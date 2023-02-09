# ch5

## package

git, go, docker, minikube / k8s, argocd

---

## config map - literal / key-value

```yaml
# shape-config.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: shape-config
data:
  shape: triangle
```

```bash
# by yaml
linux:~ $ kubectl apply -f shape-config.yaml

# by command
linux:~ $ kubectl create configmap color-config --from-literal color=yellow
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
          env:
            - name: COLOR
              valueFrom:
                configMapKeyRef:
                  name: color-config
                  key: color
            - name: SHAPE
              valueFrom:
                configMapKeyRef:
                  name: shape-config
                  key: shape
```

```bash
linux:~ $ kubectl apply -f alpine.yaml
linux:~ $ kubectl exec -it $(kubectl get pod -l app=alpine -o name) -- env
```

---

## config map - file / directory

```yaml
# game-conf.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: game-config
data:
  game.ini: |
    City = Hamburg
    State = Hamburg
    Country = Germany
```

```bash
# by yaml
linux:~ $ kubectl create configmap steak-config --from-file game-conf.yaml

linux:~ $ cat steak.properties
steak.oz=10
steak.rare=medium-well

# by command
linux:~ $ kubectl create configmap steak-config --from-file steak.properties
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
          volumeMounts:
            - name: steak-volume
              mountPath: /app/steak
            - name: game-volume
              mountPath: /app/game
      volumes:
        - name: steak-volume
          configMap:
            name: steak-config
        - name: game-volume
          configMap:
            name: game-config
```

```bash
linux:~ $ kubectl apply -f alpine.yaml
linux:~ $ kubectl exec -it $(kubectl get pod -l app=alpine -o name) -- ls /app/steak
linux:~ $ kubectl exec -it $(kubectl get pod -l app=alpine -o name) -- ls /app/game
```

---

## environment variable

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
          env:
            - name: DEMO_GREETING
              value: "Hello from the environment"
            - name: DEMO_POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
```

```bash
linux:~ $ kubectl apply -f alpine.yaml
linux:~ $ kubectl exec -it $(kubectl get pod -l app=alpine -o name) -- env
```

---

## example

```bash
# checkout v0.3
linux:~/delinkcious $ git checkout v0.3

# deploy database
linux:~/delinkcious $ kubectl apply -f svc/link_service/k8s/db.yaml

# deploy config map
linux:~/delinkcious $ kubectl apply -f svc/link_service/k8s/configmap.yaml.yaml

# deploy service
linux:~/delinkcious $ kubectl apply -f svc/link_service/k8s/link_manager.yaml
```
