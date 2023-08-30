/*
	container_of()
	copy_to_user()
	copy_from_user()
	struct inode
	struct file
*/
#include <linux/module.h>
#include <linux/fs.h>
#include <linux/cdev.h>
#include <linux/slab.h>

#include <linux/uaccess.h>

#define HELLO_MAJOR 0
#define HELLO_NR_DEVS 2
#define DEV_NAME "hc_op"

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
	char *c;
	int n;
};

struct hello_char_dev *hc_devp;
struct class *hc_cls;

int hc_open(struct inode *inode, struct file *filp)
{
	struct hello_char_dev *hc_dev;
	printk(KERN_INFO "open %s%d %d\n", DEV_NAME, iminor(inode), MINOR(inode->i_cdev->dev));
	hc_dev = container_of(inode->i_cdev, struct hello_char_dev, cdev);
	filp->private_data = hc_dev;

	return 0;
}
ssize_t hc_read(struct file *filp, char __user *buf, size_t count, loff_t *f_pos)
{
	ssize_t retval = 0;
	struct hello_char_dev *hc_dev = filp->private_data;
	printk(KERN_INFO "read %s %p\n", DEV_NAME, hc_dev);

	// hc_dev->n - *f_pos <= 0, no space
	if (*f_pos >= hc_dev->n)
		goto out;

	// count > hc_dev->n - *f_pos, no enough space
	if (*f_pos + count > hc_dev->n)
		count = hc_dev->n - *f_pos;

	if (copy_to_user(buf, hc_dev->c, count))
	{
		retval = -EFAULT;
		goto out;
	}

	*f_pos += count;
	return count;
out:
	return retval;
}
ssize_t hc_write(struct file *filp, const char __user *buf, size_t count, loff_t *f_pos)
{
	struct hello_char_dev *hc_dev = filp->private_data;
	int retval = -ENOMEM, i;
	printk(KERN_INFO "write %s %p\n", DEV_NAME, hc_dev);
	kfree(hc_dev->c);
	hc_dev->c = NULL;
	hc_dev->n = 0;
	hc_dev->c = kzalloc(count, GFP_KERNEL);
	printk(KERN_INFO "addr %p", hc_dev->c);
	if (!hc_dev->c)
		goto out;
	if (copy_from_user(hc_dev->c, buf, count))
	{
		retval = -EFAULT;
		goto fail_copy;
	}
	for (i = 0; i < count; i++)
		printk(KERN_INFO "get %c", hc_dev->c[i]);
	hc_dev->n = count;

	return count;
fail_copy:
	kfree(hc_dev->c);
out:
	return retval;
}

int hc_release(struct inode *inode, struct file *filp)
{
	printk(KERN_INFO "release %s\n", DEV_NAME);
	return 0;
}

struct file_operations hc_fops = {
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
		ret = register_chrdev_region(devt, hello_nr_devs, DEV_NAME);
	}
	else
	{
		ret = alloc_chrdev_region(&devt, hello_minor, hello_nr_devs, DEV_NAME);
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
		goto failure_kzalloc;
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

	hc_cls = class_create(THIS_MODULE, DEV_NAME);
	if (!hc_cls)
	{
		printk(KERN_WARNING "fail create class");
		ret = PTR_ERR(hc_cls);
		goto failure_class;
	}
	for (i = 0; i < hello_nr_devs; i++)
	{
		device_create(hc_cls, NULL, MKDEV(hello_major, hello_minor + i), NULL, "%s%d", DEV_NAME, i);
	}
	printk(KERN_INFO "---end %s module---\n", DEV_NAME);
	return 0;

failure_class:
	kfree(hc_devp);
failure_kzalloc:
	unregister_chrdev_region(devt, hello_nr_devs);
fail:
	return ret;
}

static void __exit hello_exit(void)
{
	int i;
	for (i = 0; i < hello_nr_devs; i++)
	{
		device_destroy(hc_cls, MKDEV(hello_major, hello_minor + i));
	}
	class_destroy(hc_cls);
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
