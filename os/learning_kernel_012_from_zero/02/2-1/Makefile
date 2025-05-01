CC = gcc
CFLAGS = 

AS = as
ASFLAGS = -g

LD = ld
LDFLAGS = -m elf_x86_64 -Ttext 0x0 --oformat binary

IMG = linux.img
OBJECTS = bootsect setup 
QEMU = qemu-system-i386

all: ${IMG}

run: ${IMG}
	${QEMU} -boot a -fda  ${IMG}

debug: ${IMG}
	${QEMU} -s -S -drive file=${IMG},format=raw

%.o: %.s
	$(AS) $(ASFLAGS) -o $@ $< 

bootsect: bootsect.o
	$(LD) $(LDFLAGS) -o $@ $<

setup: setup.o
	$(LD) $(LDFLAGS) -e _start_setup -o $@ $<


${IMG}: ${OBJECTS}
	cat bootsect setup> $@ 

clean:
	rm -f *.o
	rm -f ${OBJECTS}
	rm -f ${IMG}

.PHONY:
	all
	run
	exe
	clean
	debug