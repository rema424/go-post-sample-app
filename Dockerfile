FROM golang:1.11

RUN mkdir -p /go/src/app && \
  : 'パッケージ管理' && \
  go get -u -v github.com/golang/dep/cmd/dep && \
  : 'デバッガー' && \
  go get -u -v github.com/derekparker/delve/cmd/dlv && \
  : 'ホットリロード' && \
  : 'fresh' && \
  go get -u -v github.com/pilu/fresh && \
  : 'gin -p 3000 -a 8080 run main.go' && \
  go get -u -v github.com/codegangsta/gin && \
  : 'realize start --no-config --run' && \
  go get -u -v github.com/oxequa/realize

ADD src /go/src
WORKDIR /go/src
# RUN dep ensure

EXPOSE 8080
CMD ["go", "run", "main.go"]