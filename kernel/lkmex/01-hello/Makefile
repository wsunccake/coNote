KDIR ?= /lib/modules/$(shell uname -r)/build
PWD  := $(CURDIR)

obj-m  := hello.o
obj-m  += hi.o printp.o
obj-m  += hey.o

all:
	$(MAKE) -C $(KDIR) M=$(PWD) modules

clean:
	$(MAKE) rm-mod
	$(MAKE) -C $(KDIR) M=$(PWD) clean
	rm -f *.ko;

rm-mod:
	-rmmod hello
	-rmmod printp
	-rmmod hi
	-rmmod hey

hello: all
	insmod hello.ko
	dmesg
	lsmod | grep hello
	rmmod hello

hi: all
	insmod hi.ko
	insmod printp.ko
	dmesg
	lsmod | grep printp
	rmmod printp
	rmmod hi

hey: all
	insmod hey.ko
	dmesg
	lsmod | grep hey
	rmmod hey
	insmod hi.ko cnt=3 hi="kernel" arr=3,4,1
	dmesg
	rmmod hey

