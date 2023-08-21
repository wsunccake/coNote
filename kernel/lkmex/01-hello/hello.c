#include <linux/module.h>

static int __init hello_init(void)
{
	printk(KERN_INFO "Hello Linux Kernel Module\n");
	return 0;
}

static void __exit hello_exit(void)
{
	printk(KERN_INFO "Bye Linux Kernel Module\n");
}

module_init(hello_init);
module_exit(hello_exit);

MODULE_LICENSE("Dual BSD/GPL");
MODULE_AUTHOR("KGZ");
MODULE_VERSION("V1.0");
