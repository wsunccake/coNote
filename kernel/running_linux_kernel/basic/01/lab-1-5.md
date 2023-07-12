# qemu + riscv64 rootfs env

## os

```text
build               --- target
debian 12, x86_64       debian 12 on qemu, riscv64
```

## package

```bash
build:~ # apt install qemu-system
build:~ # apt install mmdebstrap

build:~ # apt install build-essential
build:~ # apt install kmod
build:~ # apt install flex bison
build:~ # apt install libncurses-dev libssl-dev

build:~ # apt install gcc-riscv64-linux-gnu g++-riscv64-linux-gnu
build:~ # apt install debian-ports-archive-keyring
```

---

## prepare

```bash
build:~ # dpkg-architecture -L
build:~ # dpkg --add-architecture riscv64                                     # for riscv64
build:~ # dpkg --remove-architecture riscv64                                  # for riscv64
build:~ # dpkg --print-architecture
build:~ # dpkg --print-foreign-architectures

build:~ # apt update
build:~ # apt install qemu-system-misc qemu-user-static binfmt-support        # for riscv64
```

---

## rootfs

```bash
build:~ # rootfs_dir=/riscv64
build:~ # ARCH=riscv64
build:~ # debian_release=sid

build:~ # mmdebstrap \
  --architectures=$ARCH \
  --include="debian-ports-archive-keyring" \
  $debian_release \
  $rootfs_dir \
  "deb http://deb.debian.org/debian-ports sid main" \
  "deb http://deb.debian.org/debian-ports unreleased main"

# chroot
build:~ # chroot $rootfs_dir /bin/bash
```

```bash
target:~ # apt-get update

# setup password
target:~ # passwd

# setup network
target:~ # cat << EOF >> /etc/network/interfaces
auto lo
iface lo inet loopback

auto eth0
iface eth0 inet dhcp
EOF

# disable the getty on hvc0 as hvc0 and ttyS0 share the same console device in qemu
target:~ # ln -sf /dev/null /etc/systemd/system/serial-getty@hvc0.service

# install kernel and bootloader infrastructure
target:~ # apt-get install linux-image-riscv64 u-boot-menu

# configure syslinux-style boot menu
target:~ # cat << EOF >> /etc/default/u-boot
U_BOOT_PARAMETERS="rw noquiet root=/dev/vda1"
U_BOOT_FDT_DIR="noexist"
EOF

target:~ # u-boot-update

target:~ # exit
```

---

## image / vm disk

```bash
build:~ # rootfs_size=10G
build:~ # rootfs_image=debian_riscv64.ext4
build:~ # rootfs_dir=/riscv64

build:~ # apt-get install libguestfs-tools
build:~ # virt-make-fs --partition=gpt --type=ext4 --size=$rootfs_size $rootfs_dir $rootfs_image
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


build:~/linux-5.15.120 # ls arch/riscv/configs
build:~/linux-5.15.120 # make ARCH=riscv CROSS_COMPILE=riscv64-linux-gnu- nommu_virt_defconfig
build:~/linux-5.15.120 # make ARCH=riscv CROSS_COMPILE=riscv64-linux-gnu- menuconfig
Kernel hacking  --->
  Compile-time checks and compiler options  --->
    [*] Compile the kernel with debug info

build:~/linux-5.15.120 # grep CONFIG_DEBUG_KERNEL=y .config
build:~/linux-5.15.120 # grep CONFIG_DEBUG_INFO=y .config

build:~/linux-5.15.120 # make ARCH=riscv CROSS_COMPILE=riscv64-linux-gnu- -j `nproc`
build:~/linux-5.15.120 # ls arch/riscv/boot/
build:~/linux-5.15.120 # cp arch/riscv/boot/Image $rootfs_dir/boot/.

build:~/linux-5.15.120 # make ARCH=riscv CROSS_COMPILE=riscv64-linux-gnu- -j `nproc` modules
build:~/linux-5.15.120 # make ARCH=riscv CROSS_COMPILE=riscv64-linux-gnu- -j `nproc` modules_install INSTALL_MOD_PATH=../modules/

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
build:~ # QEMU=qemu-system-riscv64
build:~ # SMP=4
build:~ # MEM=2048
build:~ # rootfs_image=debian_arm64.ext4

build:~ # apt install opensbi u-boot-qemu

build:~ # $QEMU \
 -nographic \
 -machine virt \
 -smp $SMP \
 -m $MEM \
 -bios /usr/lib/riscv64-linux-gnu/opensbi/generic/fw_jump.elf \
 -kernel /usr/lib/u-boot/qemu-riscv64_smode/uboot.elf \
 -object rng-random,filename=/dev/urandom,id=rng0 -device virtio-rng-device,rng=rng0 \
 -append "console=ttyS0 rw root=/dev/vda1" \
 -device virtio-blk-device,drive=hd0 -drive file=$rootfs_image,format=raw,id=hd0 \
 -device virtio-net-device,netdev=usernet -netdev user,id=usernet,hostfwd=tcp::22222-:22
```

---

## ref

[RISC-V](https://wiki.debian.org/RISC-V)
