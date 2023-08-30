KDIR ?= /lib/modules/$(shell uname -r)/build
PWD  := $(CURDIR)

obj-m  := demo_chr.o
obj-m  += demo_blk.o
obj-m  += hello_chr.o
obj-m  += hc_dev.o

all:
	$(MAKE) -C $(KDIR) M=$(PWD) modules

clean:
	$(MAKE) rm-mod
	$(MAKE) -C $(KDIR) M=$(PWD) clean
	rm -f *.ko;

rm-mod:
	-rmmod demo_chr
	-rmmod demo_blk

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

hc_dev: all
	echo 'SUBSYSTEM=="hc_dev", MODE="0644"' > /etc/udev/rules.d/99-demo.rules
	insmod hc_dev
	dmesg
	lsmod | grep hc_dev
	ls -l /dev/hc_dev*
	ls -l /sys/class/hc_dev
	cat /dev/hc_dev0
	dmesg
	udevadm info -a -p `udevadm info -q path -n /dev/hc_dev0`
	rmmod hc_dev
