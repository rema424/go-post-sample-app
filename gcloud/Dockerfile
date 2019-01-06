FROM golang:1.11 AS golang
FROM google/cloud-sdk:latest

# ARG sdk="1.9.64"

# golang multi-stage build
COPY --from=golang /usr/local/go /usr/local/go/
ADD service-account-key.json /tmp/service-account-key.json

# GOPATH設定
RUN mkdir -p /go/src/app
ENV GOPATH /go

# go PATH設定
ENV PATH $PATH:/usr/local/go/bin:/go/bin

# ソースコードのマウント
WORKDIR /go/src/app

RUN go get -u -v github.com/golang/dep/cmd/dep && \
  # dep ensure && \
  # go get -u -v google.golang.org/appengine/... && \
  echo "gcloud auth activate-service-account --key-file /tmp/service-account-key.json" >> /root/.bashrc
# gcloud auth activate-service-account --key-file /tmp/service-account-key.json && \
# rm -rf /tmp/service-account-key.json

# GAE GO SDKインストール
# RUN apt-get update
# RUN apt-get install unzip
# RUN curl -O https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-${sdk}.zip
# RUN unzip go_appengine_sdk_linux_amd64-${sdk}.zip -d /usr/local
# RUN rm go_appengine_sdk_linux_amd64-${sdk}.zip

# go PATH設定
# ENV PATH $PATH:/usr/local/go_appengine

# ポート解放
EXPOSE 8000
EXPOSE 8080