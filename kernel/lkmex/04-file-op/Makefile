KDIR   ?= /lib/modules/$(shell uname -r)/build
PWD    := $(CURDIR)
# EXTRA_CFLAGS := -DLOCK_USE=0   # 0: semaphore
# EXTRA_CFLAGS := -DLOCK_USE=1   # 1: mutex

obj-m  := hc_op.o
obj-m  += hc_op_locked.o

all:
	$(MAKE) -C $(KDIR) M=$(PWD) modules

clean:
	$(MAKE) rm-mod
	$(MAKE) -C $(KDIR) M=$(PWD) clean
	rm -f *.ko;

rm-mod:
	-rmmod hc_op
	-rmmod hc_op_locked

hc_op: all
	echo 'SUBSYSTEM=="hc_dev", MODE="0644"' > /etc/udev/rules.d/99-demo.rules
	insmod hc_op.ko
	dmesg
	lsmod | grep hc_op
	ls -l /dev/hc_op*
	ls -l /sys/class/hc_op
	cat /dev/hc_op0
	dmesg
	udevadm info -a -p `udevadm info -q path -n /dev/hc_op0`
	rmmod hc_op

#	while true; do echo b > /dev/hc_op0; done

hc_op_locked: all
	echo 'SUBSYSTEM=="hc_dev", MODE="0644"' > /etc/udev/rules.d/99-demo.rules
	insmod hc_op_locked.ko
	dmesg
	lsmod | grep hc_op_locked
	ls -l /dev/hc_op_locked*
	ls -l /sys/class/hc_op_locked
	cat /dev/hc_op_locked0
	dmesg
	udevadm info -a -p `udevadm info -q path -n /dev/hc_op_locked0`
	rmmod hc_op_locked
