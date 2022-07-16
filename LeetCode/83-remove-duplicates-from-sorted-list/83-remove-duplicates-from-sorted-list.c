#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define STRING_LEN 100
typedef char *string;

void swap(char *a, char *b)
{
    *a = *a ^ *b;
    *b = *a ^ *b;
    *a = *a ^ *b;
}

char *convertToTitle(int columnNumber)
{
    int remainder;
    // char *s = (char *)malloc(sizeof(char) * 100);
    char *s = (char *)calloc(100, sizeof(char));
    int i = 0;

    while (columnNumber > 0)
    {
        remainder = columnNumber % 26;
        if (remainder == 0)
        {
            s[i] = 'Z';
        }
        else
        {
            s[i] = remainder + 'A' - 1;
        }
        i++;
        columnNumber = (columnNumber - 1) / 26;
    }

    for (int j = 0; j < i / 2; j++)
    {
        s[j] = s[j] ^ s[i - j - 1];
        s[i - j - 1] = s[j] ^ s[i - j - 1];
        s[j] = s[j] ^ s[i - j - 1];

        // swap(&s[j], &s[i - j - 1]);
    }

    return s;
}

struct quest
{
    int inp;
    string sol;
};
typedef struct quest Quest;

int checkAns(Quest q)
{
    char *out = convertToTitle(q.inp);
    int res = strcmp(out, q.sol) == 0;

    if (!res)
    {
        printf("%d -> %s != %s\n", q.inp, q.sol, out);
    }

    free(out);
    return res;
}

int main()
{
    Quest q1 = {1, "A"};
    checkAns(q1);

    Quest q2 = {26, "Z"};
    checkAns(q2);

    Quest q3 = {27, "AA"};
    checkAns(q3);

    Quest q4 = {701, "ZY"};
    checkAns(q4);

    return 0;
}
