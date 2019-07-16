
```
$ kubectl apply --filename service.yaml
$ kubectl get svc istio-ingressgateway --namespace istio-system

# EXTERNAL-IPを取得
# minikubeなど外部LBがない場合は NodeIP や NodePort を利用
$ export IP_ADDRESS=$(kubectl get svc $INGRESSGATEWAY --namespace istio-system --output 'jsonpath={.status.loadBalancer.ingress[0].ip}')

# host URLを取得
# カスタムドメインももちろん利用可能
$ kubectl get route helloworld-go
http://helloworld-go.default.example.com



# 片付け
$ kubectl delete --filename service.yaml
```