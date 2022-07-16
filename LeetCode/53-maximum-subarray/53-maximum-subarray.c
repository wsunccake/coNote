#include <stdio.h>

int maxSubArray(int *nums, int numsSize)
{
    int max_sum = nums[0];
    int sum = 0;
    for (int i = 0; i < numsSize; i++)
    {
        if (sum > 0)
        {
            sum += nums[i];
        }
        else
        {
            sum = nums[i];
        }

        if (sum > max_sum)
        {
            max_sum = sum;
        }
    }
    return max_sum;
}

struct quest
{
    int *inp;
    int size;
    int sol;
};
typedef struct quest Quest;

int checkAns(Quest q)
{
    int out = maxSubArray(q.inp, q.size);
    int res = out == q.sol;

    if (!res)
    {
        printf("%p -> %d != %d\n", q.inp, q.sol, out);
    }

    return res;
}

int main()
{
    Quest q;

    q = (Quest){(int[]){-2, 1, -3, 4, -1, 2, 1, -5, 4}, 9, 6};
    checkAns(q);

    q = (Quest){(int[]){1}, 1, 1};
    checkAns(q);

    q = (Quest){(int[]){5, 4, -1, 7, 8}, 5, 23};
    checkAns(q);

    q = (Quest){(int[]){-10, -4, -1, -8}, 4, -1};
    checkAns(q);

    return 0;
}