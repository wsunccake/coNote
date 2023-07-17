# install ubunut arm64 by virt

## content

- [package](#package)
- [install](#install)

---

## package

```bash
debian:~ # apt intall libvirt-clients
debian:~ # apt intall virt-manager virtinst
debian:~ # apt intall libguestfs-tools
```

---

## install

```bash
# cli
debian:~ # VM_NAME=ubuntu-22.04-aarch64
debian:~ # VM_DISK=$VM_DISK-vm.qcow2
debian:~ # ISO_FILE=ubuntu-22.04.2-live-server-arm64.iso

debian:~ # wget https://cdimage.ubuntu.com/releases/22.04/release/$ISO_FILE
debian:~ # qemu-img create -f qcow2 $VM_DISK 20G

debian:~ # virt-install -v --name $VM_NAME \
  --arch aarch64 \
  --vcpus 4 \
  --ram 4096 \
  --disk path=$VM_DISK,bus=virtio \
  --boot uefi \
  --cdrom $ISO_FILE \
  --import \
  --nographics

# gui
debian:~ # virt-manager
```
