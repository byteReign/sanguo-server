# 编译阶段
FROM golang:1.26.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sg-server cmd/server/main.go


# 运行阶段（轻量）
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/sg-server .

EXPOSE 8080

CMD ["./sg-server"]