```
# 事前準備
$ kubectl apply --filename https://github.com/knative/eventing-contrib/releases/download/v0.7.0/gcppubsub.yaml
$ gcloud services enable pubsub.googleapis.com

$ gcloud iam service-accounts create knative-source
$ export PROJECT_ID=xxxxxxxxx
$ gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member=serviceAccount:knative-source@$PROJECT_ID.iam.gserviceaccount.com \
  --role roles/pubsub.editor

$ gcloud iam service-accounts keys create knative-source.json \
  --iam-account=knative-source@$PROJECT_ID.iam.gserviceaccount.com

$ kubectl --namespace knative-sources create secret generic gcppubsub-source-key --from-file=key.json=knative-source.json --dry-run --output yaml | kubectl apply --filename -

$ kubectl --namespace default create secret generic google-cloud-key --from-file=key.json=knative-source.json


$ kubectl label namespace default knative-eventing-injection=enabled
namespace/default labeled

$ gcloud pubsub topics create eventing-test-topic
Created topic [projects/xxxxxxx/topics/eventing-test-topic].

# demo
$ kubectl apply --filename gcp-pubsub-source.yaml
$ kubectl apply --filename trigger.yaml

$ gcloud pubsub topics publish eventing-test-topic --message="Hello world"

$ kubectl get pods --selector serving.knative.dev/service=event-display
$ kubectl logs --selector serving.knative.dev/service=event-display -c user-container

# メッセージの確認
# echo "" | base64 --decode
# https://play.golang.org/p/VsXuFwHrB1f

# 片付け
$ kubectl delete --filename gcp-pubsub-source.yaml
$ kubectl delete --filename trigger.yaml
```
