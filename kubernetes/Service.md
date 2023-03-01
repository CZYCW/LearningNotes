# Service

## Creating a service

1. Create ReplicaSet: `kubectl create -f go-demo-2-db-rs.yml`
```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: go-demo-2-db
spec:
  selector:
    matchLabels:
      type: db
      service: go-demo-2
  template:
    metadata:
      labels:
        type: db
        service: go-demo-2
        vendor: MongoLabs
    spec:
      containers:
      - name: db
        image: mongo:3.3
        ports:
        - containerPort: 28017
```
2. Create Service: `kubectl create -f go-demo-2-db-svc.yml`
```yaml 
apiVersion: v1
kind: Service 
metadata:
  name: go-demo-2-db
spec:
  ports:
  - port: 27017
  selector:
    type: db
    service: go-demo-2
```
- there is no type, by default it is ClusterIP
- there is no protocol, by default it is TCP