## Production Ready Kubernetes

### Constituents
- docker: to deploy the image to container
- kubelet: it is the agent on the node
  - ensure the containers in pod are running health. 
  - get pod information from API server
- Protokube
  - only for kops
  
### System level component
- to get `kubectl --namespace kube-system get pods`

Master Components: run only on master node
- Kubernetes API Server: accept requests to create, update or remove kubernetes resources.
  - listen on port: 8080 
  - listen on port: 443 -> https for external traffic
- etcd
  - etcd-server: a key-value store which hold the state of the cluster
  - etcd-server-events: stores events
  - kops create ebs volumne for each etcd instance
- Kubernetes controller manager
  - in charge of running controller
  - monitoring services
- Kubernetes Scheduler: watches API server for new pods and assign them to a node
- DNS Controller: allow nodes and users to discover the API server. 

### Node Components
- kube-proxy: reflects services defines through API server. 
  - in charge of TCP and UDP forwarding