# build arm64 kernel

---

## content

- [install toolchain](#install-toolchain)
- [build kernel](#build-kernel)

---

## install toolchain

```bash
ubuntu:~ # apt install crossbuild-essential-arm64   # for arm64
```

---

## build kernel

```bash
ubuntu:~ # curl -OL https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.15.120.tar.xz
ubuntu:~ # tar Jxf linux-5.15.120.tar.xz
ubuntu:~ # cd linux-5.15.120/

ubuntu:~/linux-5.15.120 # export rootfs_dir=/arm64

ubuntu:~/linux-5.15.120 # ls arch/arm64/configs
ubuntu:~/linux-5.15.120 # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- defconfig
ubuntu:~/linux-5.15.120 # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- menuconfig
Kernel hacking --->
Compile-time checks and compiler options --->
[*] Compile the kernel with debug info

ubuntu:~/linux-5.15.120 # grep CONFIG_DEBUG_KERNEL=y .config
ubuntu:~/linux-5.15.120 # grep CONFIG_DEBUG_INFO=y .config

ubuntu:~/linux-5.15.120 # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- -j `nproc`
ubuntu:~/linux-5.15.120 # ls arch/arm64/boot/
ubuntu:~/linux-5.15.120 # cp arch/arm64/boot/Image $rootfs_dir/boot/.

ubuntu:~/linux-5.15.120 # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- -j `nproc` modules
ubuntu:~/linux-5.15.120 # make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- -j `nproc` modules_install INSTALL_MOD_PATH=../modules/

ubuntu:~/linux-5.15.120 # kernel_version=$(cat include/config/kernel.release)
ubuntu:~/linux-5.15.120 # kernel_build=$rootfs_dir/usr/src/linux

ubuntu:~/linux-5.15.120 # mkdir -p $kernel_build
ubuntu:~/linux-5.15.120 # cp Makefile .config Module.symvers System.map vmlinux $kernel_build
ubuntu:~/linux-5.15.120 # cp -a include $kernel_build/usr/src/linux/.

ubuntu:~/linux-5.15.120 # mkdir -p $kernel_build/arch/arm64/kernel
ubuntu:~/linux-5.15.120 # cp -a arch/arm64/include $kernel_build/arch/arm64/.
ubuntu:~/linux-5.15.120 # cp -a arch/arm64/Makefile $kernel_build/arch/arm64/.
ubuntu:~/linux-5.15.120 # cp arch/arm64/kernel/module.lds $kernel_build/arch/arm64/kernel/.

ubuntu:~/linux-5.15.120 # mkdir -p $rootfs_dir/lib/modules/$kernel_version
ubuntu:~/linux-5.15.120 # ln -s /usr/src/linux $rootfs_dir/lib/modules/$kernel_version/build
```
