# call c from assembly

## content

- [arm64](#arm64)
  - [asm](#asm)
  - [run](#run)

---

## arm64

### asm

```s
// hello.s
.data

.balign 8
/* This is the greeting message */
say_hello: .asciz "Hello world!"

.balign 8
/* We need to keep x30 otherwise we will not be able to return from main! */
keep_x30: .dword 0

.text

/* We are going to call a C-library puts function */
.globl puts

.globl main
main:
    ldr x0, addr_keep_x30     // w0 ← &keep_30   [64]
    str x30, [x0]             // *keep_30 ← x30  [64]

    ldr x0, addr_say_hello    // w0 ← &say_hello [64]
    bl puts                   // call puts

    ldr x0, addr_keep_x30     // w0 ← &keep_30   [64]
    ldr x30, [x0]             // x30 ← *keep_30  [64]

    mov w0, #0                // w0 ← 0
    ret                       // return

addr_keep_x30 : .dword keep_x30
addr_say_hello: .dword say_hello
```

### run

```bash
linux:~ # gcc -o hello hello.s
linux:~ # ./hello
```

---

## ref

[Exploring AArch64 assembler – Chapter 7](https://thinkingeek.com/2017/03/19/exploring-aarch64-assembler-chapter-7/)
