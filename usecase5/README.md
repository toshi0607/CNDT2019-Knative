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
