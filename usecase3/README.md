```
$ kubectl apply --filename blue-green-demo-config.yaml

$ export INGRESSGATEWAY=istio-ingressgateway
$ export IP_ADDRESS=$(kubectl get svc $INGRESSGATEWAY --namespace istio-system --output 'jsonpath={.status.loadBalancer.ingress[0].ip}')

### blue
# Configurationを設定したときにできるRevision名を取得
$ kubectl get configurations blue-green-demo -o=jsonpath='{.status.latestCreatedRevisionName}'

# 既存のRevisionに100%のトラフィックを流すRoute
# Route名は更新する
$ kubectl apply --filename blue-green-demo-route.yaml

$ curl -H "Host: blue-green-demo.default.example.com" http://${IP_ADDRESS}

### green
# config.yamlのTARGETをgreenに更新
$ kubectl apply --filename blue-green-demo-config.yaml
# Configurationの更新によって新しくできたRevision名を取得
$ kubectl get configurations blue-green-demo -o=jsonpath='{.status.latestCreatedRevisionName}'
# Routeに新Revision名追加、コメントイン
$ kubectl apply --filename blue-green-demo-route.yaml

# tag v2という名前をつけていて、テスト用のRouteができる
# 100%のblueと別にテスト用にアクセスできる
# テスト用Routeの確認
$ kubectl get route blue-green-demo --output jsonpath="{.status.traffic[*].url}"

$ curl -H "Host: v2-blue-green-demo.default.example.com" http://${IP_ADDRESS}

### blue50、green50
# Routeのトラフィックをどちらも50%に
$ kubectl apply --filename blue-green-demo-route.yaml

### blue0、green100
# blueを0にして tag: v1をつける
$ kubectl apply --filename blue-green-demo-route.yaml

＃ちょっと時間がかかる
$ kubectl get route blue-green-demo --output jsonpath="{.status.traffic[*].url}"

```