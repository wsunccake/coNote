#include <linux/module.h>

// export
char *hi = "hi, linux\n";
EXPORT_SYMBOL(hi);

void prt(void)
{
	printk(KERN_INFO "this is hi module\n");
}
EXPORT_SYMBOL(prt);

static int __init hi_init(void)
{
	printk(KERN_INFO "init hi\n");
	return 0;
}

static void __exit hi_exit(void)
{
	printk(KERN_INFO "exit hi\n");
}

module_init(hi_init);
module_exit(hi_exit);

MODULE_LICENSE("Dual BSD/GPL");
MODULE_AUTHOR("KGZ");
MODULE_VERSION("V1.0");
