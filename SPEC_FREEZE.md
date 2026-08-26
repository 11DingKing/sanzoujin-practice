# 三走进实践平台规格冻结表

## 基础项目档位与边界

| 项目 | 冻结结论 |
|---|---|
| foundation_profile | compact_10 |
| task_count | 10 |
| 项目形态 | backend |
| 业务边界 | 团区委、学校、街道/社区、企业、教师导师与家长协作完成暑期实践项目发布、报名授权、资质匹配、履约签到、风险处置、材料提交和评价 |
| 业务路径 | 项目发布容量管理、报名授权、匹配成组、签到签退、风险恢复、材料审核评价 |
| 禁止题材筛查 | 不属于游戏、桌面工具、订单库存、权限后台、考勤 OA、报表看板或无边界预约系统；签到只是实践履约中的一个步骤 |

## 持久化与一致性

| 项目 | 冻结结论 |
|---|---|
| 数据库 | SQLite（modernc.org/sqlite，真实 SQL，WAL 模式） |
| migration | `migrations/001_init.sql`，schema_migrations 版本表，启动时事务式升级，历史冲突阻断 |
| 关系表 | users、sessions、practice_projects、venues、enrollments、groups、group_members、attendance、risk_events、submissions、evaluations、idempotency_keys、audit_events、outbox_messages、worker_jobs、schema_migrations |
| 事务 | 报名+名额锁定+审计；匹配成组+成员转移+审计；签退+时长结算+审计；风险关闭+项目状态+通知；材料审核+评价窗口 |
| 并发 | `UPDATE ... WHERE capacity_used < capacity`、项目 version 乐观锁、唯一索引和事务锁；冲突返回稳定错误码 |
| 恢复 | outbox/worker_jobs 持久化重试；启动扫描 pending/running 任务并恢复；重启集成测试验证状态保留 |

## 运行时和接口

- service、repository、worker API 接收并传播 context；超时与取消保留错误链。
- 会话采用随机 Token 哈希存储，登录、退出撤销、过期清理和 admin/coordinator/student 三类角色均有 HTTP 验证。
- HTTP 入口提供 JSON 错误、请求 ID、panic recovery、结构化日志、`/healthz` 与验证数据库的 `/readyz`。
- 通知使用 outbox、指数退避、最大重试和永久失败审计；幂等键绑定租户、方法和路径。

## 测试与规模门禁

- 覆盖领域状态机、service 跨实体事务、真实数据库 migration、HTTP 契约、并发容量、幂等、重启、worker 重试/取消、分页过滤、时间边界和错误映射。
- 执行 `go test ./... -count=1`、`go test -race ./... -count=1`、`go vet ./...`、`go build ./...`、measure 脚本及 Docker 双架构构建/启动健康检查。
- compact_10 门禁：非测试生产 Go ≥2000 物理行、≥20 文件、≥10 package；测试 Go ≥1500 物理行。

## 后续容量（只规划正确边界，不预埋缺陷）

1. 项目发布与场所容量条件更新
2. 报名幂等键生命周期
3. 监护授权撤回与候补转正
4. 资质与时间窗口匹配
5. 分组成员并发转移
6. 教师/导师协作确认
7. 集合出发、迟到与缺席时序
8. 签到签退及服务时长结算
9. 天气/场所关闭风险恢复
10. 风险事件升级、关闭和通知重试
11. 调研材料版本提交与审核
12. 多方评价窗口与状态约束
13. 审计哈希链和不可抵赖查询
14. 分页过滤和组合查询
15. 过期会话撤销与权限边界
16. worker 重启续作
17. 批量操作部分失败
18. 数据库 migration 升级冲突
19. context 取消传播
20. 错误码/错误链映射
