# 第 1 阶段：构建阶段
FROM golang:1.22-alpine3.19 AS builder

# 设置 Go 模块代理
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 仅复制依赖相关文件
COPY go.mod go.sum ./

# 下载依赖（利用缓存）
RUN go mod download

# 复制剩余代码
COPY . .

# 编译 Go 程序
RUN go build -o /app/bin/cmd ./api/Image/image.go

# 第 2 阶段：运行阶段（轻量级）
FROM alpine:3.19

# 添加必要工具
RUN apk add --no-cache curl

# 设置工作目录
WORKDIR /app

# 复制必要的运行文件
COPY --from=builder /app/bin/cmd /app/bin/cmd
COPY --from=builder /app/api/Image/etc /app/etc

# 暴露服务端口
EXPOSE 8527

# 定义容器启动时运行的命令
CMD ["/app/bin/cmd", "-f", "/app/etc/image.yaml"]