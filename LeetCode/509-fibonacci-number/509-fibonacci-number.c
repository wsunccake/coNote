#include <stdio.h>

int fib(int n)
{
    int f[31] = {0};
    f[0] = 0;
    f[1] = 1;
    for (int i = 2; i <= n; i++)
        f[i] = f[i - 2] + f[i - 1];
    return f[n];
}

struct quest
{
    int inp;
    int sol;
};
typedef struct quest Quest;

int checkAns(Quest q)
{
    int out = fib(q.inp);
    int res = out == q.sol;

    if (!res)
    {
        printf("%d -> %d != %d\n", q.inp, q.sol, out);
    }

    return res;
}

int main()
{
    Quest q;

    q = (Quest){2, 1};
    checkAns(q);

    q = (Quest){3, 2};
    checkAns(q);

    q = (Quest){4, 3};
    checkAns(q);

    q = (Quest){10, 55};
    checkAns(q);

    return 0;
}