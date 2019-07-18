```
# Tektonインストール

# GCR用のサービスアカウントを作成してSecretとして保存
# https://github.com/GoogleContainerTools/kaniko/blob/master/README.md#kubernetes-secret
$ gcloud iam service-accounts create gcr-push
$ export PROJECT_ID=xxxxxxxxx
$ gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member=serviceAccount:gcr-push@$PROJECT_ID.iam.gserviceaccount.com \
  --role roles/storage.admin
$ gcloud iam service-accounts keys create gcr-push.json \
  --iam-account=gcr-push@$PROJECT_ID.iam.gserviceaccount.com
$ kubectl create secret generic kaniko-secret --from-file=gcr-push.json

# テンプレート登録
$ kubectl apply -f resource.yaml
$ kubectl apply -f task.yaml
$ kubectl apply -f pipeline.yaml
$ kubectl get tekton-pipelines

# CI実行
$ kubectl apply -f pipelinerun.yaml

# 確認
$ kubectl describe pipelinerun pipeline-run

# 片づけ
$ kubectl delete -f service.yaml
$ kubectl delete -f resource.yaml
$ kubectl delete -f task.yaml
$ kubectl delete -f pipeline.yaml
$ kubectl delete -f pipelinerun.yaml
```