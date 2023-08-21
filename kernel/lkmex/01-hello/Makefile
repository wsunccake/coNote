KDIR ?= /lib/modules/$(shell uname -r)/build
PWD  := $(CURDIR)

obj-m  := hello.o
obj-m  += hi.o
obj-m  += printp.o

all:
	$(MAKE) -C $(KDIR) M=$(PWD) modules

clean:
	$(MAKE) -C $(KDIR) M=$(PWD) clean
	rm -f *.ko;

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
