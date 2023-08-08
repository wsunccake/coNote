# simple char device

---

## content

- [simple_char.c](#simple_charc)
- [makefile](#makefile)
- [usage](#usage)

---

## simple_char.c

```c
// simple_char.c
#include <linux/module.h>
#include <linux/fs.h>
#include <linux/uaccess.h>
#include <linux/init.h>
#include <linux/cdev.h>

#define DEMO_NAME "my_demo_dev"

static dev_t dev;
static struct cdev *demo_cdev;
static signed count = 1;
static struct class *demo_class = NULL;
static struct device *demo_device = NULL;

static int demodrv_open(struct inode *inode, struct file *file)
{
    int major = MAJOR(inode->i_rdev);
    int minor = MINOR(inode->i_rdev);

    printk("%s: major=%d, minor=%d\n", __func__, major, minor);

    return 0;
}

static int demodrv_release(struct inode *inode, struct file *file)
{
    return 0;
}

static ssize_t
demodrv_read(struct file *file, char __user *buf, size_t lbuf, loff_t *ppos)
{
    printk("%s enter\n", __func__);
    return 0;
}

static ssize_t
demodrv_write(struct file *file, const char __user *buf, size_t count, loff_t *f_pos)
{
    printk("%s enter\n", __func__);
    return 0;
}

static const struct file_operations demodrv_fops = {
    .owner = THIS_MODULE,
    .open = demodrv_open,
    .release = demodrv_release,
    .read = demodrv_read,
    .write = demodrv_write};

static int __init simple_char_init(void)
{
    int ret;
    // teardown: unregister_chrdev_region
    ret = alloc_chrdev_region(&dev, 0, count, DEMO_NAME);
    if (ret)
    {
        printk("fail to alloc_chrdev_region");
        return ret;
    }

    // teardown: class_destroy
    // -> /sys/class/
    demo_class = class_create(THIS_MODULE, DEMO_NAME);
    // if (!demo_class)                                   // demo_class == NULL
    // {
    //     printk("fail to class_create\n");
    //     goto class_create_fail;
    // }
    if (IS_ERR(demo_class))
    {
        printk("fail to class_create \n");
        goto class_create_fail;
        // return PTR_ERR(demo_class);
    }

    // teardown: device_destroy
    // -> /dev/
    demo_device = device_create(demo_class, NULL, dev, NULL, DEMO_NAME);
    if (IS_ERR(demo_device))
    {
        pr_err("fail to device_create\n");
        goto device_create_fail;
    }

    // teardown: cdev_del
    demo_cdev = cdev_alloc();
    if (!demo_cdev)
    {
        printk("fail to cdev_alloc\n");
        goto cdev_alloc_fail;
    }

    cdev_init(demo_cdev, &demodrv_fops);

    ret = cdev_add(demo_cdev, dev, count);
    if (ret)
    {
        printk("fail to cdev_add\n");
        goto cdev_add_fail;
    }

    printk("succeeded register char device: %s\n", DEMO_NAME);
    printk("Major number = %d, minor number = %d\n",
           MAJOR(dev), MINOR(dev));

    return 0;

cdev_add_fail:
    cdev_del(demo_cdev);
cdev_alloc_fail:
    device_destroy(demo_class, dev);
device_create_fail:
    class_destroy(demo_class);
class_create_fail:
    unregister_chrdev_region(dev, count);

    return ret;
}

static void __exit simple_char_exit(void)
{
    printk("removing device\n");

    device_destroy(demo_class, dev);
    class_destroy(demo_class);

    if (demo_cdev)
        cdev_del(demo_cdev);

    unregister_chrdev_region(dev, count);
}

module_init(simple_char_init);
module_exit(simple_char_exit);

MODULE_AUTHOR("kylin");
MODULE_LICENSE("GPL v2");
MODULE_DESCRIPTION("simpe character device");
```

---

## makefile

```makefile
PWD := $(CURDIR)
KDIR ?= /lib/modules/$(shell uname -r)/build

mydemo-objs := simple_char.o
obj-m	:=   mydemo.o

all :
	$(MAKE) -C $(KDIR) M=$(PWD) modules;

clean:
	$(MAKE) -C $(KDIR) M=$(PWD) clean;
	rm -f *.ko;
```

---

## usage

```bash
linux:~ # make

linux:~ # insmod mydemo.ko

linux:~ # dmesg
linux:~ # grep my_demo_dev /proc/devices
linux:~ # ls /dev/my_demo_dev
linux:~ # ls /sys/class/my_demo_dev

linux:~ # rmmod mydemo
```
