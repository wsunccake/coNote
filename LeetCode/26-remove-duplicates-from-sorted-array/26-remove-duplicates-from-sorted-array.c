#include <stdio.h>

int removeDuplicates(int *nums, int numsSize)
{
    int current = nums[0];
    int i = 1, j = 1;
    while (i < numsSize)
    {
        if (current != nums[i])
        {
            nums[j] = nums[i];
            j += 1;
            current = nums[i];
        }
        i += 1;
    }
    return j;
}

struct quest
{
    int *inp;
    int size;
    int *sol;
    int out;
};
typedef struct quest Quest;

int checkAns(Quest q)
{
    int out = removeDuplicates(q.inp, q.size);

    if (out != q.out)
    {
        printf("out != q.out, %d != %d\n", out, q.out);
        return 1;
    }

    for (int i = 0; i < out; i++)
    {

        if (q.inp[i] != q.sol[i])
        {
            printf("q.inp != q.sol, %d != %d\n", q.inp[i], q.sol[i]);
            return 0;
        }
    }
    return 1;
}

int main()
{
    Quest q;
    q = (Quest){(int[]){1, 1, 2}, 3, (int[]){1, 2}, 2};
    checkAns(q);

    return 0;
}