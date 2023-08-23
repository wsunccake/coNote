/*
*  	register_chrdev_region()
*	alloc_chrdev_region()
*	unregister_chrdev_region()
	MAJOR()
	MINOR()
	MKDEV()
*/
#include <linux/module.h>

// #include<linux/types.h>	//dev_t
// #include<linux/kdev_t.h>	//MAJOR/MINOR/MKDEV
#include <linux/fs.h>

#define DEMO_MAJOR 0
#define DEMO_NR_DEVS 2

int demo_major = DEMO_MAJOR;
int demo_minor = 0;

// major: high 12 bit
// minor: low 20 bit
dev_t demo_dev;

int demo_nr_devs = DEMO_NR_DEVS;

module_param(demo_major, int, S_IRUGO);
module_param(demo_minor, int, S_IRUGO);
module_param(demo_nr_devs, int, S_IRUGO);

static int __init demo_init(void)
{
	int ret;
	printk(KERN_INFO "---init chr module---\n");
	// register device
	// 0: dynamic allocate
	// not 0: assign
	if (demo_major)
	{
		demo_dev = MKDEV(demo_major, demo_minor);
		ret = register_chrdev_region(demo_dev, demo_nr_devs, "demo_chr");
	}
	else
	{
		ret = alloc_chrdev_region(&demo_dev, demo_minor, demo_nr_devs, "demo_chr");
		demo_major = MAJOR(demo_dev);
	}
	if (ret < 0)
	{
		printk(KERN_WARNING "demo: can't get major %d\n", demo_major);
		return ret;
	}
	printk(KERN_INFO "demo_chr:%d demo_dev:%x", demo_major, demo_dev);
	printk(KERN_INFO "---end chr module---\n");
	return 0;
}

static void __exit demo_exit(void)
{
	// unregister device
	unregister_chrdev_region(demo_dev, demo_nr_devs);
	printk(KERN_INFO "---exit chr module---\n");
}

module_init(demo_init);
module_exit(demo_exit);

MODULE_LICENSE("Dual BSD/GPL");
MODULE_AUTHOR("KGZ");
MODULE_VERSION("V1.0");
