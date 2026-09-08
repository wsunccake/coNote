include ../../Makefile.default

BOOT_SRC := lab1_2.asm
BOOT_BIN := lab1_2.bin

.PHONY: all run clean
all: $(BOOT_BIN) run

$(BOOT_BIN): $(BOOT_SRC)
	$(AS) -f bin $(BOOT_SRC) -o $(BOOT_BIN)

run: $(BOOT_BIN)
	$(QEMU)	-fda $(BOOT_BIN)
# 	$(QEMU) -drive format=raw,file=$(BOOT_BIN) --nographic

clean:
	rm -fv $(BOOT_BIN)
