#源镜像
FROM golang:1.15 as builder
#设置工作目录

RUN mkdir -p /http_server/
WORKDIR /http_server

#代理
ENV GO111MODLE=ON
ENV GOPROXY="https://goproxy.io"
ENV GOSUMDB="off"
ENV OS=linux
ENV ARCH=amd64

copy . .

RUN go mod download

RUN GOOS=${OS} GOARCH=${ARCH} go build  -a -o bin/server main.go

#FROM alpine:3.12

#WORKDIR /

#COPY  /http_server/bin/server .

#暴露端口
EXPOSE 8080
#最终运行docker的命令
ENTRYPOINT  ["./server"]
