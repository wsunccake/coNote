#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *f(int x, int y)
{
    char *s = malloc(sizeof(char) * 25);

    if (x == y)
        sprintf(s, "%d", x);
    else
        sprintf(s, "%d->%d", x, y);

    return s;
}

char **summaryRanges(int *nums, int numsSize, int *returnSize)
{
    char **ret = (char **)malloc(sizeof(char *) * numsSize);
    *returnSize = 0;

    if (numsSize == 0)
        return ret;

    int beginNum = nums[0];
    int endNum = nums[0];

    for (int i = 0; i < numsSize; i++)
    {
        if (endNum + 1 < nums[i])
        {
            ret[*returnSize] = f(beginNum, endNum);
            beginNum = nums[i];
            *returnSize = *returnSize + 1;
        }
        endNum = nums[i];
    }
    ret[*returnSize] = f(beginNum, endNum);

    (*returnSize)++;
    return ret;
}

void checkAns(char **sol, int sol_size, char **ans, int ans_size)
{
    if (sol_size != ans_size)
        printf("fail: size -> %d != %d\n", sol_size, ans_size);
    for (size_t i = 0; i < sol_size; i++)
    {
        if (strcmp(sol[i], *(ans + i)) != 0)
            printf("fail: str -> %s != %s\n", sol[i], *(ans + i));
    }
}

int main()
{
    int num1[] = {0, 1, 2, 4, 5, 7};
    char *sol1[] = {"0->2", "4->5", "7"};
    int sol1_size = 3;
    int ans1_size = 0;
    char **ans1 = summaryRanges(num1, 6, &ans1_size);
    checkAns(sol1, sol1_size, ans1, ans1_size);

    int num2[] = {0, 2, 3, 4, 6, 8, 9};
    char *sol2[] = {"0", "2->4", "6", "8->9"};
    int sol2_size = 4;
    int ans2_size = 0;
    char **ans2 = summaryRanges(num2, 7, &ans2_size);
    checkAns(sol2, sol2_size, ans2, ans2_size);

    int num3[] = {0};
    char *sol3[] = {"0"};
    int sol3_size = 1;
    int ans3_size = 0;
    char **ans3 = summaryRanges(num3, 1, &ans3_size);
    checkAns(sol3, sol3_size, ans3, ans3_size);
    return 0;

    int num4[] = {};
    char *sol4[] = {};
    int sol4_size = 0;
    int ans4_size = 0;
    char **ans4 = summaryRanges(num4, 0, &ans4_size);
    checkAns(sol4, sol4_size, ans4, ans4_size);
    return 0;
}