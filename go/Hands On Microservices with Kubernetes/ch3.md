# ch3

## package

git, go, docker

vscode / goland

---

## prepare

```bash
# install postgres
linux:~ $ docker pull postgres
linux:~ $ docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432-d postgres
linux:~ $ docker ps

# craete database
linux:~ $ docker exec -it postgres psql -h localhost -U postgres
postgres=# CREATE DATABASE social_graph_manager;
postgres=# \l
postgres=# quit
```

## service

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

# test
linux:~ $ curl -X POST -d '{"followed": "a", "follower": "b"}' http://localhost:9090/follow

linux:~ $ docker exec -it postgres psql -h localhost -U postgres
postgres=# \c social_graph_manager
social_graph_manager=# select * from social_graph;
```
