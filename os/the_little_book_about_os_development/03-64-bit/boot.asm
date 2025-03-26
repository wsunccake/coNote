global loader                   ; 內核入口點，對應 linker script 的 `_start`
extern kernel_main              ; C 內核的主函式

section .multiboot
align 8
    dd 0xE85250D6               ; Multiboot 2 magic number
    dd 0                        ; Architecture (x86-64)
    dd 24                       ; Header length
    dd -(0xE85250D6 + 0 + 24)   ; Checksum

section .bss
align 8
kernel_stack:
    resb 4096                   ; 保留 4KB 堆疊

section .text
align 8
loader:
    cli                         ; 關閉中斷
    mov rsp, kernel_stack + 4096 ; 設定 64 位元堆疊
    call kernel_main            ; 呼叫 C 內核主函式

.loop:
    hlt                         ; 讓 CPU 進入休眠狀態
    jmp .loop                   ; 無限迴圈

