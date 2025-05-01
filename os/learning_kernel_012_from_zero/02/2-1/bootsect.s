SETUPLEN = 4
BOOTSEG = 0x07C0
INITSEG = 0x0900
SETUPSEG = 0x9020
SYSSEG = 0x1000
ENDSEG = SYSSEG + SYSSIZE
ROOT_DEV = 0x000

.code16
.text

.global _start
_start:
    ljmp $BOOTSEG, $start2

start2:
    movw $BOOTSEG, %ax
    movw %ax, %ds
    movw $INITSEG, %ax
    movw %ax, %es
    movw $256, %cx
    xorw %si, %si     # subw %si, %si
    xorw %di, %di     # subw %di, %di
    rep movsw

    jmpl $INITSEG, $go

go:
    movw %cs, %ax
    movw %ax, %ds
    movw %ax, %es
    movw %ax, %ss
    movw $0xFF00, %sp

load_setup:
    movw $SETUPSEG, %ax  # 將 setup 載入到 SETUPSEG
    movw %ax, %es
    # subw %bx, %bx        # offset 0x0000

    movw $0x0000, %dx
    movw $0x0002, %cx
    # movw $0x0200, %bx
    movb $SETUPLEN, %al
    movb $0x02, %ah
    int $0x13
    jnc ok_load_setup
    
    movw $0x0000, %dx
    movw $0x0000, %ax
    int $0x13
    jnc load_setup

ok_load_setup:
    movw %cs, %ax       # 讓 ES 指向目前 segment
    movw %ax, %es       # ES = CS
    movw $msg, %bp      # BP = offset of msg (例如 0x7xxx)
    movw $0x1301, %ax
    movw $0x000C, %bx
    movw $21, %cx
    movb $0, %dl
    movb $0, %dh
    int $0x10

    jmpl $SETUPSEG, $0x0000
    
msg:
.ascii "setup has been loaded"

.org 508
root_dev:
    .word ROOT_DEV

boot_flag:
    .word 0xAA55
