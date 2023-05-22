# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.20-rc-alpine as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY .. .

# CGO_ENABLED禁用cgo 然后指定OS等，并go build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wechat_monitor .


# 运行阶段指定scratch作为基础镜像
FROM scratch

WORKDIR /app

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /app/wechat_monitor .
COPY --from=builder /app/etc.yaml .

EXPOSE 8088

ENTRYPOINT ["./wechat_monitor"]