# ch6

## user account

```bash
# list user account
linux:~ $ kubectl config get-users
NAME
minikube

# config info
linux:~ $ kubectl config view
apiVersion: v1
clusters:
- cluster:
    certificate-authority: $HOME/.minikube/ca.crt
    extensions:
    - extension:
        last-update: Fri, 10 Feb 2023 00:22:27 CST
        provider: minikube.sigs.k8s.io
        version: v1.29.0
      name: cluster_info
    server: https://192.168.49.2:8443
  name: minikube
contexts:
- context:
    cluster: minikube
    extensions:
    - extension:
        last-update: Fri, 10 Feb 2023 00:22:27 CST
        provider: minikube.sigs.k8s.io
        version: v1.29.0
      name: context_info
    namespace: default
    user: minikube
  name: minikube
current-context: minikube
kind: Config
preferences: {}
users:
- name: minikube
  user:
    client-certificate: $HOME/.minikube/profiles/minikube/client.crt
    client-key: $HOME/.minikube/profiles/minikube/client.key
```

```bash
# convert certificate to file
linux:~ $ kubectl config view | grep client-certificate
linux:~ $ echo <client-certificate code> | base64 -d > client.crt

# convert key to file
linux:~ $ kubectl config view | grep client-key
linux:~ $ echo <client-key code> | base64 -d > client.key
```

---

## service account

```bash
# list service account
linux:~ $ kubectl get serviceaccounts
NAME                  SECRETS   AGE
default               0         23h
```

### create service account to bind role

```bash
# create service account by command
linux:~ $ kubectl create serviceaccount api-service-account

# create service account by yaml
linux:~ $ kubectl apply -f - << EOF
apiVersion: v1
kind: ServiceAccount
metadata:
  name: api-service-account
EOF

# create token by command - for short term
linux:~ $ kubectl create token api-service-account [-duration 10m]

# create token by yaml - for long term
linux:~ $ kubectl apply -f - << EOF
apiVersion: v1
kind: Secret
metadata:
  name: api-service-account-token
  annotations:
    kubernetes.io/service-account.name: api-service-account
type: kubernetes.io/service-account-token
EOF

# create role
linux:~ $ cat << EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: api-cluster-role
rules:
  - apiGroups:
      - ""
      - apps
      - autoscaling
      - batch
      - extensions
      - policy
      - rbac.authorization.k8s.io
    resources:
      - pods
      - componentstatuses
      - configmaps
      - daemonsets
      - deployments
      - events
      - endpoints
      - horizontalpodautoscalers
      - ingress
      - jobs
      - limitranges
      - namespaces
      - nodes
      - pods
      - persistentvolumes
      - persistentvolumeclaims
      - resourcequotas
      - replicasets
      - replicationcontrollers
      - serviceaccounts
      - services
    verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]
EOF

# service account bind role by command
linux:~ $ kubectl create clusterrolebinding api-service-binding --clusterrole=cluster-admin --serviceaccount=default:api-service-account

# service account bind role by yaml
linux:~ $ cat << EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: api-cluster-role-binding
subjects:
- kind: ServiceAccount
  name: api-service-account
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: api-cluster-role
EOF

# validate
linux:~ $ kubectl auth can-i get pods --as=system:serviceaccount:default:api-service-account
```

### token

```bash
# token
linux:~ $ kubectl get  secret api-service-account-token | grep token
linux:~ $ echo <token code> | base64 -d

# ca.crt
linux:~ $ kubectl get  secret api-service-account-token | grep ca.crt
linux:~ $ echo <ca crt code> | base64 -d > <ca crt file>
```

### kubectl

```bash
kubectl config set-cluster <new context> --server=https://192.168.49.2:8443 --certificate-authority <ca crt file>
kubectl config set-context <new context> --cluster=<new context>
kubectl config set-credentials user --token <token>
kubectl config set-context <new context> --user=<user>
kubectl config use-context <new context>
```

---

## api

```bash
linux:~ $ export CLUSTER_NAME=minikube
linux:~ $ export API_SERVER=https://192.168.49.2:8443

# by certificate and key
linux:~ $ curl --insecure --cert $CERT --key $KEY $API_SERVER/api
# CERT: certificate, KEY: private key

# by token
linux:~ $ curl --insecure  --header "Authorization: Bearer $TOKEN" $API_SERVER/api
# TOKEN: service account totken
```

---

## example

```bash
# argocd
linux:~ $ argocd app create link --repo https://github.com/the-gigi/delinkcious.git --path svc/link_service/k8s --dest-namespace default --dest-server https://kubernetes.default.svc --revision v0.4
linux:~ $ argocd app list
linux:~ $ argocd app sync link
```