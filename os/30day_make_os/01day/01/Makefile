include ../../Makefile.default

BOOT_BIN = boot.bin

all: $(BOOT_BIN) run

$(BOOT_BIN):
	dd if=/dev/zero of=boot.bin bs=512 count=1
	echo -ne \
	"\xB4\x0E\xB0\x48\xCD\x10\xB0\x65\xCD\x10\xB0\x6C\xCD\x10\xB0\x6C\xCD\x10"\
	"\xB0\x6F\xCD\x10\xB0\x20\xCD\x10\xB0\x57\xCD\x10\xB0\x6F\xCD\x10\xB0\x72"\
	"\xCD\x10\xB0\x6C\xCD\x10\xB0\x64\xCD\x10\xB0\x21\xCD\x10\xF4" | dd of=$@ bs=1 seek=0 conv=notrunc
	echo -ne "\x55\xAA" | dd of=$@ bs=1 seek=510 conv=notrunc

run: $(BOOT_BIN)
	$(QEMU) -drive format=raw,file=$(BOOT_BIN) --nographic

clean:
	rm -f $(BOOT_BIN)

