include ../../Makefile.default

IMG = boot.img
OBJS := ipl.bin

%.bin: %.asm
	$(AS) -f bin $< -o $@ -l $(basename $@).lst

image: ipl.bin
	dd if=/dev/zero of=$(IMG) bs=512 count=2880
	dd if=ipl.bin of=$(IMG) bs=512 seek=0 conv=notrunc

all: ${OBJS} image

run: $(IMG)
	$(QEMU) -drive format=raw,file=$(IMG),if=floppy -boot a -m 16 --nographic

clean:
	rm -f *.bin *.lst *.img

.PHONY:
	all
