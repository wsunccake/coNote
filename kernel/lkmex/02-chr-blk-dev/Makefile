KDIR ?= /lib/modules/$(shell uname -r)/build
PWD  := $(CURDIR)

obj-m  := demo_chr.o
obj-m  += demo_blk.o

all:
	$(MAKE) -C $(KDIR) M=$(PWD) modules

clean:
	$(MAKE) rm-mod
	$(MAKE) -C $(KDIR) M=$(PWD) clean
	rm -f *.ko;

rm-mod:
	-rmmod demo_chr
	-rmmod demo_blk

demo_chr: all
	insmod demo_chr.ko
	dmesg
	-lsmod | grep demo_chr
	-grep demo_chr /proc/devices
	rmmod demo_chr

demo_blk: all
	insmod demo_blk.ko
	dmesg
	-lsmod | grep demo_blk
	-grep demo_blk /proc/devices
	rmmod demo_blk
