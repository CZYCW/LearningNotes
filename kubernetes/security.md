

### Practice create role binding
1. create role binding
```bash
kubectl create rolebinding jdoe \
    --clusterrole view \
    --user jdoe \
    --namespace default \
    --save-config

kubectl get rolebindings
```

2. describe role binding
```
kubectl describe rolebinding jdoe
```

3. check the scope

```
kubectl --namespace kube-system \
    describe rolebinding jdoe
```

4. check permissions of the role
```
kubectl auth can-i get pods \
    --as jdoe

kubectl auth can-i get pods \
    --as jdoe --all-namespaces
```

5. delete the role binding

```
kubectl delete rolebinding jdoe
```

### Practice create cluster role binding
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: view
subjects:
- kind: User
  name: jdoe
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: view
  apiGroup: rbac.authorization.k8s.io
```

1. create a role binding `kubectl create -f crb-view.yml --record --save-config`

### Pracice combing role binding with namespace
create a new dev namespace for new developer to play. 

