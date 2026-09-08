[org 0x0000]        ; 將 Offset 設為 0x0000

    ; 欲存取實體位址 0x7C00：
    ; 設定 Segment = 0x07C0 -> 0x07C0 * 16 (0x10) = 0x7C00
    mov ax, 0x07C0
    mov ds, ax      ; DS = 0x07C0

    mov ah, 0x0E    ; BIOS 印字功能

    ; 從 DS:my_text (即 0x07C0:my_text) 讀取第一個位元組
    mov al, [my_text]
    int 0x10

    ; 讀取第二個位元組 (Offset + 1)
    mov al, [my_text + 1]
    int 0x10

hang:
    jmp hang

my_text:
    db 'HI', 0      ; 定義資料字串 'H', 'I'

times 510-($-$$) db 0
dw 0xAA55
