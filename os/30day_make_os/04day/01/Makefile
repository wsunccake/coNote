include ../../Makefile.default

IMG = boot.img
SYS = haribote.sys
OBJS := ipl.bin head.bin func.o bootpack.hrb ${SYS}

ipl.bin  : ipl.asm
head.bin : head.asm
func.o   : func.asm

%.bin: %.asm
	$(AS) -f bin $< -o $@ -l $(basename $@).lst

%.o: %.asm
	$(AS) -f elf $< -o $@ -l $(basename $@).lst

bootpack.hrb : bootpack.c func.o os.lds
	gcc -march=i486 -m32 -nostdlib -g -O0 \
		-T os.lds \
		-o $@ \
		-fno-pie\
		bootpack.c func.o

${SYS} : head.bin bootpack.hrb
%.sys :
	cat $^ > $@

image: ipl.bin head.bin
	dd if=/dev/zero of=$(IMG) bs=512 count=2880
	dd if=ipl.bin of=$(IMG) bs=512 seek=0 conv=notrunc
	mcopy -i $(IMG) ${SYS} ::
	# dd if=head.bin of=$(IMG) bs=512 seek=1 conv=notrunc

all: ${OBJS} image

run: $(IMG)
	$(QEMU) -drive format=raw,file=$(IMG),if=floppy -boot a -m 16 --nographic

clean:
	rm -f ${OBJS} ${IMG}

.PHONY:
	all
