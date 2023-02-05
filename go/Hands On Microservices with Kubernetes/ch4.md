# ch4

## package

git, go, docker, minikube / k8s, argocd

vscode / goland

---

## build image

```bash
# checkout v0.2
linux:~/delinkcious $ git checkout v0.2

# update dockerfile
linux:~/delinkcious $ vi svc/social_graph_service/Dockerfile
FROM golang:1.11 AS builder
ADD ./main.go main.go
ADD ./service service

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /social_graph_service -a -tags netgo -ldflags '-s -w' .
->
FROM golang:1.19 AS builder
ADD ./ /root/data

RUN cd /root/data && go get -d -v ./...
RUN cd /root/data/svc/social_graph_service && CGO_ENABLED=0 GOOS=linux go build -o /social_graph_service -a -tags netgo -ldflags '-s -w' .

# build image
linux:~/delinkcious $ docker build . -t g1g1/delinkcious-social-graph:0.2

# push image
linux:~ $ dokcer images
linux:~ $ docker login -u $DOCKERHUB_USERNAME -p $DOCKERHUB_PASSWORD
linux:~ $ docker push g1g1/delinkcious-social-graph:0.2
```

-> ci

---

## deploy image

```bash
linux:~/delinkcious $ kubectl apply -f svc/social_graph_service/k8s
```

-> cd

```bash
# port forward
linux:~ $ kubectl -- port-forward svc/argocd-server -n argocd 8080:443

# argo password
linux:~ $ kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

# argo cd
linux:~ $ argocd login :8080 --username admin --password <password> --insecure
argocd app create social --repo https://github.com/the-gigi/delinkcious.git --path svc/social_graph_service/k8s --dest-namespace default --dest-server https://kubernetes.default.svc --revision v0.2
linux:~ $ argocd app list
linux:~ $ argocd app sync social
```
