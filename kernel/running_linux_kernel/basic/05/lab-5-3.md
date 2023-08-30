# kernel module with EXPORT_SYMBOL

---

## content

- [mymodule1.c](#mymodule1c)
- [mymodule2.c](#mymodule2c)
- [makefile](#makefile)
- [usage](#usage)
- [ref](#ref)

---

## mymodule1.c

```c
// test/mymodule1.c
#include <linux/module.h>
#include <linux/init.h>

int GLOBAL_VARIABLE = 1000;

EXPORT_SYMBOL(GLOBAL_VARIABLE);

/*
 * Function to print hello for num times.
 */
void print_hello(int num)
{
	while (num--) {
		printk(KERN_INFO "Hello Friend!!!\n");
	}
}
EXPORT_SYMBOL(print_hello);

/*
 * Function to add two passed number.
 */
void add_two_numbers(int a, int b)
{
	printk(KERN_INFO "Sum of the numbers %d", a + b);
}

EXPORT_SYMBOL(add_two_numbers);

static int __init my_init(void)
{
	printk(KERN_INFO "Hello from Export Symbol 1 module.");
	return 0;
}

static void __exit my_exit(void)
{
	printk(KERN_INFO "Bye from Export Symbol 1 module.");
}

module_init(my_init);
module_exit(my_exit);

MODULE_DESCRIPTION("Module to demonstrate the EXPORT_SYMBOL functionality");
MODULE_AUTHOR("Rishi Agrawal <rishi.b.agrawal@gmail.com");
MODULE_LICENSE("GPL v2");
```

---

## mymodule2.c

```c
// test/mymodule2.c
#include <linux/module.h>
#include <linux/init.h>

extern void print_hello(int);
extern void add_two_numbers(int, int);
extern int GLOBAL_VARIABLE;

/*
 * The function has been written just to call the functions which are in other module. This way you can also write modules which does provide some functionality to the other modules.
 */
static int __init my_init(void)
{
    printk(KERN_INFO "Hello from Hello Module");
    print_hello(10);
    add_two_numbers(5, 6);
    printk(KERN_INFO "Value of GLOBAL_VARIABLE %d", GLOBAL_VARIABLE);
    return 0;
}

static void __exit my_exit(void)
{
    printk(KERN_INFO "Bye from Hello Module");
}

module_init(my_init);
module_exit(my_exit);

MODULE_DESCRIPTION("Module to demonstrate the EXPORT_SYMBOL functionality");
MODULE_AUTHOR("Rishi Agrawal <rishi.b.agrawal@gmail.com>");
MODULE_LICENSE("GPL v2");
```

---

## makefile

```Makefile
# test/Makefile
obj-m +=  mymodule1.o
obj-m +=  mymodule2.o

PWD := $(CURDIR)
KDIR := /lib/modules/$(shell uname -r)/build

allofit:  modules
modules:
	@$(MAKE) -C $(KDIR) M=$(PWD) modules
modules_install:
	@$(MAKE) -C $(KDIR) M=$(PWD) modules_install
kernel_clean:
	@$(MAKE) -C $(KDIR) M=$(PWD) clean

clean: kernel_clean
	rm -rf   Module.symvers modules.order
```

---

## usage

```bash
linux:~/test/ # make
linux:~/test/ # insmod mymodule1.ko
linux:~/test/ # insmod mymodule2.ko

linux:~/test/ # dmesg
linux:~/test/ # modinfo mymodule1.ko
linux:~/test/ # modinfo mymodule2.ko

linux:~/test/ # rmmod mymodule2
linux:~/test/ # rmmod mymodule1
```

---

## ref

[Linux Kernel Workbook](https://lkw.readthedocs.io/en/latest/doc/04_exporting_symbols.html)
