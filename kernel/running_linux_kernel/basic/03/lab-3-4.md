# c inline assembly

## content

- [x86_64](#x86_64)
- [ref](#ref)

---

## x86_64

```c
// ex.c
#include <stdio.h>

int main()
{
    int sum, num1, num2;
    num1 = 1;
    num2 = 2;
    // sum = num1 + num2;
    asm(
        "addl    %%edx, %%eax\n"
        :"=a"(sum)
        :"a"(num1), "d"(num2)
       );
    printf("sum=%d\r\n", sum);
    return 0;
}
```

```bash
linux:~ # gcc -o ex ex.c
linux:~ # ./ex
```

---

## ref

[C 語言的行內組譯](https://evshary.com/2018/05/20/C-Inline-Assembly/)
