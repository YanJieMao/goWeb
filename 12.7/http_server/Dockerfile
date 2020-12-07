#源镜像
FROM golang:alpine
#作者
MAINTAINER mkdirhao "mkdirhao@qq.com"
#设置工作目录
WORKDIR $GOPATH/src/goWeb/12.7/http_server
#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/goWeb/12.7/http_server
#代理
ENV GO111MODLE=ON
ENV GOPROXY="https://groproxy.io"
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 8080
#最终运行docker的命令
ENTRYPOINT  ["./http_server"]
