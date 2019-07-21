```
$ go get -u github.com/rakyll/hey

$ kubectl apply --filename service.yaml
$ export INGRESSGATEWAY=istio-ingressgateway
$ export IP_ADDRESS=$(kubectl get svc $INGRESSGATEWAY --namespace istio-system --output 'jsonpath={.status.loadBalancer.ingress[0].ip}')

$ hey -z 30s -c 50 \
  -host "autoscale-go.default.example.com" \
  "http://${IP_ADDRESS?}?sleep=100&prime=10000&bloat=5" \
  && kubectl get pods

# 50 並行リクエスト / Podあたり10 = 5 pods

# メトリクスをgrafanaで確認
$ kubectl port-forward --namespace knative-monitoring $(kubectl get pods --namespace knative-monitoring --selector=app=grafana  --output=jsonpath="{.items..metadata.name}") 3000

# Knarive Serving - Scaling Debuging
# Knarive Serving - Revision HTTP Requests

# 片付け
$ kubectl delete --filename service.yaml
```