# compile c program by makefile

---

## content

- [test.c](#testc)
- [makefile](#makefile)
- [run](#run)

---

## test.c

```c
// test.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define PAGE_SIZE 4096
#define MAX_SIZE 100*PAGE_SIZE

int main()
{
    char *buf = (char *)malloc(MAX_SIZE);

    memset(buf, 0, MAX_SIZE);

    printf("buffer address=0x%p\n", buf);

    free(buf);
           return 0;
}
```

---

## makefile

```makefile
# Makefile
cc = aarch64-linux-gnu-gcc
prom = test
obj = test.o
CFLAGS = -static

$(prom): $(obj)
    $(cc) -o $(prom) $(obj) $(CFLAGS)

%.o: %.c
    $(cc) -c $< -o $@

clean:
    rm -rf $(obj) $(prom)
```

---

## run

```bash
build:~ # make
build:~ # ./test

build:~ # make clean
```
