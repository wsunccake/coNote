CC = gcc
CFLAGS = -ffreestanding -mno-red-zone -m64 
AS = nasm
ASFLAGS = -f elf64
LD = ld
LDFLAGS = -T linker.ld -melf_x86_64

OBJECTS = boot.o kernel.o
IMG = os.iso
QEMU = qemu-system-x86_64


all: kernel.elf

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

%.o: %.asm
	$(AS) $(ASFLAGS) $< -o $@

kernel.elf: $(OBJECTS)
	$(LD) $(LDFLAGS) $(OBJECTS) -o kernel.elf

img: kernel.elf
	mkdir -p iso/boot/grub
	cp kernel.elf iso/boot
	cp grub.cfg iso/boot/grub
	grub-mkrescue -o ${IMG} iso

check: kernel.elf
	file $<
	readelf -h $<
	objdump -x $<

run: img
	qemu-system-x86_64 -cdrom ${IMG} -nographic -serial mon:stdio

clean:
	@rm -rfv *.o kernel.elf

clean_build: clean
	@rm -rfv iso

clean_all: clean_build
	@rm -rfv ${IMG}

help:
	@echo "help        : help"
	@echo "all         : build kernel.elf"
	@echo "check       : check kernel.elf"
	@echo "img         : build image"
	@echo "run         : use qemu boot"
	@echo "clean       : remove all object"
	@echo "clean_build : remove all object and build folder"
	@echo "clean_all   : remove all object and build folder, image"

.PHONY:
	help
	all
	img
	check
	run
	clean
	clean_build
	clean_all
