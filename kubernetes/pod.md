# Pod

### Examples

> Create a k3d cluster
1. `k3d cluster create mycluster`
2. `kubectl get nodes`: result ->
```
NAME                     STATUS   ROLES                  AGE     VERSION
k3d-mycluster-server-0   Ready    control-plane,master   2m50s   v1.21.5+k3s1
```


> Create a pod with mongo.

1. `kubectl run db --image mongo`: to create a pod with mongo image
   1. `db`: name of the pod
2. `kubectl get pods`: 
```
NAME   READY   STATUS    RESTARTS  AGE
db     1/1     Running   0         6m
```
3. `docker exec -it k3d-mycluster-server-0 ctr container ls | grep mongo `: login into cluster -> list container -> look for mongo container

```
CONTAINER ID IMAGE
  ...        docker.io/library/mongo:latest
```