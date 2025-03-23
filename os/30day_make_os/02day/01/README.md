# 02 day

## boot from assembly code

```asm
BITS 16
ORG 0x7C00

; FAT12 文件系統引導扇區
jmp short start
nop
DB "HELLOIPL"         ; 引導扇區標籤（8字節）
DW 512                ; 每個扇區大小（512字節）
DB 1                  ; 簇大小（1個扇區）
DW 1                  ; FAT 起始扇區
DB 2                  ; FAT 數量（必須是 2）
DW 224                ; 根目錄條目數（224 個）
DW 2880               ; 磁盤總扇區數（1.44MB 軟碟）
DB 0xF0               ; 介質類型（0xF0 表示 1.44MB 軟碟）
DW 9                  ; 每個 FAT 區段大小（9 扇區）
DW 18                 ; 每磁道扇區數（18 扇區）
DW 2                  ; 磁頭數量（2 個）
DD 0                  ; 隱藏扇區數（0）
DD 2880               ; 總扇區數（重新計算）

DB 0,0,0x29           ; 固定值
DD 0xFFFFFFFF         ; 卷標號
DB "HELLO-OS   "      ; 磁盤名稱（11字節）
DB "FAT12   "         ; 文件系統類型（8字節）
RESB 18               ; 填充

start:
    mov ax, 0x0000
    mov ds, ax
    mov es, ax
    mov ss, ax
    mov sp, 0x7C00
    mov si, message

print_loop:
    lodsb
    or al, al
    jz halt
    mov ah, 0x0E
    mov bh, 0x00
    int 0x10
    jmp print_loop

halt:
    hlt
    jmp halt

message:
    DB 0x0A, 0x0A
    DB "hello, world", 0x0A, 0

times 510 - ($ - $$) db 0
dw 0xAA55             ; 開機標記
```

```bash
linux:~ # nasm -f bin boot.asm -o boot.img
linux:~ # qemu-system-x86_64 -drive format=raw,file=boot.img --nographic
```
