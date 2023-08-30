#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <limits.h>

#define show_err(a1, a2, b1, b2)          \
    if (((a1) != (b1)) || ((a2) != (b2))) \
        printf("ans1: %d != *a: %d, ans2: %d != *a+1: %d\n", (a1), (b1), (a2), (b2));

// --------------

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
//                 *returnSize = 2;
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

// --------------

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

    int *h = (int *)calloc(max_value - min_value + 1, sizeof(int));
    for (int i = 0; i < numsSize; i++)
        h[nums[i] - min_value] = i;
    h[nums[0] - min_value] = -1;

    int *r = (int *)malloc(sizeof(int) * 2);
    *r = 0, *(r + 1) = 0;
    int goal;

    for (int i = 1; i < numsSize; i++)
    {
        goal = target - nums[i];

        if ((goal > max_value) || (goal < min_value))
            continue;

        if (h[goal - min_value] != 0)
        {
            printf("i: %d, goal: %d, %d\n", i, goal, h[goal - min_value]);
            if (i == h[goal - min_value])
                continue;
            *r = h[goal - min_value];
            *(r + 1) = i;
            *returnSize = 2;
            break;
        }
    }

    if (*returnSize == 2)
    {
        if (*r > *(r + 1))
        {
            *r = *r ^ *(r + 1);
            *(r + 1) = *r ^ *(r + 1);
            *r = *r ^ *(r + 1);
        }
    }

    if (*r == -1)
        *r = 0;

    free(h);
    return r;
}

// --------------

typedef struct
{
    int value;
    int index;
} Element;

int compare(const void *a, const void *b)
{
    return ((Element *)a)->value - ((Element *)b)->value;
}

// int *twoSum(int *nums, int numsSize, int target, int *returnSize)
// {
//     *returnSize = 0;
//     Element *elements = malloc(numsSize * sizeof(Element));

//     for (int i = 0; i < numsSize; ++i)
//     {
//         elements[i].value = nums[i];
//         elements[i].index = i;
//     }

//     qsort(elements, numsSize, sizeof(Element), compare);

//     int left = 0;
//     int right = numsSize - 1;
//     int *result = malloc(2 * sizeof(int));

//     while (left < right)
//     {
//         int sum = elements[left].value + elements[right].value;

//         if (sum == target)
//         {
//             result[0] = elements[left].index;
//             result[1] = elements[right].index;
//             *returnSize = 2;
//             break;
//         }
//         else if (sum < target)
//         {
//             ++left;
//         }
//         else
//         {
//             --right;
//         }
//     }

//     free(elements);
//     return result;
// }

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