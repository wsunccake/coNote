[org 0x7c00]        ; 告知編譯器，這段代碼會被 BIOS 載入到記憶體 0x7C00 位址

    mov ah, 0x0E    ; AH = 0x0E (BIOS Teletype 顯示功能)

    mov al, 'O'     ; 將字元 'O' 存入 AL 暫存器
    int 0x10        ; 呼叫 BIOS 視訊中斷

    mov al, 'S'     ; 將字元 'S' 存入 AL 暫存器
    int 0x10        ; 呼叫 BIOS 視訊中斷

    mov al, '!'     ; 將字元 '!' 存入 AL 暫存器
    int 0x10        ; 呼叫 BIOS 視訊中斷

hang:
    jmp hang        ; 無限迴圈，防止 CPU 繼續執行後方的未定義記憶體

; MBR 512 位元組填充與開機簽章
times 510-($-$$) db 0   ; 填充 0 直到第 510 個位元組
dw 0xAA55               ; 第 511-512 位元組寫入 0xAA55 (MBR 魔術數字)
