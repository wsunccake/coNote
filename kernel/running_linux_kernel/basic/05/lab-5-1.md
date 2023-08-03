# simple kernel module

---

## content

- [hello.c](#helloc)
- [makefile](#makefile)
- [usage](#usage)
- [ref](#ref)

---

## hello.c

```c
// test/hello.c
#include <linux/init.h>
#include <linux/module.h>
#include <linux/printk.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("LKMPG");
MODULE_DESCRIPTION("A sample driver");

static int __init init_hello(void)
{
    printk(KERN_INFO "Init Hello\n");
    return 0;
}

static void __exit exit_hello(void)
{
    pr_info("Exit Hello\n");
}

module_init(init_hello);
module_exit(exit_hello);
```

---

## makefile

```makefile
# test/Makefile
obj-m += hello.o

PWD := $(CURDIR)
KDIR := /lib/modules/$(shell uname -r)/build

all:
	make -C ${KDIR} M=$(PWD) modules

clean:
	make -C ${KDIR} M=$(PWD) clean
```

---

## usage

```bash
linux:~/test/ # make
linux:~/test/ # insmod hello.ko

linux:~/test/ # dmesg

linux:~/test/ # rmmod hello
```

---

## ref

[The Linux Kernel Module Programming Guide](https://sysprog21.github.io/lkmpg/#hello-world)
