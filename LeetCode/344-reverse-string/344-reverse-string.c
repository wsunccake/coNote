#include <stdio.h>
#include <stdlib.h>

void reverseString(char *s, int sSize)
{
    int h = (sSize + 1) / 2;
    char t;
    int j;
    for (int i = 0; i < h; i++)
    {
        j = sSize - 1 - i;
        t = s[i];
        s[i] = s[j];
        s[j] = t;
    }
}

struct quest
{
    char *inp;
    int size;
    char *sol;
};
typedef struct quest Quest;

int checkAns(Quest q)
{
    reverseString(q.inp, q.size);
    for (int i = 0; i < q.size; i++)
    {

        if (q.inp[i] != q.sol[i])
        {
            printf("%s != %s\n", q.inp, q.sol);
            return 0;
        }
    }
    return 1;
}

int main()
{
    Quest q;
    char *s1 = (char *)malloc(sizeof("hello"));
    s1[0] = 'h';
    s1[1] = 'e';
    s1[2] = 'l';
    s1[3] = 'l';
    s1[4] = 'o';
    char *s2 = (char *)malloc(sizeof("olleh"));
    s2[0] = 'o';
    s2[1] = 'l';
    s2[2] = 'l';
    s2[3] = 'e';
    s2[4] = 'h';
    q = (Quest){s1, 5, s2};
    checkAns(q);
    free(s1);
    free(s2);

    return 0;
}