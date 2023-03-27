### pods
- get metadata of pods `POD_NAME=$(kubectl get pods \
    -l service=jenkins,type=master \
    -o jsonpath="{.items[*].metadata.name}")`

### namespace
- get events: `kubectl --namespace dev get events`
- 