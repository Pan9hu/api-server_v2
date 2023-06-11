FROM golang:alpine as builder

WORKDIR /go/src/github.com/Pan9Hu/api-server-v2

COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o api-server .

FROM alpine:latest

LABEL author="Pan9Hu" \
      e-mail="Pan9Hu@Outlook.com" \
      version="2.0"

WORKDIR /go/src/github.com/Pan9Hu/api-server-v2

COPY --from=0 /go/src/github.com/Pan9Hu/api-server-v2/api-server ./
COPY --from=0 /go/src/github.com/Pan9Hu/api-server-v2/conf/api.properties ./

EXPOSE 8000

ENTRYPOINT ./api-server config api.properties
