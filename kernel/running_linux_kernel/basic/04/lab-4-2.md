# debug linux kernel module

---

## content

- [environment](#environment)
- [package](#package)
- [debug](#debug)

---

## environment

[qemu + arm64 on ubuntu](../01/lab-1-4-ubuntu.md)

---

## package

```bash
build:~ # apt install gdb-multiarch
```

---

## debug

```bash
build:~ # QEMU=qemu-system-aarch64
build:~ # SMP=4
build:~ # MEM=1024
build:~ # kernel_image=linux/arch/arm64/boot/Image

build:~ # cmd="$QEMU \
  -smp $SMP \
  -m $MEM \
  -cpu cortex-a57 \
  -M virt \
  -nographic \
  -kernel $kernel_image \
  -append \"noinintrd sched_debug root=/dev/vda rootfstype=ext4 rw crashkernel=256M loglevel=8\" \
  -drive if=none,file=$rootfs_image,id=hd0 \
  -device virtio-blk-device,drive=hd0 \
  -s -S
"
# -S: do not start CPU at startup, waiting debug
# -s: -gdb tcp::1234

build:~ # echo $cmd
build:~ # eval $cmd
```

```bash
build:~/linux # gdb-multiarch
(gdb) file vmlinux

(gdb) layout split

(gdb) show architecture
(gdb) show sysroot
(gdb) show solib-search-path

(gdb) set architecture aarch64
(gdb) set sysroot ~/test
(gdb) set solib-search-path /usr/lib/aarch64-linux-gnu/

(gdb) target remote :1234

(gdb) hbreak start_kernel
(gdb) contiune
(gdb) lx-dmesg

(gdb) quit
```
