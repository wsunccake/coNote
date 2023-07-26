# call assembly from c

## content

- [x86_64](#x86_64)
  - [asm](#asm)
  - [c](#c)
  - [makefile](#makefile)
  - [run](#run)
- [arm64](#arm64)
  - [asm](#asm-1)
  - [c](#c-1)
  - [makefile](#makefile-1)
  - [run](#run-1)
- [ref](#ref)

---

## x86_64

### asm

```s
# maxofthree.s
        .globl  maxofthree

        .text
maxofthree:
        mov     %rdi, %rax                # result (rax) initially holds x
        cmp     %rsi, %rax                # is x less than y?
        cmovl   %rsi, %rax                # if so, set result to y
        cmp     %rdx, %rax                # is max(x,y) less than z?
        cmovl   %rdx, %rax                # if so, set result to z
        ret                               # the max will be in eax
```

### c

```c
// callmaxofthree.c
#include <stdio.h>
#include <inttypes.h>

int64_t maxofthree(int64_t, int64_t, int64_t);

int main() {
    printf("%ld\n", maxofthree(1, -4, -7));
    printf("%ld\n", maxofthree(2, -6, 1));
    printf("%ld\n", maxofthree(2, 3, 1));
    printf("%ld\n", maxofthree(-2, 4, 3));
    printf("%ld\n", maxofthree(2, -6, 5));
    printf("%ld\n", maxofthree(2, 4, 6));
    return 0;
}
```

### makefile

```makefile
# Makefile
CC = gcc
AS = as

objs = main.o add.o

run:
	make clean
	make all
	./run.exe

all: ${objs}
	${CC} -o run.exe ${objs}

.c .o:
	${CC} $<

.s .o:
	${AS} $<

clean:
	-rm *.o
	-rm *.exe
```

### run

```bash
# method 1
linux:~ # as -o maxofthree.o maxofthree.s
linux:~ # gcc -o run.exe maxofthree.o callmaxofthree.c
linux:~ # ./run.exe

# method 2
linux:~ # gcc -o run.exe callmaxofthree.c maxofthree.s
linux:~ # ./run.exe

# method 3
linux:~ # make run
```

---

## arm64

### asm

```s
# add.s
        .text
        .global a_add
a_add:
        add     w0, w0, w1
        ret
```

### c

```c
// main.c
#include<stdio.h>

extern int a_add(int, int);

int main() {
  int a = 5;
  int b = 3;
  printf("%d + %d = %d \n",  a, b, a_add(a, b));
}
```

### makefile

```Makefile
# Makefile
CC = gcc
AS = as

objs = callmaxofthree.o maxofthree.o

run:
	make clean
	make all
	./run.exe

all: ${objs}
	${CC} -o run.exe ${objs}

.c .o:
	${CC} $<

.s .o:
	${AS} $<

clean:
	-rm *.o
	-rm *.exe
```

### run

```bash
# method 1
linux:~ # as -o add.o add.s
linux:~ # gcc -o run.exe add.o main.c
linux:~ # ./run.exe

# method 2
linux:~ # gcc -o run.exe main.c add.s
linux:~ # ./run.exe

# method 3
linux:~ # make run
```

---

## ref

[GNU Assembler Examples](https://cs.lmu.edu/~ray/notes/gasexamples/)
[How to Mix C and Assembly](https://www.devdungeon.com/content/how-mix-c-and-assembly)
