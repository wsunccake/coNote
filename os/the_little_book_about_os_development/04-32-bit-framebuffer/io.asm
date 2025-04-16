global outb
; outb - send a byte to an I/O port
; stack: [esp + 8] the data byte
;	[esp + 4] the I/O port
;	[esp] return address
; 輸出一個位元組到 I/O 埠
outb:
    mov al, [esp + 8]    ; 把 data byte 移到 AL (因為 out 指令會從 AL 輸出)
    mov dx, [esp + 4]    ; 把 port number 移到 DX (out 指令使用 DX 作為目標埠)
    out dx, al           ; 把 AL 的值輸出到指定的 port
    ret                  ; 回到呼叫處

global inb
; inb - send a byte from the given I/O port
; stack: [esp + 4] the address of the I/O port
;	[esp] return address
inb:
    mov dx, [esp + 4]    ; port number 傳入 DX
    in al, dx            ; 從該 port 讀取一個 byte 到 AL
    ret                  ; 回傳結果（注意：回傳值會存在 AL，呼叫者需處理）
