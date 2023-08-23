/*
	kzalloc()/kmalloc()
	kfree()
	cdev_init()
	cdev_add()
	cdev_del()
	struct file_operations
*/
#include <linux/module.h>
#include <linux/fs.h>
#include <linux/cdev.h>
#include <linux/slab.h>

#define HELLO_MAJOR 0
#define HELLO_NR_DEVS 2
#define DEV_NAME "hello_chr"

int hello_major = HELLO_MAJOR;
int hello_minor = 0;

dev_t devt;

int hello_nr_devs = HELLO_NR_DEVS;

module_param(hello_major, int, S_IRUGO);
module_param(hello_minor, int, S_IRUGO);
module_param(hello_nr_devs, int, S_IRUGO);

struct hello_char_dev
{
	struct cdev cdev;
	char c;
};

struct hello_char_dev *hc_devp;

int hc_open(struct inode *inode, struct file *filp)
{
	printk(KERN_INFO "open hc_dev%d %d\n", iminor(inode), MINOR(inode->i_cdev->dev));
	return 0;
}
ssize_t hc_read(struct file *filp, char __user *buf, size_t count, loff_t *f_pos)
{
	printk(KERN_INFO "read hc_dev\n");
	return 0;
}
ssize_t hc_write(struct file *filp, const char __user *buf, size_t count, loff_t *f_pos)
{
	printk(KERN_INFO "write %s\n", DEV_NAME);
	return count; // if count = 0, always run
}

int hc_release(struct inode *inode, struct file *filp)
{
	printk(KERN_INFO "release %s\n", DEV_NAME);
	return 0;
}

struct file_operations hc_fops = {
	// character device file operator
	.owner = THIS_MODULE,
	.read = hc_read,
	.write = hc_write,
	.open = hc_open,
	.release = hc_release,
};

static int __init hello_init(void)
{
	int ret, i;
	printk(KERN_INFO "---init %s module---\n", DEV_NAME);
	if (hello_major)
	{
		devt = MKDEV(hello_major, hello_minor);
		ret = register_chrdev_region(devt, hello_nr_devs, "hello_chr");
	}
	else
	{
		ret = alloc_chrdev_region(&devt, hello_minor, hello_nr_devs, "hello_chr");
		hello_major = MAJOR(devt);
	}
	if (ret < 0)
	{
		printk(KERN_WARNING "hello: can't get major %d\n", hello_major);
		goto fail;
	}

	hc_devp = kzalloc(sizeof(struct hello_char_dev) * hello_nr_devs, GFP_KERNEL);
	if (!hc_devp)
	{
		printk(KERN_WARNING "alloc mem failed");
		ret = -ENOMEM;
		goto failure_kzalloc; // when fail
	}

	for (i = 0; i < hello_nr_devs; i++)
	{
		cdev_init(&hc_devp[i].cdev, &hc_fops);
		hc_devp[i].cdev.owner = THIS_MODULE;
		ret = cdev_add(&hc_devp[i].cdev, MKDEV(hello_major, hello_minor + i), 1);
		if (ret)
		{
			printk(KERN_WARNING "fail add %s%d", DEV_NAME, i);
		}
	}

	printk(KERN_INFO "---end %s module---\n", DEV_NAME);
	return 0;

failure_kzalloc:
	unregister_chrdev_region(devt, hello_nr_devs);
fail:
	return ret;
}

static void __exit hello_exit(void)
{
	int i;
	for (i = 0; i < hello_nr_devs; i++)
		cdev_del(&hc_devp[i].cdev);
	kfree(hc_devp);
	unregister_chrdev_region(devt, hello_nr_devs);
	printk(KERN_INFO "---exit %s module---\n", DEV_NAME);
}

module_init(hello_init);
module_exit(hello_exit);

MODULE_LICENSE("Dual BSD/GPL");
MODULE_AUTHOR("KGZ");
MODULE_VERSION("V1.0");
