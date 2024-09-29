# ch2 行程管理 (基礎)

## requirement

```bash
# for rhel / fedora
fedora:~ # dnf install procps-ng    # ps
fedora:~ # dnf install binutils     # readelf
fedora:~ # dnf install file         # file

# for debian / ubuntu
debian:~ # apt install

```

## process

```bash
linux:~ # ps aux
linux:~ # ps aux --no-head | wc -l
```

### 練習 fork & execve

```python
# fork-execve.py
import os, sys

pid = os.fork()
# pid > 0: current pid / parent pid
# pid = 0: new pid / child pid
# pid < 0: fail to craete process & OSError

env = os.environ

if pid == 0:
    print(f"child  procrss - child  pid: {os.getpid()}, parent pid: {os.getppid()}")
    os.execve("/bin/echo", ["echo", f"c current pid: {os.getpid()}"], env)
    exit()
else:
    print(f"parent process - parent pid: {os.getpid()}, child  pid: {pid}")
    os.execve("/bin/echo", ["echo", f"p current pid: {os.getpid()}"], env)
    exit()

sys.exit(0)
```

### 練習 posix spawn

```python
# posix_spawn.py
import os, sys

env = os.environ

pid = os.posix_spawn("/bin/echo", ["echo", f"current pid: {os.getpid()}"], env)
# posix_spawn = fork + execve
print(f"pid: {ret}")

sys.exit(0)
```

### 記憶體位置

```bash
linux:~ # gcc -o pause-pie -pie pause.c
linux:~ # gcc -o pause-no-pie -no-pie pause.c
# -pie: produce dynamically linked position independent executable

linux:~ # readelf -h pause-pie
linux:~ # readelf -h pause-no-pie
linux:~ # readelf -S pause-pie
linux:~ # readelf -S pause-no-pie

linux:~ # file pause-pie
linux:~ # file pause-no-pie
# ELF 64-bit LSB pie executable 或 ELF 64-bit LSB executable

linux:~ # pause-no-pie &
linux:~ # cat /proc/<pid>/maps
# 每次執行都使用相同的 memory address
linux:~ # pause-pie &
linux:~ # cat /proc/<pid>/maps
# 每次執行都使用不同的 memory address
```

```bash
linux:~ # pstree -p
linux:~ # ps aux
linux:~ # ps ajx
```

```bash
#!/bin/bash

false &
wait $!
echo "false command to end: $?"
```

---

## signal

```python
# int-ignore.py
import signal

signal.signal(signal.SIGINT, signal.SIG_IGN)

while True:
    pass
```
