# 第一阶段：构建阶段
FROM golang:1.24.3-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的构建工具和依赖
RUN apk add --no-cache git make gcc musl-dev

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 设置构建参数（可注入版本信息等）
ARG APP_VERSION=1.0.0
ARG BUILD_TIME
ARG COMMIT_SHA

# 构建应用（启用静态链接、禁用CGO、优化二进制大小）
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "\
    -w -s \
    -X main.version=${APP_VERSION} \
    -X main.buildTime=${BUILD_TIME} \
    -X main.commit=${COMMIT_SHA}" \
    -a -installsuffix cgo -o /go/bin/app .

# 第二阶段：运行时镜像
FROM alpine:3.18

# 安装必要的运行时依赖
RUN apk add --no-cache ca-certificates tzdata && \
    update-ca-certificates

# 设置时区
ENV TZ=Asia/Shanghai

# 创建非root用户
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder --chown=appuser:appgroup /go/bin/app /app/app
COPY --from=builder --chown=appuser:appgroup /app/config.yaml /app/config.yaml

# 设置健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# 设置容器用户
USER appuser

# 暴露端口
EXPOSE 8080

# 设置容器启动命令
ENTRYPOINT ["/app/app"]