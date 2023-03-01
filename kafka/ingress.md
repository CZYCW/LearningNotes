## Ingress

### Overview
Kubernetes's ingress controler is not provided by kubernetes. Kubernetes only provides its api, you need to install another controller for that.

### In Practice
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
  annotations:
    kubernetes.io/ingress.class: "nginx"
    ingress.kubernetes.io/ssl-redirect: "false"
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - http:
      paths:
      - path: /demo
        pathType: ImplementationSpecific
        backend:
          service:
            name: go-demo-2-api
            port:
              number: 8080
```