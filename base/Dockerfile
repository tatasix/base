# 使用 Alpine 版本的 Go 镜像
FROM golang:1.23-alpine

# 设置工作目录
WORKDIR /app

# 启用 Go Modules 并设置国内代理（解决网络问题）
ENV GOPROXY=https://goproxy.cn,direct

# 新增：安装时区数据并设置时区
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

# 复制 go.mod 和 go.sum 文件（利用 Docker 缓存，避免重复下载依赖）
COPY go.mod go.sum ./
RUN go mod download

# 复制整个项目代码
COPY . .

# 构建应用
RUN go build -o main cmd/main.go

# 暴露端口
EXPOSE 8897

# 运行应用
CMD ["./main"]