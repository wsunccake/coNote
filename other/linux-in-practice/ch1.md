# ch1 Linux 概論

## requirement

```bash
# for rhel / fedora
fedora:~ # dnf install strace           # strace
fedora:~ # dnf util-linux-core          # taskset
fedora:~ # dnf install sysstat          # sar
fedora:~ # dnf install glibc-common     # ldd
fedora:~ # dnf glibc-static

# for debian / ubuntu
debian:~ # apt install strace           # strace
debian:~ # apt install util-linux       # taskset
debian:~ # apt install sysstat          # sar
debian:~ # apt install libc6            # ldd
debian:~ # apt install libc6-dev-x86
```

## system call

### 練習 strace

```c
// hello.c
#include <stdio.h>

int main()
{
    printf("hello\n");
    return 0;
}
```

```go
// hello.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
}
```

```rust
// hello.rs
fn main() {
    println!("hello");
}
```

```python
# hello.py
print("hello")
```

```bash
linux:~ # gcc -o hello hello.c
linux:~ # starce ./hello |& grep write

linux:~ # go build -o hello hello.go
linux:~ # starce ./hello |& grep write

linux:~ # rustc -o hello hello.rs
linux:~ # starce ./hello |& grep write

linux:~ # strace python3 hello.py |& grep write

linux:~ # strace -T -o hello.log python3 hello.py
# -T: show system call time
# -o: write trace output

linux:~ # grep write hello.log
# write is system call api
```

### 練習 sar

```python
# inf-loop.py
import os

while True:
    os.getppid()
```

```bash
linux:~ # taskset -c 0 python3 inf-loop.py &
# -c: cpu list

linux:~ # sar -p 0 1
# -p: cpu list
# 注意 %user 跟 %idle 變化
# CPU     %user     %nice   %system   %iowait    %steal     %idle
```

---

## library

```c
// pause.c
#include <unistd.h>

int main()
{
    pause();
    return 0;
}
```

```bash
# dynamic link
linux:~ # gcc -o pause pause.c
linux:~ # ldd pause

# static link
linux:~ # gcc -static -o pause pause.c
linux:~ # ldd pause
```
