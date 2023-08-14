# qemu + arm64 rootfs env

---

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
ubuntiu 22, x86_64      ubuntu on qemu, arm64
```

## package

```bash
build:~ # apt install qemu-system
build:~ # apt install debootstrap
build:~ # apt install arch-install-scripts
build:~ # apt install virt-manager virtinst

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
build:~ # rootfs_image=ubuntu_arm64.ext4
build:~ # rootfs_dir=/arm64

build:~ # dd if=/dev/zero of=$rootfs_image bs=1M count=$rootfs_size
build:~ # mkfs.ext4 $rootfs_image
build:~ # mkdir -p $rootfs_dir
build:~ # mount -o loop -t ext4 $rootfs_image $rootfs_dir
build:~ # lsblk

# umount
build:~ # umount $rootfs_dir
build:~ # chmod 777 $rootfs_image
```

---

## rootfs

```bash
build:~ # rootfs_dir=/arm64
build:~ # ARCH=arm64
build:~ # release=jammy                                   # jammy, focal
build:~ # MIRROT=http://ports.ubuntu.com/ubuntu-ports

build:~ # debootstrap \
  --arch $ARCH \
  --foreign \
  --verbose \
  $release \
  $rootfs_dir \
  $MIRROT

build:~ # cp /usr/bin/qemu-aarch64-static $rootfs_dir/usr/bin/.   # for arm64 -> aarch64
build:~ # genfstab -U $rootfs_dir >> $rootfs_dir/etc/fstab

# chroot
build:~ # chroot $rootfs_dir /bin/bash
target:~ # /debootstrap/debootstrap --second-stage
target:~ # passwd

build:~ # chroot $rootfs_dir /bin/bash
```

---

## kernel

```bash
build:~ # curl -OL https://cdn.kernel.org/pub/linux/kernel/v6.x/linux-6.1.43.tar.xz
build:~ # tar Jxf linux-6.1.43.tar.xz
build:~ # ln -s linux-6.1.43/ linux
build:~ # cd linux

build:~/linux # ls arch/arm64/configs
build:~/linux # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- defconfig

# enable debug option by menu
build:~/linux # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- menuconfig
Kernel hacking  --->
  Compile-time checks and compiler options  --->

# enable debug option by script
build:~/linux # ./scripts/config -e CONFIG_DEBUG_KERNEL -e CONFIG_DEBUG_INFO

# check debug option in config
build:~/linux # grep CONFIG_DEBUG_KERNEL=y .config
build:~/linux # grep CONFIG_DEBUG_INFO=y .config

# build kernel
build:~/linux # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- -j `nproc`
build:~/linux # ls arch/arm64/boot/
build:~/linux # cp arch/arm64/boot/Image $rootfs_dir/boot/.

# check enable option after building
build:~/linux # file linux/vmlinux | grep debug

build:~/linux # kernel_version=$(cat include/config/kernel.release)
build:~/linux # mkdir -p $rootfs_dir/modules/$kernel_version
build:~/linux # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- modules
build:~/linux # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- INSTALL_MOD_PATH=$rootfs_dir/modules/$kernel_version modules_install

build:~/linux # kernel_build=$rootfs_dir/usr/src/linux
build:~/linux # mkdir -p $kernel_build
build:~/linux # cp Makefile .config Module.symvers System.map vmlinux $kernel_build
build:~/linux # cp -ar include $kernel_build/usr/src/linux/.

build:~/linux # mkdir -p $kernel_build/arch/arm64/kernel
build:~/linux # cp -a arch/arm64/include $kernel_build/arch/arm64/.
build:~/linux # cp -a arch/arm64/Makefile $kernel_build/arch/arm64/.

build:~/linux # mkdir -p $rootfs_dir/lib/modules/$kernel_version
build:~/linux # ln -s /usr/src/linux $rootfs_dir/lib/modules/$kernel_version/build
```

---

## qemu

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
  -device virtio-blk-device,drive=hd0
"

build:~ # echo $cmd
build:~ # eval $cmd
```
