BOOTSEG = 0x07C0

.code16
.text

.global _start
_start:
    ljmp $BOOTSEG, $start2

start2:
    movw $BOOTSEG, %ax
    movw %ax, %ds
    movw %ax, %es
    movw %ax, %fs
    movw %ax, %gs

    lea msg, %bp
    movw $0x1301, %ax
    movw $0x000C, %bx
    movw $12, %cx
    movb $0, %dl
    int $0x10

loop:
    jmp loop

msg:
.ascii "Hello World!"

.org 510
boot_flag:
    .word 0xAA55
