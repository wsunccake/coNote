# prepare

---

## content

- [intstall](#install)
  - [linux](#linux)
  - [macosx](#macosx)
  - [windows](#windows)
- [editor / ide](#editor--ide)
  - [vi](#vi)
  - [vscode](#vscode)
  - [goland / intellij](#goland--intellij)
- [get started](#get-started)
  - [hello](#hello)
  - [module](#module)

---

## install

### download

[All releases](https://go.dev/dl/)

### linux

```bash
# for rhel / centos / fedora
linux:~ # yum install golang
linux:~ # dnf install golang

# for debian / ubuntu
linux:~ # apt install golang

# for binary
linux:~ # curl -OL https://go.dev/dl/go1.21.4.linux-amd64.tar.gz
linux:~ # tar -xzf go1.21.4.linux-amd64.tar.gz -C /usr/local
linux:~ # ln -s /usr/local/go/bin/go /usr/local/bin/.
```

### macosx

### windows

### gvm

```bash
# for rhel / centos / fedora
linux:~ # yum install bison
linux:~ # dnf install bison

# for debian / ubuntu
linux:~ # apt install bison

linux:~ $ bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
linux:~ $ source /home/$(whoami)/.gvm/scripts/gvm
```

```bash
linux:~ $ gvm help
linux:~ $ gvm install go1.21.1 [--binary]
linux:~ $ gvm uninstall go1.21.1

linux:~ $ gvm list
linux:~ $ gvm use go1.21.1 [--default]
```

[gvm](https://github.com/moovweb/gvm)

---

## editor / ide

### vi

- [vim](https://www.vim.org/) with plugin [vim-go](https://github.com/fatih/vim-go)

### vscode

- [vscode](https://code.visualstudio.com/) with plugin [Go](https://marketplace.visualstudio.com/items?itemName=golang.go)

### goland / intellij

- [goland](https://www.jetbrains.com/go/) or [intellij](https://www.jetbrains.com/idea/) with plugin [go](https://plugins.jetbrains.com/plugin/9568-go)

---

## get started

## hello

- [Get started with Go](https://go.dev/doc/tutorial/getting-started)

## module

- [Create a Go module](https://go.dev/doc/tutorial/create-module)
