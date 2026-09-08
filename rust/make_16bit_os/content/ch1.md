# ch1. X86 實模式與底層架構

歡迎來到作業系統開發的第一站。在現代 64 位元作業系統（如 Linux、macOS）中，硬體細節被 OS 核心與 CPU 保護模式高度抽象化；然而，當電腦剛按下電源鍵時，x86 CPU 為了向下相容，會強制進入 16-bit 實模式 (Real Mode)。

本文件梳理 x86 實模式的核心機制、暫存器用途、記憶體定址與 BIOS 中斷，這是撰寫 Bootloader 與 16-bit OS 核心不可或缺的理論基礎。

## 一、 X86 實模式與 16-bit 暫存器

CPU 暫存器 (Registers) 是位於 CPU 內部最高速的儲存單元。在 16-bit 實模式下，所有通用暫存器大小皆為 16 位元 (2 Bytes)。

1. 通用暫存器 (General Purpose Registers)

通用暫存器通常可以拆分為高 8 位元 (High) 與低 8 位元 (Low) 獨立存取。例如 AX 可拆為 AH 與 AL。

- AX (Accumulator)：累加暫存器。常用於算術運算、I/O 埠讀寫以及 BIOS 中斷呼叫的指令碼傳遞。
- BX (Base)：基底暫存器。常用作記憶體定址的基底位址 (Base Address)。
- CX (Count)：計數暫存器。用於迴圈（如 LOOP 指令）或字串操作的重複計數器。
- DX (Data)：資料暫存器。常與 AX 搭配進行 32 位元乘除法，或存放 I/O 埠位址。

2. 變址與指標暫存器 (Index & Pointer Registers)

無法拆分為 8 位元，主要用於記憶體定址與堆疊操作。

- SI (Source Index)：來源變址暫存器。字串複製/比較時指向來源記憶體位址。
- DI (Destination Index)：目的變址暫存器。字串操作時指向目的記憶體位址。
- SP (Stack Pointer)：堆疊頂端指標。時刻指向目前堆疊 (Stack) 最頂端的位址。
- BP (Base Pointer)：堆疊基底指標。常用於函式呼叫時，存取區域變數或傳遞的參數。

3. 段暫存器 (Segment Registers)

實模式下決定記憶體分段的核心暫存器（無法拆分）。

- CS (Code Segment)：程式碼段暫存器。搭配指令指標 IP (Instruction Pointer)，決定 CPU 下一條要執行的指令位址 (CS:IP)。
- DS (Data Segment)：資料段暫存器。預設的資料讀寫記憶體段。
- SS (Stack Segment)：堆疊段暫存器。搭配 SP 決定堆疊記憶體區域 (SS:SP)。
- ES (Extra Segment)：附加段暫存器。常與 DI 搭配用於跨段記憶體複製。

---

## 二、 記憶體定址原理：分段機制 (Segmented Memory)

在 16-bit 實模式下，暫存器只有 16 位元，最大只能存取 $2^{16} = 65,536\text{ Bytes} = 64\text{ KB}$ 的位址空間。然而，早期 8086 CPU 擁有 20 條位址線 (Address Bus)，硬體實際上支援 $2^{20} = 1\text{ MB}$ 的記憶體空間。

為了用 16 位元的暫存器存取 20 位元的實體記憶體，Intel 設計了 分段機制 (Segmentation)：

$$\text{實體記憶體位址 (Physical Address)} = (\text{Segment} \times 16) + \text{Offset}$$

寫法通常表示為：Segment:Offset（段位址 : 偏移量）。

**定址算術範例**

假設 CS = 0x07C0，IP = 0x0005：

1. 將段位址左移 4 個 Bit（即乘以 16，相當於十六進位末尾補一個 0）：0x07C0 $\rightarrow$ 0x07C00
2. 加算偏移量：0x07C00 + 0x0005 = 0x07C05
3. CPU 最終會向實體記憶體匯流排發出 0x07C05 這個 20-bit 位址。
   Bootloader 關鍵位址：BIOS 開機載入 MBR Bootloader 時，預設會將 512 位元組載入至實體位址 0x7C00。這代表將段暫存器設為 0x07C0 (Offset 0x0000)，或是段暫存器設為 0x0000 (Offset 0x7C00)。

---

## 三、 BIOS 中斷機制 (BIOS Interrupts)

