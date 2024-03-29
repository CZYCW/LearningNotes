export INGRESS_NAME=istio-ingressgateway
export INGRESS_NS=istio-system
export NAMESPACE=kserve-test
export MY_INGRESS_GATEWAY_HOST=istio.$NAMESPACE.bookinfo.com
echo $MY_INGRESS_GATEWAY_HOST

kubectl get svc "$INGRESS_NAME" -n "$INGRESS_NS"
export INGRESS_HOST=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export SECURE_INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="https")].port}')
export TCP_INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="tcp")].port}')
echo $INGRESS_HOST $MY_INGRESS_GATEWAY_HOST $INGRESS_PORT
curl -s $MY_INGRESS_GATEWAY_HOST:$INGRESS_PORT/productpage 
