section .text
    global main
    extern sum_of_three
    extern print_result

main:
    ; 設定參數
    mov rdi, 5      ; 第一個參數 arg1 = 5
    mov rsi, 10     ; 第二個參數 arg2 = 10
    mov rdx, 15     ; 第三個參數 arg3 = 15

    ; 呼叫 sum_of_three 函數
    call sum_of_three

    ; 返回值在 RAX，傳給 print_result
    mov rdi, rax    ; 設定 print_result 的參數 (sum_of_three 的回傳值)
    call print_result

    ; 結束程式
    mov eax, 0
    ret

