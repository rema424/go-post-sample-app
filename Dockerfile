FROM golang:1.11 AS golang
FROM google/cloud-sdk:latest

# golang multi-stage build
COPY --from=golang /usr/local/go /usr/local/go
COPY --from=golang /usr/bin/make /usr/local/bin/
ADD ./gcloud/service-account-key.json /tmp/service-account-key.json

# GOPATH設定
RUN mkdir -p /go/src/app
ENV GOPATH /go

# go PATH設定
ENV PATH $PATH:/usr/local/go/bin:/go/bin

# goappコマンドにPATHを通す
ENV PATH $PATH:/usr/lib/google-cloud-sdk/platform/google_appengine

# ソースコードのマウント
ADD ./src/app /go/src/app
WORKDIR /go/src/app

RUN chmod +x /usr/lib/google-cloud-sdk/platform/google_appengine/goapp && \
  chmod +x /usr/lib/google-cloud-sdk/platform/google_appengine/appcfg.py && \
  # curl -sL https://deb.nodesource.com/setup_10.x | bash - && \
  # apt-get install -y nodejs && \
  # go get -u -v github.com/golang/dep/cmd/dep && \
  # dep ensure && \
  echo "gcloud auth activate-service-account --key-file /tmp/service-account-key.json" >> /root/.bashrc && \
  echo 'export PROJECT_ID=$(gcloud config get-value account | cut -d "@" -f 2 | cut -d "." -f 1)' >> /root/.bashrc && \
  echo 'gcloud config set project $PROJECT_ID' >> /root/.bashrc

# ポート解放
EXPOSE 8000
EXPOSE 8080