# dev env

## VM setup

```
OS: Ubuntu 26 - Desktop
CPU: 4
RAM: 8 Gb
HD:  50 GB
```

## package

```
hypervisor: qemu-kvm
editor:     vscode
extension:  rust-analyzer
lang:       nasm, rust
```

## install

```bash
# 1. 更新套件庫並安裝開發工具、QEMU 與 NASM
ubuntu:~ # apt update
ubuntu:~ # apt install -y build-essential nasm qemu-system-x86 gdb
ubuntu:~ # apt install -y curl

# 2. 安裝 Rust (若尚未安裝)
ubuntu:~ $ curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# 3. 新增 16-bit / 32-bit Bare-metal Target
ubuntu:~ $ rustup target add i686-unknown-linux-gnu
```
