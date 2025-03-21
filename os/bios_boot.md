# bios bootloader

## content

- [BIOS](#bios)
- [16-bit bios bootloader](#16-bit-bios-bootloader)
  - [vdisk](#vdisk)
  - [vm](#vm)
- [32-bit bios bootloader](#32-bit-bios-bootloader)
- [64-bit bios bootloader](#64-bit-bios-bootloader)
- [boot sector](#boot-sector)

---

## BIOS

- 16-bit bootloader, Real Mode
- 32-bit bootloader, Protected Mode
- 64-bit bootloader, Long Mode

需要 NASM 和 QEMU

```bash
# for debian / ubuntu
linux:~ # apt update
linux:~ # apt install nasm qemu-system-x86

# for rhel / fedora
linux:~ # dnf makecache
linux:~ # dnf install nasm qemu-system-x86
```

---

## 16-bit bios bootloader

1. 清空螢幕
2. 在螢幕上顯示 "Hello, World!"
3. 停在這裡，等待使用者重新開機

```asm
; boot.asm
[BITS 16]            ; 16-bit real mode
[ORG 0x7C00]         ; BIOS 會將 bootloader 加載到 0x7C00

start:
    mov si, message  ; 將 SI 註冊指向 message 字串
    call print_string

    jmp $            ; 讓 CPU 停留在此，不進一步執行

print_string:
    mov ah, 0x0E     ; BIOS 文字輸出功能
.loop:
    lodsb            ; 從 SI 讀取一個字元到 AL
    test al, al      ; 檢查是否為 null 結尾
    jz .done
    int 0x10         ; 呼叫 BIOS 中斷輸出字元
    jmp .loop
.done:
    ret

message db "Hello, World!", 0  ; 訊息字串，以 null 結尾

times 510-($-$$) db 0  ; 填充，使程式總長達 510 bytes
dw 0xAA55              ; boot sector 結尾的 magic number
```

```bash
linux:~ # nasm -f bin boot.asm -o boot.bin
# boot.bin must be 512 bytes

# check boot.bin size
linux:~ # ls -lh boot.bin
-rw-r--r-- 1 root root 512 Mar 20 12:22 boot.bin
linux:~ # file boot.bin
boot.bin: DOS/MBR boot sector
linux:~ # stat boot.bin
  File: boot.bin
  Size: 512 ...

# check magic number
linux:~ # hexdump -C boot.bin | head -n 32
...
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|
linux:~ # hexdump -C boot.bin | tail -n 4
...
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|

# test
linux:~ # qemu-system-x86_64 -drive format=raw,file=boot.bin -nographic
```

### vdisk

```bash
linux:~ # dd if=/dev/zero of=floppy.img bs=512 count=2880
linux:~ # dd if=boot.bin of=floppy.img conv=notrunc
linux:~ # hexdump -C floppy.img | head -n 32
```

### vm

```bash
linux:~ # qemu-system-x86_64 -drive format=raw,file=floppy.img -nographic
```

---

## 32-bit bios bootloader

1. 關閉中斷（CLI）
2. 啟用 A20 地址線（A20 Line）
3. 設置 GDT（全域描述表，Global Descriptor Table）
4. 切換到 32-bit 受保護模式
5. 跳轉並執行 32-bit 代碼
6. 顯示 Hello, World!

```asm
; boot32.asm
[BITS 16]               ; 16-bit 真實模式
[ORG 0x7C00]            ; BIOS 會將 Bootloader 加載到 0x7C00

start:
    cli                 ; 關閉中斷
    cld                 ; 清除方向標誌，確保字串操作向前讀取

    ; 啟用 A20 地址線
    in al, 0x92
    or al, 2
    out 0x92, al

    ; 加載 GDT
    lgdt [gdt_descriptor]

    ; 設置 CR0 啟動受保護模式
    mov eax, cr0
    or eax, 1
    mov cr0, eax

    ; 跳轉到 32-bit 受保護模式
    jmp CODE_SEG:init_protected_mode

[BITS 32]
init_protected_mode:
    ; 更新段寄存器
    mov ax, DATA_SEG
    mov ds, ax
    mov es, ax
    mov fs, ax
    mov gs, ax
    mov ss, ax
    mov esp, 0x90000    ; 設置堆疊指標

    ; 顯示訊息
    mov esi, message
    call print_string

    jmp $               ; 讓 CPU 停留在這裡

print_string:
    mov ah, 0x0E        ; BIOS 文字輸出功能
.loop:
    lodsb               ; 讀取一個字元
    test al, al         ; 若為 0 則結束
    jz .done
    int 0x10            ; 顯示字元
    jmp .loop
.done:
    ret

; GDT 定義
gdt_start:
    dq 0                ; NULL descriptor

gdt_code:
    dw 0xFFFF           ; Limit (bits 0-15)
    dw 0x0000           ; Base (bits 0-15)
    db 0x00             ; Base (bits 16-23)
    db 0x9A             ; Access byte: 1 00 1 1 0 1 0 (執行、可讀、存在)
    db 0xCF             ; Flags: 1100 1111 (4K, 32-bit)
    db 0x00             ; Base (bits 24-31)

gdt_data:
    dw 0xFFFF           ; Limit
    dw 0x0000
    db 0x00
    db 0x92             ; Access byte: 1 00 1 0 0 1 0 (讀寫)
    db 0xCF
    db 0x00

gdt_end:

gdt_descriptor:
    dw gdt_end - gdt_start - 1
    dd gdt_start

CODE_SEG equ gdt_code - gdt_start
DATA_SEG equ gdt_data - gdt_start

message db "Hello, 32-bit World!", 0

times 510-($-$$) db 0  ; 填充至 510 bytes
dw 0xAA55              ; Bootloader magic number
```

---

## 64-bit bios bootloader

1. 關閉中斷（CLI）
2. 啟用 A20 地址線
3. 設置 GDT（全域描述表，Global Descriptor Table）
4. 進入 32-bit 受保護模式（Protected Mode）
5. 啟用 Paging 並設置 4 級頁表
6. 切換到 64-bit 長模式（Long Mode）
7. 跳轉到 64-bit 代碼並執行
8. 顯示 Hello, 64-bit World!

```asm
; boot64.asm
[BITS 16]
[ORG 0x7C00]

start:
    cli                 ; 關閉中斷
    cld                 ; 清除方向標誌

    ; 啟用 A20 地址線
    in al, 0x92
    or al, 2
    out 0x92, al

    ; 加載 GDT
    lgdt [gdt_descriptor]

    ; 進入 32-bit 受保護模式
    mov eax, cr0
    or eax, 1
    mov cr0, eax
    jmp CODE_SEG:init_protected_mode

[BITS 32]
init_protected_mode:
    mov ax, DATA_SEG
    mov ds, ax
    mov es, ax
    mov fs, ax
    mov gs, ax
    mov ss, ax
    mov esp, 0x90000  ; 設置堆疊指標

    ; 設置 PAE（物理地址擴展）
    mov eax, cr4
    or eax, 1 << 5
    mov cr4, eax

    ; 設置長模式（Long Mode）
    mov ecx, 0xC0000080
    rdmsr
    or eax, 1 << 8
    wrmsr

    ; 設置頁表（Identity Mapping）
    mov eax, pml4_table
    mov cr3, eax

    ; 啟用 Paging 和 Long Mode
    mov eax, cr0
    or eax, (1 << 31) | (1 << 16)
    mov cr0, eax

    ; 切換到 64-bit 長模式
    jmp CODE64_SEG:init_long_mode

[BITS 64]
init_long_mode:
    ; 更新段寄存器
    mov ax, DATA64_SEG
    mov ds, ax
    mov es, ax
    mov fs, ax
    mov gs, ax
    mov ss, ax
    mov rsp, 0x90000  ; 設置堆疊

    ; 顯示訊息
    mov rsi, message
    call print_string

    jmp $

print_string:
    mov ah, 0x0E
.loop:
    lodsb
    test al, al
    jz .done
    int 0x10
    jmp .loop
.done:
    ret

; ===== GDT 定義 =====
gdt_start:
    dq 0                ; NULL descriptor

gdt_code:
    dw 0xFFFF, 0, 0x9A, 0xAF, 0, 0  ; 32-bit Code Segment
gdt_data:
    dw 0xFFFF, 0, 0x92, 0xCF, 0, 0  ; 32-bit Data Segment
gdt_code64:
    dw 0xFFFF, 0, 0x9A, 0xAF, 0, 0  ; 64-bit Code Segment
gdt_data64:
    dw 0xFFFF, 0, 0x92, 0xAF, 0, 0  ; 64-bit Data Segment

gdt_end:

gdt_descriptor:
    dw gdt_end - gdt_start - 1
    dd gdt_start

CODE_SEG equ gdt_code - gdt_start
DATA_SEG equ gdt_data - gdt_start
CODE64_SEG equ gdt_code64 - gdt_start
DATA64_SEG equ gdt_data64 - gdt_start

; ===== 64-bit Paging =====
section .bss
align 4096
pml4_table: resq 512
pdpt_table: resq 512
pd_table:   resq 512

section .data
message db "Hello, 64-bit World!", 0

times 510-($-$$) db 0  ; 填充至 510 bytes
dw 0xAA55              ; Bootloader magic number
```

---

## boot sector

BIOS 在載入開機磁區（Boot Sector）時，要求開機程式必須 剛好是 512 bytes，這是由硬體和歷史標準決定的。

1. 硬體規範與歷史標準

傳統 BIOS（Basic Input/Output System）在開機時，會從磁碟的 第 1 個磁區（sector 0） 載入開機碼到記憶體的 0x7C00，然後開始執行。

- 標準磁碟扇區大小（Sector Size）
  早期的軟碟（Floppy Disk）和硬碟（Hard Disk）大多使用 512-byte 扇區，這成為業界標準。因此，BIOS 在讀取開機磁區時 只會讀取 512 bytes。

- 固定大小避免兼容性問題
  如果開機程式小於 512 bytes，BIOS 還是會載入 512 bytes（但程式可以填充 0x00）。
  如果開機程式超過 512 bytes，BIOS 只會載入 前 512 bytes，剩下的部分會被忽略，可能導致程式無法正常運行。

2. Magic Number (0xAA55) 確保可開機

BIOS 在讀取 512-byte 開機磁區後，會檢查最後 2 個 bytes 是否是 0xAA55（Little Endian 格式：55 AA）。這稱為 Boot Signature，代表這個磁區是可開機

```asm
times 510-($-$$) db 0   ; 填充 0，使總大小達 510 bytes
dw 0xAA55               ; 加上 boot signature
```

如果沒有這個標誌，BIOS 會認為磁碟 不是可開機的，然後跳轉到其他開機裝置（如 USB、網路開機等）。

3. 如何載入更大程式？

   由於 BIOS 只會載入 第一個 512 bytes，如果想載入更大的作業系統，通常會：

   a. Bootloader 階段（第一階段）

   - 這個 512-byte 程式只負責 最基本的初始化，例如顯示訊息、讀取磁碟等。
   - 之後它會手動載入 第二階段 bootloader 或 核心（Kernel）。

   b. 使用 BIOS 中斷讀取更多磁區

   - 利用 int 0x13 來讀取磁碟上的其他部分，例如

   ```asm
   mov ah, 0x02  ; BIOS 讀取磁碟功能
   mov al, 2     ; 讀取 2 個磁區
   mov ch, 0     ; 磁柱 0
   mov dh, 0     ; 磁頭 0
   mov cl, 2     ; 從磁區 2 開始
   mov bx, 0x9000 ; 將資料存到 0x9000 記憶體
   int 0x13
   ```

   - 這樣，開機磁區就能載入更大的 bootloader 或作業系統核心。
