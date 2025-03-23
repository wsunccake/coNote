include ../../Makefile.default

BOOT_SRC = boot.asm
IMG = boot.img

all: $(IMG) run

$(IMG): $(BOOT_SRC)
	$(AS) -f bin $(BOOT_SRC) -o $(IMG)

run: $(IMG)
	$(QEMU) -drive format=raw,file=$(IMG) --nographic

clean:
	rm -f $(IMG)

