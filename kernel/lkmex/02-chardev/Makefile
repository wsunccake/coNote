KDIR ?= /lib/modules/$(shell uname -r)/build
PWD  := $(CURDIR)

obj-m  := demo_chr.o
obj-m  += demo_blk.o
obj-m  += hello_chr.o

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

add-chr:
#	awk '"hello_chr" == $2 {print $1}' /proc/devices
#	mknod /dev/hello_chr0 c 239 0
	@cid=$(shell awk '"hello_chr" == $$2 {print $$1}' /proc/devices); \
	       	$(foreach n, $(shell seq 0 1), mknod /dev/hello_chr$(n) c $${cid} $(n);)

rm-chr:
	-@for n in $(shell seq 0 1); do \
		rm -fv /dev/hello_chr$${n};\
	done


op-chr:
	$(foreach n, $(shell seq 0 1), \
		echo /dev/hello_chr$(n); \
		dmesg | tail -5; \
		echo; \
		hostname > /dev/hello_chr$(n); \
		dmesg | tail -5; \
		echo; \
		cat /dev/hello_chr$(n); \
		dmesg | tail -5; \
		echo;)

hello_chr: all
	insmod hello_chr.ko
	dmesg
	lsmod | grep hello_chr
	$(MAKE) add-chr
	-ls -l /dev/hello_chr*
	$(MAKE) op-chr
	$(MAKE) rm-chr
	rmmod hello_chr
