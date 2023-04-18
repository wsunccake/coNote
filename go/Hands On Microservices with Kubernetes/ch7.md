# ch7

## service

```bash
# api_gateway_service
linux:~ $ argocd app create api-gateway-service --repo https://github.com/the-gigi/delinkcious.git --path svc/api_gateway_service/k8s --dest-namespace default --dest-server https://kubernetes.default.svc --revision v0.5
linux:~ $ argocd app sync api-gateway-service

# link_service
linux:~ $ argocd app create link-service --repo https://github.com/the-gigi/delinkcious.git --path svc/link_service/k8s --dest-namespace default --dest-server https://kubernetes.default.svc --revision v0.5
linux:~ $ argocd app sync link-service

# news_service
linux:~ $ argocd app create news-service --repo https://github.com/the-gigi/delinkcious.git --path svc/news_service/k8s --dest-namespace default --dest-server https://kubernetes.default.svc --revision v0.5
linux:~ $ argocd app sync news-service

# social_graph_service
linux:~ $ argocd app create social-graph-service --repo https://github.com/the-gigi/delinkcious.git --path svc/social_graph_service/k8s --dest-namespace default --dest-server https://kubernetes.default.svc --revision v0.5
linux:~ $ argocd app sync social-graph-service

# user_service
linux:~ $ argocd app create user-service --repo https://github.com/the-gigi/delinkcious.git --path svc/user_service/k8s --dest-namespace default --dest-server https://kubernetes.default.svc --revision v0.5
linux:~ $ argocd app sync user-service

# install ingress
linux:~ $ minikube addons enable ingress
linux:~ $ minikube service api-gateway --url

# yaml config
linux:~ $ cat svc/api_gateway_service/k8s/api_gateway.yaml
...
apiVersion: v1
kind: Service
metadata:
  name: api-gateway
spec:
  type: LoadBalancer
  ports:
  - port:  80
    targetPort: 5000
...
# minikube addons enable ingress, support LoadBalancer.
# if k8s, maybe need install other ingress
```

---

## dns

```bash
linux:~ $ kubectl exec -it pods/link-db-<xxxxxxxxxx>-<xxxxx> -- dig link-manager.default.svc.cluster.local
linux:~ $ kubectl exec -it pods/link-db-<xxxxxxxxxx>-<xxxxx> -- nslookup link-manager
linux:~ $ kubectl exec -it pods/link-db-<xxxxxxxxxx>-<xxxxx> -- nslookup link-manager.default.svc.cluster.local
linux:~ $ kubectl exec -it pods/link-db-<xxxxxxxxxx>-<xxxxx> -- cat /etc/resolv.conf

# service name format
# <svc>.<namespace>.svc.cluster.local
# dig <host> [@<server>]
# nslookup [-port=<port>] [-timeout=<number>] <host> [<server>]
```

---

## environment variable

```bash
linux:~ $ kubectl exec -it pods/link-db-<xxxxxxxxxx>-<xxxxx> -- env
linux:~ $ kubectl exec -it pods/social-graph-db-<xxxxxxxxxx>-<xxxxx> -- env

# service environment variable
# <svc>_SERVICE_HOST, <svc>_SERVICE_PORT
# same namespace service export <svc env var> to other service
```

---

## end point

```bash
linux:~ $ kubectl get endpoints
```

---

## ingress

```bash
linux:~ $ minikube addons enable ingress

linux:~ $ kubectl create namespace demo
linux:~ $ kubectl -n demo create deployment hello-minikube --image=k8s.gcr.io/echoserver:1.4
linux:~ $ kubectl -n demo expose deployment hello-minikube --type=NodePort --port=8080

# auto port-forwarding
linux:~ $ minikube -n demo service hello-minikube
linux:~ $ curl $(minikube -n demo service hello-minikube --url)

# manual ingress
linux:~ $ kubectl -n demo apply -f - << EOF
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: example-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /$1
spec:
  rules:
    - host: hello-minikube.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: hello-minikube
                port:
                  number: 8080
EOF
linux:~ $ kubectl -n demo get ingress
linux:~ $ curl -H Host:hello-minikube.example.com  $(minikube -n demo service hello-minikube --url)
```
