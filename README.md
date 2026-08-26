# 三走进实践协同平台

面向团区委、学校、街道/社区、企业、教师导师与家长的暑期社会实践履约后端，把项目发布、容量、报名授权、资质匹配、分组、签到、风险、材料和评价串成可恢复生命周期。

## 环境

- Go 1.23 或更高版本
- SQLite（由 modernc.org/sqlite 提供，无需系统数据库）

## 运行

```bash
go run .
```

默认监听 `:8080`，可通过 `HTTP_ADDR`、`DB_PATH`、`SESSION_TTL` 和 `WORKER_INTERVAL` 配置。

## 测试与检查

```bash
go test ./... -count=1
go test -race ./... -count=1
go vet ./...
go build ./...
```

## 主要 API

`POST /api/v1/auth/login` 登录，`POST /api/v1/auth/logout` 撤销会话；`/api/v1/projects` 管理实践项目；`/api/v1/projects/{id}/enroll` 报名；`/api/v1/enrollments/{id}/authorize` 监护授权；`/api/v1/groups` 匹配分组；`/api/v1/attendance/check-in` 与 `check-out` 记录履约；`/api/v1/risks` 管理风险；`/api/v1/submissions` 和 `/api/v1/evaluations` 完成材料与评价。

## Docker

```bash
docker build -t sanzoujin-practice .
docker run --rm -p 8080:8080 -v "$PWD/data:/app/data" sanzoujin-practice
```

健康检查为 `/healthz`，就绪检查为 `/readyz`。
