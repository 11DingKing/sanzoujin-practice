# 官方 Go 镜像，自带完整工具链
FROM golang:1.23-bookworm
WORKDIR /app
COPY go.mod go.sum ./
RUN GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go mod download
COPY . .
RUN GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go build ./...
CMD ["bash"]
