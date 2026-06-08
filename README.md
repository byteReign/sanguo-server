# sanguo-server

## 项目结构与开发规范
``` 
sg-server/                              # 项目根目录（单体后端服务：三国砍王服务端）

├── cmd/
│   └── server/
│       └── main.go                    # 程序入口：负责启动整个服务
│                                      # 职责：加载配置 / 初始化DB-Redis-Logger / 注册路由 / 启动HTTP服务
│                                      # 禁止：写任何业务逻辑
│
├── configs/                           # 配置文件目录（按环境区分）
│   ├── dev.yaml                       # 开发环境配置（本地调试用）
│   ├── test.yaml                      # 测试环境配置
│   └── prod.yaml                      # 生产环境配置
│                                      # 内容：mysql / redis / jwt / server端口等
│
├── internal/                          # 核心业务代码（按业务模块拆分）
│   ├── user/                          # 用户系统（登录 / 注册 / 等级 / 经验）
│   ├── hero/                          # 武将系统（属性 / 获取 / 强化）
│   ├── skill/                         # 技能系统（技能数据 / 技能效果）
│   ├── battle/                        # 战斗系统（核心玩法：战斗结算）
│   ├── bag/                           # 背包系统（道具 / 物品）
│   ├── title/                         # 称号系统（达成条件 / 称号解锁）
│   ├── achievement/                   # 成就系统（任务型目标）
│   ├── task/                          # 任务系统（日常 / 主线任务）
│   └── ranking/                       # 排行榜系统（战力 / 等级排名）
│
│   # 每个模块统一结构如下：
│   # model.go       -> 数据结构（对应数据库表）
│   # repository.go  -> 数据层（只负责 MySQL/Redis 读写）
│   # service.go     -> 业务层（核心逻辑都在这里）
│   # handler.go     -> 接口层（HTTP入口）
│   # dto.go         -> 请求/响应结构体
│
├── pkg/                               # 公共基础组件（无业务逻辑）
│   ├── database/                      # MySQL初始化 & 连接池封装
│   ├── redis/                         # Redis初始化 & 基础封装
│   ├── logger/                        # 日志系统封装
│   ├── jwt/                           # Token生成与解析
│   ├── response/                      # 统一API返回格式封装
│   ├── errors/                        # 统一错误定义
│   └── util/                          # 工具函数（时间/加密/字符串等）
│
├── router/
│   └── router.go                      # 路由注册中心
│                                      # 职责：把 user/hero/battle 等模块 handler 统一挂载成 HTTP API
│
├── migrations/                        # 数据库版本管理（建表 & 字段变更SQL）
│                                      # 示例：001_user.sql / 002_hero.sql
│
├── docs/                              # 项目文档
│                                      # API说明 / 游戏设计 / 表结构 / 规则说明
│
├── scripts/                           # 运维脚本
│                                      # 启动 / 部署 / 初始化 / 数据迁移脚本
│
├── go.mod                             # Go依赖管理
└── go.sum                             # 依赖校验文件
```

## 数据流转
``` 
Client
↓
Handler（接收请求 + 参数校验）
↓
Service（业务逻辑）
↓
Repository（数据读写）
↓
MySQL / Redis
```