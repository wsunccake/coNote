#include <stdio.h>

int main()
{
    int a = 5, b = 7, c;

    __asm__ __volatile__(
        "movl %1, %%eax;\n\t"
        "addl %2, %%eax;\n\t"
        "movl %%eax, %0;"
        : "=r"(c)
        : "r"(a), "r"(b)
        : "%eax");

    printf("a + b = %d\n", c);
    return 0;
}