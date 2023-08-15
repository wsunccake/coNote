#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

#define show_err(a1, a2, b1, b2)          \
    if (((a1) != (b1)) || ((a2) != (b2))) \
        printf("ans1: %d != *a: %d, ans2: %d != *a+1: %d\n", (a1), (b1), (a2), (b2));

// int *twoSum(int *nums, int numsSize, int target, int *returnSize)
// {
//     int *r = (int *)malloc(sizeof(int) * 2);
//     int flag = 0;

//     for (int i = 0; i < numsSize - 1; i++)
//     {
//         for (int j = i + 1; j < numsSize; j++)
//         {
//             if (nums[i] + nums[j] == target)
//             {
//                 *r = i;
//                 *(r + 1) = j;
//                 flag = 1;
//                 break;
//             }
//         }

//         if (flag & 1)
//         {
//             break;
//         }
//     }

//     return r;
// }

int *twoSum(int *nums, int numsSize, int target, int *returnSize)
{
    int min_value = INT_MAX;
    int max_value = INT_MIN;

    for (int i = 0; i < numsSize; i++)
    {
        if (min_value > nums[i])
            min_value = nums[i];

        if (max_value < nums[i])
            max_value = nums[i];
    }

    int h[max_value - min_value + 1];
    for (int i = 0; i < max_value - min_value + 1; i++)
        h[i] = -1;

    int *r = (int *)malloc(sizeof(int) * 2);
    *r = 0, *(r + 1) = 0;
    int goal, idx;

    for (int i = 0; i < numsSize; i++)
    {
        goal = target - nums[i];

        if ((goal > max_value) || (goal < min_value))
            continue;
        if (h[goal - min_value] == -1)
        {
            h[nums[i] - min_value] = i;
            continue;
        }
        else
        {
            *r = h[goal - min_value];
            *(r + 1) = i;
            break;
        }
    }

    return r;
}

int main()
{
    int q[3];
    int target, ans1, ans2;
    int *a;

    q[0] = 1, q[1] = 2, q[2] = 3;
    target = 6, ans1 = 0, ans2 = 0;
    a = twoSum(q, 3, target, 2);
    show_err(ans1, ans2, *a, *(a + 1));
    free(a);

    q[0] = 3, q[1] = 2, q[2] = 4;
    target = 6, ans1 = 1, ans2 = 2;
    a = twoSum(q, 3, target, 2);
    show_err(ans1, ans2, *a, *(a + 1));
    free(a);

    q[0] = 3, q[1] = 3, q[2] = 0;
    target = 6, ans1 = 0, ans2 = 1;
    a = twoSum(q, 3, target, 2);
    show_err(ans1, ans2, *a, *(a + 1));
    free(a);

    return 0;
}