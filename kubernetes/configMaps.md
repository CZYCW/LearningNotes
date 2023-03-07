## Overview
ConfigMaps allow us to keep configurations separate from application images.

Purpose:
- ConfigMap allows us to inject configuration into containers.
- ConfigMap takes a configuration from a source and mounts it into running containers as a volume.


Source:
- files
- directories
- literal values

destination: 
- file or environment variables

## Pratice

### a new start
1. `kubectl config current-context`
2. `kubectl create cm my-config --from-file=prometheus-conf.yml` We created a ConfigMap (cm) called my-config. The data of the map is the content of the prometheus-conf.yml file.
3. `kubectl describe cm my-config`
4. pod with mounted configMaps
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: alpine
spec:
  containers:
  - name: alpine
    image: alpine
    command: ["sleep"]
    args: ["100000"]
    volumeMounts:
    - name: config-vol
      mountPath: /etc/config
  volumes:
  - name: config-vol
    configMap:
      name: my-config
```
5. creating the pod 
```bash
kubectl create -f alpine.yml
kubectl get pods
```
6. `kubectl exec -it alpine -- ls /etc/config`, output `prometheus.yml`

### practice 2: injecting configuraiton files from multiple files
```
kubectl create cm my-config \
    --from-file=cm/prometheus-conf.yml \
    --from-file=cm/prometheus.yml

kubectl create -f cm/alpine.yml

#Run the following command separately
kubectl exec -it alpine -- \
    ls /etc/config
```

### practice 3: injecting configuraiton from a directory
```
kubectl create cm my-config --from-file=cm # cm is a directory 
```
### practice 3 injecting configuration from key/value literals
you can use `--from-literal`, The --from-literal argument is useful when we’re in need to set a relatively small set of configuration entries in different clusters. It makes more sense to specify only the things that change, than all the configuration options.



```
kubectl create cm my-config \
    --from-literal=something=else \
    --from-literal=weather=sunny

kubectl get cm my-config -o yaml
```
### pracice 4: injecting configurations from environment files
we can use `--from-env-file` argument. 

### practice 5: convert configmaps output into environment variables
1. pod definition
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: alpine-env
spec:
  containers:
  - name: alpine
    image: alpine
    command: ["sleep"]
    args: ["100000"]
    env:
    - name: something
      valueFrom:
        configMapKeyRef:
          name: my-config
          key: something
    - name: weather
      valueFrom:
        configMapKeyRef:
          name: my-config
          key: weather
```

2. create the pod
```bash
kubectl create \
    -f alpine-env.yml
#Wait for a few seconds before executing the below command
kubectl exec -it alpine-env -- env # execute env command in the container
```