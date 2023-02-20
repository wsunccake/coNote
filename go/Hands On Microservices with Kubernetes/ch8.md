# ch8

## service

```yaml
# https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/service/networking/run-my-nginx.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
spec:
  selector:
    matchLabels:
      run: my-nginx
  replicas: 2
  template:
    metadata:
      labels:
        run: my-nginx
    spec:
      containers:
        - name: my-nginx
          image: nginx
          ports:
            - containerPort: 80
```

```yaml
# https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/service/networking/nginx-svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: my-nginx
  labels:
    run: my-nginx
spec:
  ports:
    - port: 80
      protocol: TCP
  selector:
    run: my-nginx
```

```bash
# deployment
linux:~ $ kubectl apply -f https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/service/networking/run-my-nginx.yaml
linux:~ $ kubectl get pods -l run=my-nginx -o wide

# service
linux:~ $ kubectl apply -f https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/service/networking/nginx-svc.yaml
linux:~ $ kubectl get services

# port forward
linux:~ $ kubectl port-forward services/my-nginx 8080:80 &
linux:~ $ curl http://localhost:8080
```

---

## stateful set

```yaml
# https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/application/web/web.yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  ports:
    - port: 80
      name: web
  clusterIP: None # None, cluster 內部存取 service 時，可以直接連到 pod 而不是 service VIP
  selector:
    app: nginx
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web
spec:
  serviceName: "nginx"
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: registry.k8s.io/nginx-slim:0.8
          ports:
            - containerPort: 80
              name: web
          volumeMounts:
            - name: www
              mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
    - metadata:
        name: www
      spec:
        accessModes: ["ReadWriteOnce"]
        resources:
          requests:
            storage: 1Gi
```

```bash
linux:~ $ kubectl apply -f https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/application/web/web.yaml
linux:~ $ kubectl get all

linux:~ $ kubectl exec -it web-0 -- sh -c "hostname > /usr/share/nginx/html/index.html"
linux:~ $ for i in $(seq 0 1); do; kubectl exec -it web-$i -- sh -c "hostname > /usr/share/nginx/html/index.html"; done
```

```yaml
# alpine.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: alpine
spec:
  replicas: 1
  selector:
    matchLabels:
      app: alpine
  template:
    metadata:
      labels:
        app: alpine
    spec:
      containers:
        - name: alpine
          image: alpine
          stdin: true
          tty: true
```

```bash
linux:~ $ kubectl apply -f alpine.html

linux:~ $ nslookup nginx.default.svc.cluster.local
linux:~ $ curl http://nginx.default.svc.cluster.local
```

---

## stateful set - advance

---

## ref

[Connecting Applications with Services](https://kubernetes.io/docs/tutorials/services/connect-applications-service/)

[StatefulSet Basics](https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/)

[Run a Replicated Stateful Application](https://kubernetes.io/docs/tasks/run-application/run-replicated-stateful-application/)
