KDIR ?= /lib/modules/$(shell uname -r)/build
PWD  := $(CURDIR)

obj-m  := hc_op.o

all:
	$(MAKE) -C $(KDIR) M=$(PWD) modules

clean:
	$(MAKE) rm-mod
	$(MAKE) -C $(KDIR) M=$(PWD) clean
	rm -f *.ko;

rm-mod:
	-rmmod demo_chr
	-rmmod demo_blk


hc_op: all
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
