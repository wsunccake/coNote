#include <linux/module.h>

#define CNT 1

// parameter
static int cnt = CNT;
static char *hi = "hey, linux\n";
static int arr[] = {1, 2, 3, 4, 5, 6};
static int nums = sizeof(arr) / sizeof(int);

module_param(cnt, int, S_IRUGO);
MODULE_PARM_DESC(cnt, "count, int");
module_param(hi, charp, S_IRUGO);
MODULE_PARM_DESC(hi, "hi, char");
module_param_array(arr, int, &nums, S_IRUGO);
MODULE_PARM_DESC(arr, "arr: int[]");

static int __init hey_init(void)
{
    printk(KERN_INFO "init hey\n");

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

static void __exit hey_exit(void)
{
    printk(KERN_INFO "exit hey\n");
}

module_init(hey_init);
module_exit(hey_exit);

MODULE_LICENSE("Dual BSD/GPL");
MODULE_AUTHOR("KGZ");
MODULE_VERSION("V1.0");
