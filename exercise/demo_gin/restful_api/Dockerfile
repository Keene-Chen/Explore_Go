#源镜像
FROM golang:alpine
 
#将二进制文件拷贝进容器的GOPATH目录中
ADD restful_api /go/src/restful_api
 
#为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
 
#暴露端口
EXPOSE 8080
 
#工作目录
WORKDIR /go/src
 
#最终运行docker的命令 #此处运行容器内执行# CMD ["./restful_api"]
ENTRYPOINT ./restful_api