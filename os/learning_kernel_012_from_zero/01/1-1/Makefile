CC = gcc
CFLAGS = -m32

AS = as
ASFLAGS =

LD = ld
LDFLAGS = -m elf_x86_64 -Ttext 0x0 --oformat binary

IMG = linux.img
OBJECTS = bootsect.o
QEMU = qemu-system-i386

all: ${IMG}

run: ${IMG}
	${QEMU} -boot a -fda  ${IMG}

%.o: %.s
	$(AS) $(ASFLAGS) $< -o $@

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

exe: asm.o
	$(CC) $(CFLAGS) $< -o asm

linux.img: ${OBJECTS}
	$(LD) $(LDFLAGS) -o ${IMG} ${OBJECTS}

clean:
	rm -f *.o
	rm -f ${IMG}

.PHONY:
	all
	run
	exe
	clean