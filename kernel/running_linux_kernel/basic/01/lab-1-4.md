# qemu + arm64 rootfs env

## content

- [os](#os)
- [package](#package)
- [prepare](#prepare)
- [image / vm disk](#image--vm-disk)
- [rootfs](#rootfs)
- [kernel](#kernel)
- [qemu](#qemu)

---

## os

```text
build               --- target
debian 12, x86_64       debian 12 on qemu, arm64
```

## package

```bash
build:~ # apt install qemu-system
build:~ # apt install debootstrap

build:~ # apt install build-essential
build:~ # apt install kmod
build:~ # apt install flex bison
build:~ # apt install libncurses-dev libssl-dev

build:~ # apt install crossbuild-essential-arm64
```

---

## prepare

```bash
build:~ # dpkg-architecture -L
build:~ # dpkg --add-architecture arm64                                    # for arm64
build:~ # dpkg --remove-architecture arm64                                 # for arm64
build:~ # dpkg --print-architecture
build:~ # dpkg --print-foreign-architectures

build:~ # apt update
build:~ # apt install qemu-system:arm64 qemu-user:arm64 qemu-user-static   # for arm64
```

---

## image / vm disk

```bash
build:~ # rootfs_size=2048
build:~ # rootfs_image=debian_arm64.ext4
build:~ # rootfs_dir=/arm64

build:~ # dd if=/dev/zero of=$rootfs_image bs=1M count=$rootfs_size
build:~ # mkfs.ext4 $rootfs_image
build:~ # mkdir -p $rootfs_dir
build:~ # mount -t ext4 $rootfs_image $rootfs_dir -o loop
build:~ # lsblk

build:~ # umount $rootfs_dir
build:~ # chmod 777 $rootfs_image
```

---

## rootfs

```bash
build:~ # rootfs_dir=/arm64
build:~ # ARCH=arm64
build:~ # debian_release=bookworm                                 # bookworm, bulleye, buster
build:~ # MIRROT=http://ftp.tw.debian.org/debian

build:~ # debootstrap \
  --arch $ARCH \
  --foreign \
  --keyring /usr/share/keyrings/debian-archive-keyring.gpg \
  --verbose \
  $debian_release \
  $rootfs_dir \
  $MIRROT

build:~ # cp /usr/bin/qemu-aarch64-static $rootfs_dir/usr/bin/.   # for arm64 -> aarch64

# chroot
build:~ # chroot $rootfs_dir /bin/bash
target:~ # /debootstrap/debootstrap --second-stage
target:~ # passwd

# test
build:~ # chroot $rootfs_dir /bin/bash
```

---

## kernel

```bash
build:~ # curl -OL https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.15.120.tar.xz
build:~ # tar Jxf linux-5.15.120.tar.xz
build:~ # cd linux-5.15.120/

build:~ # curl -OL https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.10.186.tar.xz
build:~ # tar Jxf linux-5.10.186.tar.xz
build:~ # cd linux-5.10.186/


build:~/linux-5.15.120 # ls arch/arm64/configs
build:~/linux-5.15.120 # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- defconfig
build:~/linux-5.15.120 # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- menuconfig
Kernel hacking  --->
  Compile-time checks and compiler options  --->
    [*] Compile the kernel with debug info

build:~/linux-5.15.120 # grep CONFIG_DEBUG_KERNEL=y .config
build:~/linux-5.15.120 # grep CONFIG_DEBUG_INFO=y .config

build:~/linux-5.15.120 # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- -j `nproc`
build:~/linux-5.15.120 # ls arch/arm64/boot/
build:~/linux-5.15.120 # cp arch/arm64/boot/Image $rootfs_dir/boot/.

build:~/linux-5.15.120 # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- -j `nproc` modules
build:~/linux-5.15.120 # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- -j `nproc` modules_install INSTALL_MOD_PATH=../modules/

build:~/linux-5.15.120 # kernel_version=$(cat include/config/kernel.release)
build:~/linux-5.15.120 # kernel_build=$rootfs_dir/usr/src/linux

build:~/linux-5.15.120 # mkdir -p $kernel_build
build:~/linux-5.15.120 # cp Makefile .config Module.symvers System.map vmlinux $kernel_build
build:~/linux-5.15.120 # cp -a include $kernel_build/usr/src/linux/.

build:~/linux-5.15.120 # mkdir -p $kernel_build/arch/arm64/kernel
build:~/linux-5.15.120 # cp -a arch/arm64/include $kernel_build/arch/arm64/.
build:~/linux-5.15.120 # cp -a arch/arm64/Makefile $kernel_build/arch/arm64/.
build:~/linux-5.15.120 # cp arch/arm64/kernel/module.lds $kernel_build/arch/arm64/kernel/.

build:~/linux-5.15.120 # mkdir -p $rootfs_dir/lib/modules/$kernel_version
build:~/linux-5.15.120 # ln -s /usr/src/linux $rootfs_dir/lib/modules/$kernel_version/build
```

---

## qemu

```bash
build:~ # QEMU=qemu-system-aarch64
build:~ # SMP=4
build:~ # MEM=1024
build:~ # rootfs_image=debian_arm64.ext4
build:~ # kernel_image=linux-5.15.120/arch/arm64/boot/Imag

build:~ # cmd="$QEMU \
  -smp $SMP \
  -m $MEM \
  -cpu cortex-a57 \
  -M virt \
  -nographic \
  -kernel $kernel_image \
  -append "noinintrd sched_debug root=/dev/vda rootfstype=ext4 rw crashkernel=256M loglevel=8" \
  -drive if=none,file=$rootfs_image,id=hd0 \
  -device virtio-blk-device,drive=hd0
"

build:~ # echo $cmd
build:~ # eval $cmd
```
