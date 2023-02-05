# ch3

## package

git, go, docker

---

## prepare

```bash
linux:~ $ docker pull postgres
linux:~ $ docker run --name postgres -e POSTGRES_PASSWORD=postgres -d postgres
linux:~ $ docker ps
```

## run

```bash
# clone code
linux:~ $ git clone https://github.com/the-gigi/delinkcious.git
linux:~ $ cd delinkcious
linux:~/delinkcious $ git checkout v0.1

# import module
linux:~/delinkcious $ rm go.mod go.sum
linux:~/delinkcious $ go mod init github.com/the-gigi/delinkcious
linux:~/delinkcious $ go mod tiny
linux:~/delinkcious $ run go svc/social_graph_service/main.go
```
