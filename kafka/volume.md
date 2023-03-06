## Overview

Kubernetes Volumes solve the need to . In essence, Volumes are references to files and directories made accessible to containers that form a Pod. The significant difference between different types of Kubernetes Volumes is in the way these files and directories are created.

### First Practice
- create a cluster by copying `usercode/volume` from user directory to `/file` in the cluster: `k3d cluster create mycluster --volume "/usercode/volume/prometheus-conf.yml:/files/prometheus-conf.yml"`

### Practice: run docker in kubernetes
1.
```bash
kubectl run docker \
    --image=docker:17.11  --restart=Never \
    docker image ls

kubectl get pods
```
2. 
error occurs
3. `kubectl logs -f docker`
4. output: `Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?`
5. Why: Docker consists of two main pieces. There is a client, and there is a server. When we executed docker image ls, we invoked the client which tried to communicate with the server through its API. The problem is that Docker server is not running in that container. What we should do is tell the client (inside a container) to use Docker server that is already running on the host. Docker consists of two main pieces. There is a client, and there is a server. When we executed docker image ls, we invoked the client which tried to communicate with the server through its API. The problem is that Docker server is not running in that container. What we should do is tell the client (inside a container) to use Docker server that is already running on the host.
6. `kubectl delete pod docker` delete this resource

**The correct way to do that**
To mount the file `/var/run/docker.sock` `
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: docker
spec:
  containers:
  - name: docker
    image: docker:17.11
    command: ["sleep"]
    args: ["100000"]
    volumeMounts:
    - mountPath: /var/run/docker.sock
      name: docker-socket
  volumes:
  - name: docker-socket
    hostPath:
      path: /var/run/docker.sock
      type: Socket
```

### Volumne Types
- Host Path: A hostPath Volume maps a directory from a host to where the Pod is running. Using it to “inject” configuration files into containers would mean that we’d have to make sure that the file is present on every node of the cluster.
  - Do use hostPath to mount host resources like /var/run/docker.sock and /dev/cgroups. Do not use it to inject configuration files or store the state of an application.


