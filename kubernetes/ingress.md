# Ingress

## Overview
Kubernetes's ingress controler is not provided by kubernetes. Kubernetes only provides its api, you need to install another controller for that.

- We can define multiple Ingress resources that will configure a single Ingress Controller.
- Ingress is a (kind of) Service that runs on all nodes of a cluster. A user can send requests to any and, as long as they match one of the rules, they will be forwarded to the appropriate Service.


## First Practice
- enable ingress on k3d: `kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.3.0/deploy/static/provider/cloud/deploy.yaml`
    - the ingress image is based on nginx ingress controller
- get status for ingress controller: `kubectl get pods --namespace=ingress-nginx | grep ingress`
  - `-n` stands for namespace
- To check controller's health
```bash
nohup kubectl port-forward -n ingress-nginx service/ingress-nginx-controller 3000:80 --address 0.0.0.0 > /dev/null 2>&1 &
curl -i "0.0.0.0:3000/healthz"
```
  
### Yaml definition
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: go-demo-2
  annotations: # allow us to provide additional information to the ingress controller
    kubernetes.io/ingress.class: "nginx"
    ingress.kubernetes.io/ssl-redirect: "false"
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - http:
      paths:
      - path: /demo # all request with path starting with /demo will be forwarded to the service go-demo-2-api on port 8080
        pathType: ImplementationSpecific
        backend:
          service:
            name: go-demo-2-api
            port:
              number: 8080
```

- create ingress resource: 

```bash
kubectl create -f go-demo-2-ingress.yml
kubectl get -f go-demo-2-ingress.yml`
```

- send request to `/demo`
  
```bash
nohup kubectl port-forward -n ingress-nginx service/ingress-nginx-controller 3000:80 --address 0.0.0.0 > /dev/null 2>&1 &
curl -i "http://0.0.0.0:3000/demo/hello" 
# -i means include header in the http response
```

- delete ingress resource

```bash
kubectl delete -f go-demo-2-ingress.yml

kubectl delete -f go-demo-2-deploy.yml
```

- Create ingress resource using unified YAML
  
```bash
kubectl create -f go-demo-2.yml \
    --record --save-config

nohup kubectl port-forward -n ingress-nginx service/ingress-nginx-controller 3000:80 --address 0.0.0.0 > /dev/null 2>&1 &
curl -i "http://0.0.0.0:3000/demo/hello"
```

## Second Practice

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: devops-toolkit
  annotations:
    kubernetes.io/ingress.class: "nginx"
    ingress.kubernetes.io/ssl-redirect: "false"
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - http:
      paths:
      - path: /
        pathType: ImplementationSpecific
        backend:
          service:
            name: devops-toolkit
            port:
              number: 80

---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: devops-toolkit
spec:
  replicas: 3
  selector:
    matchLabels:
      type: frontend
      service: devops-toolkit
  template:
    metadata:
      labels:
        type: frontend
        service: devops-toolkit
    spec:
      containers:
      - name: frontend
        image: vfarcic/devops-toolkit-series

---

apiVersion: v1
kind: Service
metadata:
  name: devops-toolkit
spec:
  ports:
  - port: 80
  selector:
    type: frontend
    service: devops-toolkit
```
    

## Practice: create ingress resource based on domain

- apply the new definition: `kubectl apply -f devops-toolkit-dom.yml --record`
- send a domain-less request `nohup kubectl port-forward -n ingress-nginx service/ingress-nginx-controller 3000:80 --address 0.0.0.0 > /dev/null 2>&1 curl -i "http://0.0.0.0:3000"`

```bash
HTTP/1.1 404 Not Found
Server: nginx/1.15.9
Date: Wed, 19 Jun 2019 11:12:42 GMT
Content-Type: text/plain; charset=utf-8
Content-Length: 21
Connection: keep-alive
```

### Yaml
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: devops-toolkit
  annotations:
    kubernetes.io/ingress.class: "nginx"
    ingress.kubernetes.io/ssl-redirect: "false"
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - host: devopstoolkitseries.com
    http:
      paths:
      - path: /
        pathType: ImplementationSpecific
        backend:
          service:
            name: devops-toolkit
            port:
              number: 80

---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: devops-toolkit
spec:
  replicas: 3
  selector:
    matchLabels:
      type: frontend
      service: devops-toolkit
  template:
    metadata:
      labels:
        type: frontend
        service: devops-toolkit
    spec:
      containers:
      - name: frontend
        image: vfarcic/devops-toolkit-series

---

apiVersion: v1
kind: Service
metadata:
  name: devops-toolkit
spec:
  ports:
  - port: 80
  selector:
    type: frontend
    service: devops-toolkit
```

### Create a ingress resource with default backend

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: default
  annotations:
    kubernetes.io/ingress.class: "nginx"
    ingress.kubernetes.io/ssl-redirect: "false"
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - http:
      paths:
      - path: /
        pathType: ImplementationSpecific
        backend:
          service:
            name: devops-toolkit
            port:
              number: 80
```