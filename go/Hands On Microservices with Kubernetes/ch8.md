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

```yaml
# https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/application/mysql/mysql-configmap.yaml

apiVersion: v1
kind: ConfigMap
metadata:
  name: mysql
  labels:
    app: mysql
    app.kubernetes.io/name: mysql
data:
  primary.cnf: |
    # Apply this config only on the primary.
    [mysqld]
    log-bin
  replica.cnf: |
    # Apply this config only on replicas.
    [mysqld]
    super-read-only
```

```yaml
# https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/application/mysql/mysql-services.yaml
# Headless service for stable DNS entries of StatefulSet members.
apiVersion: v1
kind: Service
metadata:
  name: mysql
  labels:
    app: mysql
    app.kubernetes.io/name: mysql
spec:
  ports:
    - name: mysql
      port: 3306
  clusterIP: None
  selector:
    app: mysql
---
# Client service for connecting to any MySQL instance for reads.
# For writes, you must instead connect to the primary: mysql-0.mysql.
apiVersion: v1
kind: Service
metadata:
  name: mysql-read
  labels:
    app: mysql
    app.kubernetes.io/name: mysql
    readonly: "true"
spec:
  ports:
    - name: mysql
      port: 3306
  selector:
    app: mysql
```

```yaml
# https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/application/mysql/mysql-statefulset.yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mysql
spec:
  selector:
    matchLabels:
      app: mysql
      app.kubernetes.io/name: mysql
  serviceName: mysql
  replicas: 3
  template:
    metadata:
      labels:
        app: mysql
        app.kubernetes.io/name: mysql
    spec:
      initContainers:
        - name: init-mysql
          image: mysql:5.7
          command:
            - bash
            - "-c"
            - |
              set -ex
              # Generate mysql server-id from pod ordinal index.
              [[ $HOSTNAME =~ -([0-9]+)$ ]] || exit 1
              ordinal=${BASH_REMATCH[1]}
              echo [mysqld] > /mnt/conf.d/server-id.cnf
              # Add an offset to avoid reserved server-id=0 value.
              echo server-id=$((100 + $ordinal)) >> /mnt/conf.d/server-id.cnf
              # Copy appropriate conf.d files from config-map to emptyDir.
              if [[ $ordinal -eq 0 ]]; then
                cp /mnt/config-map/primary.cnf /mnt/conf.d/
              else
                cp /mnt/config-map/replica.cnf /mnt/conf.d/
              fi
          volumeMounts:
            - name: conf
              mountPath: /mnt/conf.d
            - name: config-map
              mountPath: /mnt/config-map
        - name: clone-mysql
          image: gcr.io/google-samples/xtrabackup:1.0
          command:
            - bash
            - "-c"
            - |
              set -ex
              # Skip the clone if data already exists.
              [[ -d /var/lib/mysql/mysql ]] && exit 0
              # Skip the clone on primary (ordinal index 0).
              [[ `hostname` =~ -([0-9]+)$ ]] || exit 1
              ordinal=${BASH_REMATCH[1]}
              [[ $ordinal -eq 0 ]] && exit 0
              # Clone data from previous peer.
              ncat --recv-only mysql-$(($ordinal-1)).mysql 3307 | xbstream -x -C /var/lib/mysql
              # Prepare the backup.
              xtrabackup --prepare --target-dir=/var/lib/mysql
          volumeMounts:
            - name: data
              mountPath: /var/lib/mysql
              subPath: mysql
            - name: conf
              mountPath: /etc/mysql/conf.d
      containers:
        - name: mysql
          image: mysql:5.7
          env:
            - name: MYSQL_ALLOW_EMPTY_PASSWORD
              value: "1"
          ports:
            - name: mysql
              containerPort: 3306
          volumeMounts:
            - name: data
              mountPath: /var/lib/mysql
              subPath: mysql
            - name: conf
              mountPath: /etc/mysql/conf.d
          resources:
            requests:
              cpu: 500m
              memory: 1Gi
          livenessProbe:
            exec:
              command: ["mysqladmin", "ping"]
            initialDelaySeconds: 30
            periodSeconds: 10
            timeoutSeconds: 5
          readinessProbe:
            exec:
              # Check we can execute queries over TCP (skip-networking is off).
              command: ["mysql", "-h", "127.0.0.1", "-e", "SELECT 1"]
            initialDelaySeconds: 5
            periodSeconds: 2
            timeoutSeconds: 1
        - name: xtrabackup
          image: gcr.io/google-samples/xtrabackup:1.0
          ports:
            - name: xtrabackup
              containerPort: 3307
          command:
            - bash
            - "-c"
            - |
              set -ex
              cd /var/lib/mysql

              # Determine binlog position of cloned data, if any.
              if [[ -f xtrabackup_slave_info && "x$(<xtrabackup_slave_info)" != "x" ]]; then
                # XtraBackup already generated a partial "CHANGE MASTER TO" query
                # because we're cloning from an existing replica. (Need to remove the tailing semicolon!)
                cat xtrabackup_slave_info | sed -E 's/;$//g' > change_master_to.sql.in
                # Ignore xtrabackup_binlog_info in this case (it's useless).
                rm -f xtrabackup_slave_info xtrabackup_binlog_info
              elif [[ -f xtrabackup_binlog_info ]]; then
                # We're cloning directly from primary. Parse binlog position.
                [[ `cat xtrabackup_binlog_info` =~ ^(.*?)[[:space:]]+(.*?)$ ]] || exit 1
                rm -f xtrabackup_binlog_info xtrabackup_slave_info
                echo "CHANGE MASTER TO MASTER_LOG_FILE='${BASH_REMATCH[1]}',\
                      MASTER_LOG_POS=${BASH_REMATCH[2]}" > change_master_to.sql.in
              fi

              # Check if we need to complete a clone by starting replication.
              if [[ -f change_master_to.sql.in ]]; then
                echo "Waiting for mysqld to be ready (accepting connections)"
                until mysql -h 127.0.0.1 -e "SELECT 1"; do sleep 1; done

                echo "Initializing replication from clone position"
                mysql -h 127.0.0.1 \
                      -e "$(<change_master_to.sql.in), \
                              MASTER_HOST='mysql-0.mysql', \
                              MASTER_USER='root', \
                              MASTER_PASSWORD='', \
                              MASTER_CONNECT_RETRY=10; \
                            START SLAVE;" || exit 1
                # In case of container restart, attempt this at-most-once.
                mv change_master_to.sql.in change_master_to.sql.orig
              fi

              # Start a server to send backups when requested by peers.
              exec ncat --listen --keep-open --send-only --max-conns=1 3307 -c \
                "xtrabackup --backup --slave-info --stream=xbstream --host=127.0.0.1 --user=root"
          volumeMounts:
            - name: data
              mountPath: /var/lib/mysql
              subPath: mysql
            - name: conf
              mountPath: /etc/mysql/conf.d
          resources:
            requests:
              cpu: 100m
              memory: 100Mi
      volumes:
        - name: conf
          emptyDir: {}
        - name: config-map
          configMap:
            name: mysql
  volumeClaimTemplates:
    - metadata:
        name: data
      spec:
        accessModes: ["ReadWriteOnce"]
        resources:
          requests:
            storage: 10Gi
```

