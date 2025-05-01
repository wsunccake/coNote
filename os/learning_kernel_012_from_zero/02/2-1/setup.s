.code16
.text
.global _start_setup

_start_setup:
    movw %cs, %ax
    movw %ax, %ds
    movw %ax, %es
    movw $setup_msg, %bp
    movw $0x1301, %ax
    movw $0x0002, %bx
    movw $16, %cx
    movb $3, %dh
    movb $0, %dl
    int $0x10

setup_msg:
    .ascii "setup is running"
