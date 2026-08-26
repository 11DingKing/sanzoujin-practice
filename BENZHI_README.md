# BENZHI_README

三走进实践协同平台面向团区委、学校、街道社区、企业导师与家长，提供暑期实践项目从发布、报名授权、分组匹配到履约签到、风险处置、材料审核和评价的可恢复协作服务。

## 标准构建、运行和测试命令

进入容器后执行：

```bash
GOTOOLCHAIN=local go build ./...
GOTOOLCHAIN=local go run .
GOTOOLCHAIN=local go test ./... -count=1
```

## Docker 构建和进入容器

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh sanzoujin-practice-amd64 linux/amd64
./build_benzhi_docker.sh sanzoujin-practice-arm64 linux/arm64
docker run -it sanzoujin-practice-amd64
docker run -it --platform linux/arm64 sanzoujin-practice-arm64
```
