# initialize project

```bash
linux:~ $ mkdir monkey
linux:~ $ cd monkey
linux:~/monkey $ go mod init monkey
linux:~/monkey $ ls
linux:~/monkey $ cat go.mod
```


---

# source code

## repl / Read-Eval-Print Loop

```bash
linux:~/monkey $ cat repl/repl.go
```

## token

```bash
linux:~/monkey $ cat token/token.go
```

## lexer

```bash
linux:~/monkey $ cat lexer/lexer.go
```


---

# build tool

## make

```bash
linux:~/monkey $ cat makefile

linux:~/monkey $ make init      # create go module

linux:~/monkey $ make           # build binary
linux:~/monkey $ make main      # build binary

linux:~/monkey $ make clean     # remove binary
```
