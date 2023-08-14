# kernel module with parameter

---

## content

- [hello.c](#helloc)
  - [function](#function)
- [makefile](#makefile)
- [usage](#usage)
- [ref](#ref)

---

## hello.c

```c
// test/hello.c
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/module.h>
#include <linux/moduleparam.h>
#include <linux/printk.h>
#include <linux/stat.h>

MODULE_LICENSE("GPL");

static short int myshort = 1;
static int myint = 420;
static long int mylong = 9999;
static char *mystring = "blah";
static int myintarray[2] = { 420, 420 };
static int arr_argc = 0;

module_param(myshort, short, S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP);
MODULE_PARM_DESC(myshort, "A short integer");
module_param(myint, int, S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH);
MODULE_PARM_DESC(myint, "An integer");
module_param(mylong, long, S_IRUSR);
MODULE_PARM_DESC(mylong, "A long integer");
module_param(mystring, charp, 0000);
MODULE_PARM_DESC(mystring, "A character string");

module_param_array(myintarray, int, &arr_argc, 0000);
MODULE_PARM_DESC(myintarray, "An array of integers");

static int __init hello_init(void)
{
    int i;

    pr_info("Hello, world\n=============\n");
    pr_info("myshort is a short integer: %hd\n", myshort);
    pr_info("myint is an integer: %d\n", myint);
    pr_info("mylong is a long integer: %ld\n", mylong);
    pr_info("mystring is a string: %s\n", mystring);

    for (i = 0; i < ARRAY_SIZE(myintarray); i++)
        pr_info("myintarray[%d] = %d\n", i, myintarray[i]);

    pr_info("got %d arguments for myintarray.\n", arr_argc);
    return 0;
}

static void __exit hello_exit(void)
{
    printk(KERN_INFO "Goodbye, world\n");
}

module_init(hello_init);
module_exit(hello_exit);
```

### function

```c
// linux/moduleparam.h
#define module_param(name, type, perm)				\
	module_param_named(name, name, type, perm)

#define module_param_array(name, type, nump, perm)		\
	module_param_array_named(name, name, type, nump, perm)

#define MODULE_PARM_DESC(_parm, desc) \
	__MODULE_INFO(parm, _parm, #_parm ":" desc)
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
linux:~/test/ # insmod hello.ko [mystring="bebop"] [myintarray=-1,-1]

linux:~/test/ # dmesg

linux:~/test/ # rmmod hello
```

---

## ref

[The Linux Kernel Module Programming Guide](https://sysprog21.github.io/lkmpg/#hello-world)
