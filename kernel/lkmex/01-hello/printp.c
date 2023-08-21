#include <linux/module.h>

extern char *hi;
extern void prt(void);

static int __init printp_init(void)
{
	printk(KERN_INFO "printp:%s", hi);
	prt();
	return 0;
}

static void __exit printp_exit(void)
{
}

module_init(printp_init);
module_exit(printp_exit);

MODULE_LICENSE("Dual BSD/GPL");
MODULE_AUTHOR("KGZ");
MODULE_VERSION("V1.0");
