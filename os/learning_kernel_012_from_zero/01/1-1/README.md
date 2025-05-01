# 16 bit boot loader

## bootsect.s

```S
BOOTSEG = 0x7c0
# 定義常數，叫做 BOOTSEG，值是 0x7C0。
# 為了稍後做 segment:offset 計算的。
# 注意：因為 BIOS 把 bootloader 載入到 實體位址 0x7C00，而在 real mode（16 位元）中，地址是用 segment × 16 + offset 計算的。0x7C0 × 0x10 = 0x7C00
# BOOTSEG=0x7C0，代表 segment 位置設定。

.code16
# assembler：16-bit 模式的程式碼。
# x86 CPU 在開機時是進入 real mode (實模式)，只有 16-bit 寄存器。

.text
# 進入 assembler 的 text 段（code section）。
# 開始是要編譯成機器碼的指令。

.global _start
# 宣告 _start 是全域符號。
# linker (ld) 在找入口點時，可以知道 _start 是 bootloader 開始的地方。

_start: # 標記一個 label，叫 _start，從這裡開始執行。
    ljmp $BOOTSEG, $start2
# ljmp = long jump
# ljmp 語法是這樣：ljmp segment, offset
# 一開機，BIOS 把 Bootloader 複製到記憶體 0x0000:0x7C00。
# CPU 的 CS:IP （Code Segment : Instruction Pointer）會設成：
# CS = 0x0000
# IP = 0x7C00
# CPU 在「segment + offset」的概念上，其實是 0x0000 × 16 + 0x7C00 = 0x7C00
# 要做的：
# CS = 0x07C0
# IP = offset(start2)
# physical address / 真正物理位址是：segment × 16 + offset
# segment × 16: 0x07C0 × 0x10 = 0x7C00
# offset: start2


start2: # 新的 label，代表 long jump 後程式執行到這裡。
    movw $BOOTSEG, %ax
# 把 0x7C0 放到 AX 暫存器。
# $ 表示「立即數」（immediate value），不是從記憶體讀，而是直接把數值塞進暫存器。
    movw %ax, %ds
    movw %ax, %es
    movw %ax, %fs
    movw %ax, %gs
# 把 AX 的值（0x7C0）設定到 資料段寄存器（DS、ES、FS、GS）。
# 要存取資料的時候，資料段（DS）需要是正確的（就是 0x7C0）。
# 存取 msg 這種 label 時，才能正確地計算物理地址。

    lea msg, %bp
# lea 是 Load Effective Address。
# 把 msg 的 offset 值載入到 BP。
# lea 不是真的去存取記憶體，它只是把地址計算好。
# 這裡做的是：「BP = msg的offset」。
# 注意：為什麼用 BP？因為後面 BIOS 中斷會用 BP 當作字串地址參數。

    movw $0x1301, %ax
# 設定 AH=0x13，AL=0x01。
# 呼叫 BIOS interrupt 0x10 的 AH=0x13 功能：列印字串到螢幕上（Write String）。
# AL=1 代表：每個字元印出時套用 attribute（顏色）

    movw $0x000c, %bx
# 設定 BH = 0（Page number 0），BL = 0x0C（文字顏色，亮紅色）
# BIOS int 0x10 的 Write String 需要這樣設定顏色。

    movw $12, %cx
# 設定 CX=12。
# 字串長度是 12（"Hello World!" 有 12 個字元）
    movb $0, %dl
# 設定 DL=0。
# DL 代表顯示用的行列座標（一般來說初始位置是 (0, 0)）。
    int $0x10
# 呼叫 BIOS 服務。
# int 0x10 = 呼叫 BIOS 的「螢幕文字模式功能」。
# 根據你前面設好的參數（AH=0x13, AL=1, BX=顏色, CX=字數, BP=字串指標），BIOS 會把 msg 裡的字串印出到畫面上。

loop: # 標記一個 loop 位置。
    jmp loop
# 無限迴圈
# 這樣程式不會亂跑，不然 CPU 會繼續跑到不確定的地方出錯。
# 等於說：「印完 Hello World 之後，CPU就卡在這裡等」。

msg: # 標記字串的開始。
.ascii "Hello World!"
# 把純 ASCII 字元 "Hello World!" 直接塞進目標檔案
。
.org 510
# 將組譯位置跳到第 510 位元組。
# 因為 BIOS boot sector 要求整個 boot sector 必須是 512 bytes。
# 510 個 byte 前面是你的程式碼，後面兩個 byte 放 boot flag。

boot_flag: # 標記 boot flag 開始位置。
    .word 0xaa55
# 寫入 0xAA55（小端序），即 55 AA。
# 這是標準 BIOS boot sector 的魔法字（Magic Number），BIOS 會檢查這個
# 如果不是 0xAA55，BIOS 不會認為這是個可開機磁區。
```

---

## asm.c

syntax

```c
__asm__ __volatile__ (
    "汇编指令模板 / Assembler Template"
    : 输出操作数 / Output Operands
    : 输入操作数 / Input Operands
    : 被破坏的寄存器列表 / Clobbered Registers
);
```

```c
#include <stdio.h>

int main() {
    int a = 5, b = 7, c;

    __asm__ __volatile__ (
        "movl %1, %%eax;\n\t"   // 將變數 a 的值移動到 eax 寄存器中。
        "addl %2, %%eax;\n\t"   // 將變數 b 的值加到 eax 寄存器中。
        "movl %%eax, %0;"       // 將 eax 寄存器中的結果移動到變數 c 中。​
        : "=r" (c)              // 輸出操作數
        : "r" (a), "r" (b)      // 輸入操作數
        : "%eax"                // 被破壞的暫存器
    );
// "=r"：指定變數為僅輸出的通用暫存器。
// "r"：指定變數為輸入的通用暫存器。
// "+r"：指定變數為既可讀又可寫的通用暫存器。

    printf("a + b = %d\n", c);
    return 0;
}
```

---

## makefile

```makefile
CC = gcc
CFLAGS = -m32

AS = as
ASFLAGS =

LD = ld
LDFLAGS = -m elf_x86_64 -Ttext 0x0 --oformat binary

IMG = linux.img
OBJECTS = bootsect.o
QEMU = qemu-system-i386

all: ${IMG}

run: ${IMG}
	${QEMU} -boot a -fda  ${IMG}

%.o: %.S
        $(AS) $(ASFLAGS) $< -o $@

%.o: %.c
        $(CC) $(CFLAGS) -c $< -o $@

exe: asm.o
        $(CC) $(CFLAGS) $< -o asm

linux.img: ${OBJECTS}
        $(LD) $(LDFLAGS) -o ${IMG} ${OBJECTS}

clean:
        rm -f *.o
        rm -f ${IMG}

.PHONY:
        all
	    run
        exe
        clean
```
