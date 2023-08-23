/*
 *	register_blkdev()
 *	unregister_blkdev()
 */

#include <linux/module.h>

// #include<linux/types.h>	//dev_t
// #include<linux/kdev_t.h>	//MAJOR/MINOR/MKDEV
#include <linux/fs.h>
#include <linux/blkdev.h>

#define DEMO_MAJOR 0

int demo_major = DEMO_MAJOR;

module_param(demo_major, int, S_IRUGO);

static int __init demo_init(void)
{
	int ret;
	printk(KERN_INFO "---init blk moudule---\n");
	ret = register_blkdev(demo_major, "hello_blk");
	if (ret < 0)
	{
		printk(KERN_WARNING "demo_blk: can't get major %d\n", demo_major);
		return ret;
	}
	if (!demo_major)
	{
		demo_major = ret;
	}
	printk(KERN_INFO "demo_blk:%d ret:%d", demo_major, ret);
	printk(KERN_INFO "---end blk moudle---\n");
	return 0;
}

static void __exit demo_exit(void)
{
	unregister_blkdev(demo_major, "dem_blk");
	printk(KERN_INFO "---exit blk moudle---\\n");
}

module_init(demo_init);
module_exit(demo_exit);

MODULE_LICENSE("Dual BSD/GPL");
MODULE_AUTHOR("KGZ");
MODULE_VERSION("V1.0");
