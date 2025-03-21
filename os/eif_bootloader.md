# bootloader

## content

- [UEFI](#uefi)

---

## UEFI

GNU-EFI

```bash
# for debian / ubuntu
linux:~ # apt update
linux:~ # apt install gnu-efi

# for rhel / fedora
linux:~ # dnf makecache
linux:~ # dnf gnu-efi gnu-efi-devel
```

---

## 64-bit uefi bootloader

```c
// boot.c
#include <efi.h>
#include <efilib.h>

EFI_STATUS
EFIAPI
efi_main (EFI_HANDLE ImageHandle, EFI_SYSTEM_TABLE *SystemTable)
{
  InitializeLib(ImageHandle, SystemTable);
  Print(L"Hello, world!\n");
  return EFI_SUCCESS;
}
```

```bash
linux:~ # gcc -c boot.c \
  -fno-stack-protector \
  -fpic \
  -fshort-wchar \
  -mno-red-zone \
  -I /usr/include/efi \
  -DEFI_FUNCTION_WRAPPER \
  -o boot.o

linux:~ # ld boot.o \
     /usr/lib/crt0-efi-x86_64.o \
     -nostdlib \
     -znocombreloc \
     -T /usr/lib/elf_x86_64_efi.lds \
     -shared \
     -Bsymbolic \
     -L /usr/lib \
     -l:libgnuefi.a \
     -l:libefi.a \
     -o boot.so

linux:~ # objcopy -j .text \
          -j .sdata \
          -j .data \
          -j .dynamic \
          -j .dynsym \
          -j .rel \
          -j .rela \
          -j .reloc \
          --target=efi-app-x86_64 \
          boot.so \
          boot.efi

linux:~ # file boot.efi
boot.efi: PE32+ executable (EFI application) x86-64 (stripped to external PDB), for MS Windows, 7 sections
```

---

### vdisk

```bash
# create raw img disk
linux:~ # dd if=/dev/zero of=uefi.img bs=1M count=1024

# partition
linux:~ # parted -s uefi.img mklabel msdos
linux:~ # parted -s uefi.img mkpart primary fat32 1MiB 81MiB
linux:~ # parted -s uefi.img set 1 lba off
linux:~ # parted -s uefi.img mkpart primary ext4 81MiB 100%
linux:~ # parted -s uefi.img print
linux:~ # fdisk -l uefi.img

# dev map
linux:~ # losetup /dev/loop1 uefi.img
linux:~ # losetup
linux:~ # kpartx -av /dev/loop1
linux:~ # kpartx -lv /dev/loop1

# format
linux:~ # mkfs -t msdos /dev/mapper/loop1p1
linux:~ # mkfs -t ext4 -F /dev/mapper/loop1p2

# copy efi
linux:~ # mount /dev/mapper/loop1p1 /mnt
linux:~ # mkdir -p /mnt/efi/EFI/BOOT
linux:~ # cp boot.efi /mnt/efi/EFI/BOOT/BOOTX64.EFI
linux:~ # umount /mnt

# umount
linux:~ # kpartx -dv /dev/loop1
linux:~ # losetup -d /dev/loop1
```

### vm

```bash
linux:~ # qemu-system-x86_64 \
  -cpu qemu64 \
  -bios /usr/share/edk2/ovmf/OVMF_CODE.fd \
  -drive file=uefi.img,format=raw,if=ide \
  --nographic
```

```bash
Shell> FS0:
Shell> efi/EFI/BOOT/BOOTX64.EFI
```
