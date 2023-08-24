/*
	sema_init()
	down_interruptible()
	up()
	mutex_init()
	mutex_lock_interruptible()
	mutex_unlock()

*/
#include <linux/module.h>
#include <linux/fs.h>
#include <linux/cdev.h>
#include <linux/slab.h>
#include <linux/uaccess.h>
#include <linux/jiffies.h>
#include <linux/sched.h>

#include <linux/semaphore.h>
#include <linux/mutex.h>

#ifndef LOCK_USE
#define LOCK_USE 1 // 0:semaphore,1:mutex
#endif

#define HELLO_MAJOR 0
#define HELLO_NR_DEVS 2
#define DEV_NAME "hc_op_locked"

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
	struct semaphore sema;
	struct mutex mtx;
};

struct hello_char_dev *hc_devp;
struct class *hc_cls;

int hc_open(struct inode *inode, struct file *filp)
{
	struct hello_char_dev *hc_dev;
	printk(KERN_INFO "%s %s open \n", DEV_NAME, current->comm);
	hc_dev = container_of(inode->i_cdev, struct hello_char_dev, cdev);
	filp->private_data = hc_dev;

	return 0;
}
ssize_t hc_read(struct file *filp, char __user *buf, size_t count, loff_t *f_pos)
{
	ssize_t retval = 0;
	struct hello_char_dev *hc_dev = filp->private_data;
	printk(KERN_INFO "%s read %p\n", DEV_NAME, hc_dev);

	if (*f_pos >= hc_dev->n)
		goto out;
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
	int retval = -ENOMEM;
	unsigned long jiff1;
	printk(KERN_INFO "%s write begin\n", current->comm);
#if (LOCK_USE == 0)
	if (down_interruptible(&hc_dev->sema)) //-EINTR
		return -ERESTARTSYS;
	printk(KERN_INFO "%s get sema\n", current->comm);
#endif
#if (LOCK_USE == 1)
	if (mutex_lock_interruptible(&hc_dev->mtx)) //-EINTR
		return -ERESTARTSYS;
	printk(KERN_INFO "%s get mutex\n", current->comm);
#endif

	kfree(hc_dev->c);
	hc_dev->c = NULL;
	hc_dev->n = 0;
	printk(KERN_INFO "%s 1", current->comm);
	jiff1 = jiffies;
	while (jiffies - jiff1 < HZ)
		;
	hc_dev->c = kzalloc(count, GFP_KERNEL);
	if (!hc_dev->c)
		goto out;
	printk(KERN_INFO "%s 2 addr:%p", current->comm, hc_dev->c);
	jiff1 = jiffies;
	while (jiffies - jiff1 < HZ)
		;
	printk(KERN_INFO "%s %s 3 addr:%p", DEV_NAME, current->comm, hc_dev->c);
	if (copy_from_user(hc_dev->c, buf, count))
	{
		retval = -EFAULT;
		goto fail_copy;
	}
	hc_dev->n = count;

#if (LOCK_USE == 0)
	up(&hc_dev->sema);
	printk(KERN_INFO "%s up sema\n", current->comm);
#endif
#if (LOCK_USE == 1)
	mutex_unlock(&hc_dev->mtx);
	printk(KERN_INFO "%s unlock mutex\n", current->comm);
#endif
	printk(KERN_INFO "%s %s write done", DEV_NAME, current->comm);
	return count;
fail_copy:
	kfree(hc_dev->c);
out:
#if (LOCK_USE == 0)
	up(&hc_dev->sema);
#endif
#if (LOCK_USE == 1)
	mutex_unlock(&hc_dev->mtx);
#endif
	return retval;
}
int hc_release(struct inode *inode, struct file *filp)
{
	printk(KERN_INFO "%s %s release\n", DEV_NAME, current->comm);
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
#if (LOCK_USE == 0)
		sema_init(&hc_devp[i].sema, 1);
#elif (LOCK_USE == 1)
		mutex_init(&hc_devp[i].mtx);
#endif

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
