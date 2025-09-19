# 构建阶段
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o go-server .

# 运行阶段（轻量）
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/go-server .
EXPOSE 8080
CMD ["./go-server"]

