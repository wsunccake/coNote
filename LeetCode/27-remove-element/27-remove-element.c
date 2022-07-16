#include <stdio.h>

int removeElement(int *nums, int numsSize, int val)
{
    int i = 0, j = 0;
    while (i + j < numsSize)
    {
        if (nums[i] == val)
        {
            if (nums[i + j] == val)
            {
                j++;
            }
            else
            {
                nums[i] = nums[i] ^ nums[i + j];
                nums[i + j] = nums[i] ^ nums[i + j];
                nums[i] = nums[i] ^ nums[i + j];
                i++;
                j = 0;
            }
        }
        else
        {
            i++;
        }
    }

    return i;
}

struct quest
{
    int *inpNums;
    int inpSize;
    int inpVal;
    int *solNums;
    int solVal;
};
typedef struct quest Quest;

int checkAns(Quest q)
{
    int out = removeElement(q.inpNums, q.inpSize, q.inpVal);
    int res = out == q.solVal;

    if (!res)
    {
        printf("%p -> %d != %d\n", q.inpNums, q.solVal, out);
    }

    return res;
}

int main()
{
    Quest q;

    // q = (Quest){(int[]){3, 2, 2, 3}, 4, 3, (int[]){2, 2, 0, 0}, 2};
    // checkAns(q);

    q = (Quest){(int[]){0, 1, 2, 2, 3, 0, 4, 2}, 8, 2, (int[]){0, 1, 4, 0, 3}, 5};
    checkAns(q);

    return 0;
}