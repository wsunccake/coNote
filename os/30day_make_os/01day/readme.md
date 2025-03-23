# 01 day

## boot from binary code

```bash
linux:~ # dd if=/dev/zero of=boot.bin bs=512 count=1

linux:~ # echo -ne \
"\xB4\x0E\xB0\x48\xCD\x10\xB0\x65\xCD\x10\xB0\x6C\xCD\x10\xB0\x6C\xCD\x10"\
"\xB0\x6F\xCD\x10\xB0\x20\xCD\x10\xB0\x57\xCD\x10\xB0\x6F\xCD\x10\xB0\x72"\
"\xCD\x10\xB0\x6C\xCD\x10\xB0\x64\xCD\x10\xB0\x21\xCD\x10\xF4" | dd of=boot.bin bs=1 seek=0 conv=notrunc
linux:~ # echo -ne "\x55\xAA" | dd of=boot.bin bs=1 seek=510 conv=notrunc

linux:~ # hexedit boot.bin
linux:~ # qemu-system-x86_64 -drive format=raw,file=boot.bin --nographic
```

## boot from assembly code

```asm
[BITS 16]
ORG 0x7C00   ; BIOS 會載入 MBR 到 0x7C00

start:
    DB 0xB4, 0x0E  ; MOV AH, 0Eh   ; 設定 BIOS 文字輸出功能
    DB 0xB0, 0x48  ; MOV AL, 'H'   ; 字元 'H'
    DB 0xCD, 0x10  ; INT 10h        ; 顯示字元

    DB 0xB0, 0x65  ; MOV AL, 'e'
    DB 0xCD, 0x10  ; INT 10h

    DB 0xB0, 0x6C  ; MOV AL, 'l'
    DB 0xCD, 0x10  ; INT 10h

    DB 0xB0, 0x6C  ; MOV AL, 'l'
    DB 0xCD, 0x10  ; INT 10h

    DB 0xB0, 0x6F  ; MOV AL, 'o'
    DB 0xCD, 0x10  ; INT 10h

    DB 0xB0, 0x20  ; MOV AL, ' '   ; 顯示空格
    DB 0xCD, 0x10  ; INT 10h

    DB 0xB0, 0x57  ; MOV AL, 'W'
    DB 0xCD, 0x10  ; INT 10h

    DB 0xB0, 0x6F  ; MOV AL, 'o'
    DB 0xCD, 0x10  ; INT 10h

    DB 0xB0, 0x72  ; MOV AL, 'r'
    DB 0xCD, 0x10  ; INT 10h

    DB 0xB0, 0x6C  ; MOV AL, 'l'
    DB 0xCD, 0x10  ; INT 10h

    DB 0xB0, 0x64  ; MOV AL, 'd'
    DB 0xCD, 0x10  ; INT 10h

    DB 0xB0, 0x21  ; MOV AL, '!'
    DB 0xCD, 0x10  ; INT 10h

    DB 0xF4        ; HLT (停止執行)

    ; 透過 DB 定義 0 填充到 510 Bytes
    %assign i ($ - $$)
    %rep 510 - i
        DB 0
    %endrep

    ; 透過 DB 定義 0xAA55 (Little Endian)
    DB 0x55, 0xAA
```

```bash
linux:~ # nasm -f bin boot.asm -o boot.img
linux:~ # qemu-system-x86_64 -drive format=raw,file=boot.img --nographic
```