```bash
# launch service
linux:~ $ kubectl apply -f https://k8s.io/examples/application/mysql/mysql-configmap.yaml
linux:~ $ kubectl apply -f https://k8s.io/examples/application/mysql/mysql-services.yaml
linux:~ $ kubectl apply -f https://k8s.io/examples/application/mysql/mysql-statefulset.yaml
linux:~ $ kubectl get pods -l app=mysql --watch

# test
linux:~ $ kubectl run mysql-client --image=mysql:5.7 -i --rm --restart=Never --\
  mysql -h mysql-0.mysql <<EOF
CREATE DATABASE test;
CREATE TABLE test.messages (message VARCHAR(250));
INSERT INTO test.messages VALUES ('hello');
EOF

linux:~ $ kubectl run mysql-client --image=mysql:5.7 -i -t --rm --restart=Never --\
  mysql -h mysql-read -e "SELECT * FROM test.messages"

linux:~ $ kubectl run mysql-client-loop --image=mysql:5.7 -i -t --rm --restart=Never --\
  bash -ic "while sleep 1; do mysql -h mysql-read -e 'SELECT @@server_id,NOW()'; done"

# simulate pod and node failure
linux:~ $ kubectl exec mysql-2 -c mysql -- mv /usr/bin/mysql /usr/bin/mysql.off
linux:~ $ kubectl get pod mysql-2
linux:~ $ kubectl exec mysql-2 -c mysql -- mv /usr/bin/mysql.off /usr/bin/mysql

# delete pod
linux:~ $ kubectl delete pod mysql-2
linux:~ $ kubectl get pod mysql-2 -o wide

# drain node
linux:~ $ kubectl drain <node-name> --force --delete-emptydir-data --ignore-daemonsets
linux:~ $ kubectl get pod mysql-2 -o wide --watch
linux:~ $ kubectl uncordon <node-name>

# scaling the number of replica
linux:~ $ kubectl scale statefulset mysql  --replicas=5
linux:~ $ kubectl get pods -l app=mysql --watch
linux:~ $ kubectl run mysql-client --image=mysql:5.7 -i -t --rm --restart=Never --\
  mysql -h mysql-3.mysql -e "SELECT * FROM test.messages"
linux:~ $ kubectl scale statefulset mysql --replicas=3
linux:~ $ kubectl get pvc -l app=mysql
linux:~ $ kubectl delete pvc data-mysql-3
linux:~ $ kubectl delete pvc data-mysql-4

# clean up
linux:~ $ kubectl delete pod mysql-client-loop --now
linux:~ $ kubectl delete statefulset mysql
linux:~ $ kubectl get pods -l app=mysql
linux:~ $ kubectl delete configmap,service,pvc -l app=mysql
```

---

## serivce

```bash
linux:~/delinkcious $ git checkout v0.6
linux:~/delinkcious $ cat user_service/k8s/statefulset.yaml
```

---

## ref

[Connecting Applications with Services](https://kubernetes.io/docs/tutorials/services/connect-applications-service/)

[StatefulSet Basics](https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/)

[Run a Replicated Stateful Application](https://kubernetes.io/docs/tasks/run-application/run-replicated-stateful-application/)