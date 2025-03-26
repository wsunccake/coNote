CC = gcc
CFLAGS = -m32 -nostdlib -nostdinc -fno-builtin -fno-stack-protector \
         -nostartfiles -nodefaultlibs -Wall -Wextra -Werror -c
AS = nasm
ASFLAGS = -f elf
LD = ld
LDFLAGS = -T link.ld -melf_i386

OBJECTS = loader.o kmain.o
IMG = os.iso
QEMU = qemu-system-i386


all: kernel.elf

%.o: %.c
	$(CC) $(CFLAGS)  $< -o $@

%.o: %.asm
	$(AS) $(ASFLAGS) $< -o $@

kernel.elf: $(OBJECTS)
	$(LD) $(LDFLAGS) $(OBJECTS) -o kernel.elf

img: kernel.elf
	mkdir -p iso/boot/grub
	cp kernel.elf iso/boot
	cp menu.lst iso/boot/grub
	cp stage2_eltorito iso/boot/grub
	genisoimage -R \
      -b boot/grub/stage2_eltorito \
      -no-emul-boot \
      -boot-load-size 4 \
      -A os \
      -input-charset utf8 \
      -quiet \
      -boot-info-table \
      -o os.iso \
      iso

run: img
	${QEMU} -cdrom ${IMG} -nographic -serial mon:stdio

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