在沒有 OS 驅動程式的裸機 (Bare-metal) 環境下，BIOS 提供了基礎的軟體中斷服務，讓開發者透過 CPU 指令 INT <向量號碼> 呼叫硬體功能。

呼叫 BIOS 中斷的固定模式：

1.  將功能碼寫入 AH 暫存器。
2.  將所需的參數放入 AL, BX, CX, DX 等暫存器。
3.  觸發中斷指令（例如 INT 10h）。

### 1. INT 10h：VGA 螢幕顯示控制

- AH = 0x0E（Teletype 模式印字）：在游標位置印出單一字元，並自動將游標後移。
  - AL = ASCII 字元編碼
  - BH = 頁碼 (通常設 0)
  - BL = 前景色彩 (文字模式下可設 0x07)

```assembly
mov ah, 0x0E    ; 設定 BIOS 功能碼：印出字元
mov al, 'A'     ; 欲印出的字元
int 0x10        ; 呼叫 BIOS 視訊中斷，螢幕印出 'A'
```

### 2. INT 13h：軟碟/磁碟低階讀寫

OS 載入器需要透過此中斷將磁碟上的 OS 核心讀入記憶體。

- AH = 0x02（讀取磁碟區塊 Sector）：
  - AL = 欲讀取的磁碟區塊數量
  - CH = 柱面號 (Cylinder, 0 起算)
  - CL = 磁區號 (Sector, 1 起算，Bit 0-5)
  - DH = 磁頭號 (Head, 0 起算)
  - DL = 驅動器編碼 (0x00 代表第一個軟碟 A:, 0x80 代表第一個硬碟)
  - ES:BX = 讀取後的資料要寫入的記憶體目標位址 (Segment:Offset)

---

## 四、 x86 組合語言語法與指標概念對比

在 OS 開發中，高階語言的指標本質上就是實體記憶體的位址。

### 1. 常用 x86 彙編指令 (NASM 語法)

```assembly
mov ax, 0x1234      ; 將立即值 0x1234 賦予 AX
mov ds, ax          ; 將 AX 的值轉存至 DS (段暫存器無法直接 mov 立即值)

mov [0x1000], ax    ; 【記憶體寫入】將 AX 的值寫入 DS:0x1000 位址（括號表示取位址內容）
mov cx, [0x1000]    ; 【記憶體讀取】讀取 DS:0x1000 的資料寫入 CX

add ax, bx          ; AX = AX + BX
cmp ax, 0           ; 比較 AX 與 0 (會影響 Flags 暫存器)
je label_equal      ; 若相等 (Jump if Equal)，跳轉至 label_equal 標籤處執行

push ax             ; 將 AX 壓入堆疊 (SP = SP - 2)
pop bx              ; 從堆疊彈出資料至 BX (SP = SP + 2)
```

### 2. 記憶體指標 (Pointer) 對照：C / Rust vs Assembly

假設要將記憶體位址 0xB8000（VGA 文字模式顯示記憶體首位址）寫入字元 'H' (ASCII 0x48) 與顏色屬性 0x0F（黑底白字）：

- Assembly (NASM)

```assembly
mov ax, 0xB800
mov es, ax          ; ES = 0xB800 (實體位址 0xB8000)
mov byte [es:0x0000], 'H'   ; 寫入字元
mov byte [es:0x0001], 0x0F  ; 寫入顏色屬性
```

- C 語言 (裸機指標操作)

```c
// 建立一個指向實體位址 0xB8000 的 volatile 指標
volatile unsigned char *vga = (unsigned char *)0xB8000;
vga[0] = 'H';   ; 寫入字元
vga[1] = 0x0F;  ; 寫入顏色屬性
```

- Rust (裸機不安全指標操作 #![no_std])

```rust
// 使用 Rust 裸指標 (Raw Pointer) 直接操作硬體記憶體
unsafe {
    let vga = 0xB8000 as *mut u8;
    *vga = b'H';          // 寫入字元
    *vga.add(1) = 0x0F;   // 寫入顏色屬性
}
```

關鍵理解：不論是 C 還是 Rust，在 #![no_std] 裸機環境下沒有作業系統幫你管理記憶體分頁或指標保護，指標 (Pointer) 的轉型與解引用就是直接對實體記憶體暫存器進行電位開關讀寫。這也是為何 Rust 需要使用 unsafe 區塊宣告的原因。