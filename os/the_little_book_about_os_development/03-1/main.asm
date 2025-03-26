section .text
    global main
    extern sum_of_three
    extern print_result

main:
    push rbp
    mov rbp, rsp

    ; 依照 cdecl 風格，從右到左 push 參數
    push 3            ; arg3
    push 2            ; arg2
    push 1            ; arg1

    ; 從堆疊中提取參數到正確的暫存器 (System V AMD64 ABI)
    pop rdi           ; 第一個參數 arg1 = 1
    pop rsi           ; 第二個參數 arg2 = 2
    pop rdx           ; 第三個參數 arg3 = 3

    ; 呼叫 C 函數 sum_of_three
    call sum_of_three

    ; 傳遞返回值給 print_result
    mov rdi, rax
    call print_result

    ; 結束程式
    mov eax, 0
    pop rbp
    ret

