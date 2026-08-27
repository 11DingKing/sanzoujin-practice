# BENZHI_README

这是一个 Go 后端服务，面向团区委、学校、街道/社区、企业、教师导师与家长的暑期社会实践履约后端，把项目发布、容量、报名授权、资质匹配、分组、签到、风险、材料和评价串成可恢复生命周期。

## 标准构建、运行和测试命令

进入容器后执行：

```bash
# 编译
cd '/app' && GOTOOLCHAIN=local go build ./...

# 启动
cd '/app' && GOTOOLCHAIN=local go run .
cd '/app' && GOTOOLCHAIN=local go run ./cmd/server

# 测试
cd '/app' && GOTOOLCHAIN=local go test ./...
```

## Docker 构建和进入容器

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh benzhi-task-1-amd64 linux/amd64
./build_benzhi_docker.sh benzhi-task-1-arm64 linux/arm64
docker run -it benzhi-task-1-amd64:latest
docker run -it --platform linux/arm64 benzhi-task-1-arm64:latest
```
