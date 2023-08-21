#include <linux/module.h>

#define CNT 1

// export
char *hi = "Hi,Linux\n";
EXPORT_SYMBOL(hi);

void prt(void)
{
	printk(KERN_INFO "this is hello module\n");
}
EXPORT_SYMBOL(prt);

// parameter
static int cnt = CNT;
static int arr[] = {1, 2, 3, 4, 5, 6};
static int nums = sizeof(arr) / sizeof(int);

module_param(cnt, int, S_IRUGO);
MODULE_PARM_DESC(cnt, "count, int");
module_param(hi, charp, S_IRUGO);
MODULE_PARM_DESC(hi, "hi, char");
module_param_array(arr, int, &nums, S_IRUGO);
MODULE_PARM_DESC(arr, "arr: int[]");

static int __init hi_init(void)
{
	printk(KERN_INFO "Hello Linux Kernel Module\n");

	for (int i = 0; i < cnt; i++)
	{
		printk(KERN_INFO "%d:%s", i, hi);
	}
	for (int i = 0; i < 6; i++)
	{
		printk(KERN_INFO "%d ", arr[i]);
	}
	printk(KERN_INFO "nums:%d\n", nums);

	return 0;
}

static void __exit hi_exit(void)
{
	printk(KERN_INFO "Bye Linux Kernel Module\n");
}

module_init(hi_init);
module_exit(hi_exit);

MODULE_LICENSE("Dual BSD/GPL");
MODULE_AUTHOR("KGZ");
MODULE_VERSION("V1.0");
