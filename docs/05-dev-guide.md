# 三国砍王 - 开发规范

## Git 分支策略

```
main          # 主分支，可部署版本
├── dev       # 开发分支
│   ├── feature/login      # 功能分支
│   ├── feature/battle     # 功能分支
│   └── fix/xxx            # 修复分支
```

### 提交规范

```
feat: 新增登录功能
fix: 修复血条显示异常
docs: 更新API文档
refactor: 重构战斗逻辑
```

---

## 后端编码规范

### 目录结构

```
sanguo-server/
├── cmd/server/          # 入口
├── internal/            # 业务逻辑（不对外暴露）
│   ├── auth/            # 认证模块
│   ├── user/            # 用户模块
│   ├── battle/          # 战斗模块
│   ├── illustration/    # 图鉴模块
│   ├── title/           # 称号模块
│   └── rank/            # 排行榜模块
├── pkg/                 # 公共组件
│   ├── config/          # 配置
│   ├── database/        # 数据库
│   ├── redis/           # Redis
│   ├── logger/          # 日志
│   ├── response/        # 响应
│   ├── errors/          # 错误
│   └── middleware/      # 中间件
└── router/              # 路由
```

### 代码风格

- 使用 `gofmt` 格式化
- 错误处理：优先返回 error，不 panic（除启动阶段）
- 命名：驼峰式（导出）/ 小驼峰（私有）

### 接口开发模板

```go
// handler.go
func (h *Handler) Attack(c *gin.Context) {
    var req AttackRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Fail(c, errors.CodeBadRequest, "参数错误")
        return
    }

    userId := middleware.GetUserId(c)
    result, err := h.service.Attack(userId, req.Damage)
    if err != nil {
        response.HandleError(c, err)
        return
    }

    response.Success(c, result)
}

// service.go
func (s *Service) Attack(userId int64, damage int64) (*AttackResult, error) {
    // 1. 校验伤害是否合理
    // 2. 计算敌人剩余血量
    // 3. 判断是否击杀
    // 4. 结算奖励
    // 5. 更新数据库
    // 6. 更新Redis排行榜
    return result, nil
}
```

---

## 前端编码规范

### 目录结构

```
NewProject/
├── assets/
│   ├── scripts/
│   │   ├── controllers/     # 场景控制器
│   │   ├── ui/              # UI组件
│   │   ├── models/          # 数据模型
│   │   └── utils/           # 工具函数
│   ├── prefabs/             # 预制体
│   ├── textures/            # 贴图
│   └── scenes/              # 场景
```

### 代码风格

- 使用 TypeScript
- 组件使用 `@ccclass` 装饰器
- 命名：大驼峰（类）/ 小驼峰（方法变量）

---

## 数据库规范

### 表命名

- 使用复数：`users`、`enemies`
- 关联表：`user_illustrations`

### 字段命名

- 小写下划线：`user_id`、`total_damage`
- 时间字段：`created_at`、`updated_at`
- 布尔字段：`is_boss`、`is_deleted`

### 索引

- 主键必建
- 外键必建索引
- 查询频繁字段建索引

---

## 配置管理

### 后端配置

```yaml
# configs/dev.yaml
server:
  host: 0.0.0.0
  port: 8081
  mode: debug

database:
  host: mysql
  port: 3306
  user: root
  password: root
  db: sanguo_server

redis:
  host: redis
  port: 6379
  password: ""
```

### 前端配置

```typescript
// utils/config.ts
export const CONFIG = {
  API_BASE_URL: 'http://服务器IP:8081',
  OFFLINE_MAX_SECONDS: 8 * 3600, // 离线最大8小时
  SYNC_INTERVAL: 30, // 数据同步间隔（秒）
};
```

---

## 安全规范

### 后端

1. **防作弊校验**：
   - 每次攻击校验伤害合理性
   - 伤害不能超过（战力 + DPS × 时间间隔）
   - 异常请求记录日志并拒绝

2. **SQL 注入防护**：
   - 使用 GORM 参数化查询
   - 不拼接 SQL

3. **限流**：
   - 攻击接口限流：每秒最多 10 次
   - 登录接口限流：每分钟最多 5 次

### 前端

1. **不信任客户端数据**：所有关键计算在服务器完成
2. **不存储敏感信息**：设备ID存在本地，其他数据从服务器获取

---

## 测试规范

### 后端测试

```bash
# 运行测试
go test ./...

# 测试覆盖率
go test -cover ./...
```

### 前端测试

- 微信开发者工具真机调试
- 测试不同分辨率设备
- 测试弱网环境

---

## 部署规范

### 后端

```bash
# 编译
go build -o sg-server cmd/server/main.go

# 运行
./sg-server -config configs/prod.yaml
```

### 前端

- 微信开发者工具上传代码
- 提交审核
- 发布

---

## 日志规范

### 后端日志

```go
// 信息日志
logger.Infof("user %d attacked enemy %d, damage %d", userId, enemyId, damage)

// 警告日志
logger.Warnf("user %d damage abnormal: %d", userId, damage)

// 错误日志
logger.WithError(err).Errorf("attack failed for user %d", userId)
```

### 日志级别

- `DEBUG`：开发环境，详细调试信息
- `INFO`：正常业务流程
- `WARN`：可恢复的异常
- `ERROR`：需要关注的错误
