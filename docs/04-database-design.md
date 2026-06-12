# 三国砍王 - 数据库设计

## 概述

- **数据库类型**：MySQL 8.0
- **字符集**：utf8mb4
- **命名规范**：小写下划线，复数表名

---

## 表结构

### 1. users - 用户表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT UNSIGNED AUTO_INCREMENT | 主键 |
| device_id | VARCHAR(64) UNIQUE | 设备ID |
| nickname | VARCHAR(32) | 昵称 |
| level | INT UNSIGNED DEFAULT 1 | 等级 |
| exp | BIGINT UNSIGNED DEFAULT 0 | 经验值 |
| click_true_damage | BIGINT UNSIGNED DEFAULT 1 | 点击真伤 |
| dps_true_damage | BIGINT UNSIGNED DEFAULT 1 | 每秒真伤（DPS） |
| title | VARCHAR(32) DEFAULT '乡勇' | 当前称号 |
| total_damage | BIGINT UNSIGNED DEFAULT 0 | 总伤害（排行榜用） |
| total_kills | INT UNSIGNED DEFAULT 0 | 总击杀数 |
| chapter | INT UNSIGNED DEFAULT 1 | 当前章节 |
| stage | INT UNSIGNED DEFAULT 1 | 当前关卡 |
| last_online_time | DATETIME | 最后在线时间 |
| created_at | DATETIME DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| updated_at | DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE | 更新时间 |

**说明**：
- `click_true_damage`：点击造成的真伤，升级/装备/称号都会增加
- `dps_true_damage`：每秒自动造成的真伤，升级/装备/称号都会增加
- 前期全部真伤，无防御/减伤机制
- 后续加防御系统时，只需在伤害计算层处理，不影响存储结构

**索引**：
- PRIMARY KEY (id)
- UNIQUE KEY (device_id)
- INDEX (total_damage) -- 排行榜用

---

### 2. illustrations - 图鉴表（玩家已解锁）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT UNSIGNED AUTO_INCREMENT | 主键 |
| user_id | BIGINT UNSIGNED | 用户ID |
| illustration_id | INT UNSIGNED | 图鉴ID（敌人ID） |
| unlock_time | DATETIME DEFAULT CURRENT_TIMESTAMP | 解锁时间 |
| created_at | DATETIME DEFAULT CURRENT_TIMESTAMP | 创建时间 |

**索引**：
- PRIMARY KEY (id)
- UNIQUE KEY (user_id, illustration_id)
- INDEX (user_id)

---

### 3. illustration_defs - 图鉴定义表（静态数据）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | INT UNSIGNED AUTO_INCREMENT | 主键（图鉴ID） |
| enemy_id | INT UNSIGNED | 对应敌人ID |
| name | VARCHAR(64) | 图鉴名称 |
| description | VARCHAR(256) | 描述 |
| sort_order | INT UNSIGNED | 排序 |

**索引**：
- PRIMARY KEY (id)

---

### 4. titles - 称号定义表（静态数据）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | INT UNSIGNED AUTO_INCREMENT | 主键 |
| name | VARCHAR(32) | 称号名称 |
| required_kills | INT UNSIGNED | 需要击杀数 |
| dps_bonus | INT UNSIGNED DEFAULT 0 | DPS加成 |
| sort_order | INT UNSIGNED | 排序 |

**索引**：
- PRIMARY KEY (id)

---

### 5. level_defs - 等级定义表（静态数据）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | INT UNSIGNED AUTO_INCREMENT | 主键（等级） |
| required_exp | BIGINT UNSIGNED | 升到该级所需总经验 |
| exp_to_next | BIGINT UNSIGNED | 升到下一级所需经验 |
| click_true_damage | INT UNSIGNED DEFAULT 0 | 升级奖励：点击真伤 |
| dps_true_damage | INT UNSIGNED DEFAULT 0 | 升级奖励：每秒真伤 |
| sort_order | INT UNSIGNED | 排序 |

**索引**：
- PRIMARY KEY (id)

**说明**：
- `required_exp`：累计经验，用于判断是否升级
- `exp_to_next`：当前级升到下一级还差多少经验
- 升级奖励直接加真伤，简单直观

**示例数据**：

```sql
INSERT INTO level_defs (id, required_exp, exp_to_next, click_true_damage, dps_true_damage, sort_order) VALUES
(1,  0,       50,    1, 1, 1),
(2,  50,      100,   1, 1, 2),
(3,  150,     200,   2, 1, 3),
(4,  350,     350,   2, 2, 4),
(5,  700,     500,   3, 2, 5),
(6,  1200,    800,   3, 3, 6),
(7,  2000,    1200,  4, 3, 7),
(8,  3200,    1800,  4, 4, 8),
(9,  5000,    2500,  5, 4, 9),
(10, 7500,    3500,  5, 5, 10),
(11, 11000,   5000,  6, 5, 11),
(12, 16000,   7000,  6, 6, 12),
(13, 23000,   10000, 7, 6, 13),
(14, 33000,   14000, 7, 7, 14),
(15, 47000,   20000, 8, 7, 15),
(16, 67000,   28000, 8, 8, 16),
(17, 95000,   40000, 9, 8, 17),
(18, 135000,  55000, 9, 9, 18),
(19, 190000,  75000, 10, 9, 19),
(20, 265000,  100000,10,10, 20);
```

---

### 6. enemies - 敌人定义表（静态数据）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | INT UNSIGNED AUTO_INCREMENT | 主键（敌人ID） |
| name | VARCHAR(64) | 敌人名称 |
| hp | BIGINT UNSIGNED | 血量 |
| reward_exp | INT UNSIGNED | 奖励经验 |
| chapter | INT UNSIGNED | 所属章节 |
| stage | INT UNSIGNED | 所属关卡 |
| is_boss | TINYINT(1) DEFAULT 0 | 是否Boss |
| description | VARCHAR(256) | 描述 |
| sort_order | INT UNSIGNED | 排序 |

**索引**：
- PRIMARY KEY (id)
- INDEX (chapter, stage)

---

### 6. chapters - 章节定义表（静态数据）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | INT UNSIGNED AUTO_INCREMENT | 主键 |
| name | VARCHAR(64) | 章节名称 |
| description | VARCHAR(256) | 描述 |
| sort_order | INT UNSIGNED | 排序 |

**索引**：
- PRIMARY KEY (id)

---

## 初始化数据

### chapters 表

```sql
INSERT INTO chapters (id, name, description, sort_order) VALUES
(1, '黄巾之乱', '东汉末年，黄巾起义，天下大乱。', 1);
```

### enemies 表（黄巾之乱章节）

```sql
-- 黄巾士卒（普通怪）
INSERT INTO enemies (id, name, hp, reward_exp, chapter, stage, is_boss, description, sort_order) VALUES
(1, '一级黄巾兵', 100, 10, 1, 1, 0, '东汉末年黄巾起义军士卒。', 1),
(2, '二级黄巾兵', 200, 15, 1, 2, 0, '略通武艺的黄巾士卒。', 2),
(3, '三级黄巾兵', 350, 20, 1, 3, 0, '经验丰富的黄巾老兵。', 3),
(4, '四级黄巾兵', 500, 30, 1, 4, 0, '黄巾军中的小头目。', 4),
(5, '五级黄巾兵', 800, 40, 1, 5, 0, '黄巾军中的精锐。', 5),
(6, '六级黄巾兵', 1200, 50, 1, 6, 0, '黄巾军中的老兵。', 6),
(7, '七级黄巾兵', 1800, 60, 1, 7, 0, '黄巾军中的悍卒。', 7),
(8, '八级黄巾兵', 2500, 80, 1, 8, 0, '黄巾起义军中的精锐士卒。', 8),
(9, '九级黄巾兵', 3500, 100, 1, 9, 0, '黄巾军中的精英战士。', 9),
(10, '十级黄巾兵', 5000, 120, 1, 10, 0, '黄巾军中的百战老兵。', 10);

-- 黄巾渠帅（精英怪）
INSERT INTO enemies (id, name, hp, reward_exp, chapter, stage, is_boss, description, sort_order) VALUES
(11, '一级渠帅', 8000, 300, 1, 11, 0, '黄巾军中的渠帅，统领一方。', 11),
(12, '二级渠帅', 15000, 500, 1, 12, 0, '经验丰富的黄巾渠帅。', 12),
(13, '三级渠帅', 25000, 800, 1, 13, 0, '黄巾军中的大渠帅。', 13);

-- 黄巾首领（Boss）
INSERT INTO enemies (id, name, hp, reward_exp, chapter, stage, is_boss, description, sort_order) VALUES
(14, '张宝', 50000, 2000, 1, 14, 1, '地公将军，黄巾军首领之一。', 14),
(15, '张梁', 80000, 3000, 1, 15, 1, '人公将军，黄巾军首领之一。', 15),
(16, '张角', 150000, 5000, 1, 16, 1, '天公将军，黄巾起义领袖。', 16);
```

### illustration_defs 表

```sql
-- 与 enemies 表一一对应
INSERT INTO illustration_defs (id, enemy_id, name, description, sort_order)
SELECT id, id, name, description, sort_order FROM enemies;
```

### titles 表

```sql
INSERT INTO titles (id, name, required_kills, dps_bonus, sort_order) VALUES
(1, '乡勇', 10, 1, 1),
(2, '伍长', 50, 2, 2),
(3, '什长', 100, 5, 3),
(4, '百夫长', 500, 10, 4),
(5, '千夫长', 1000, 20, 5),
(6, '校尉', 5000, 50, 6),
(7, '将军', 10000, 100, 7);
```

---

## 数据流转说明

```
击杀敌人
    ↓
获得经验（reward_exp）
    ↓
经验达到 required_exp → 升级
    ↓
升级获得真伤（click_true_damage + dps_true_damage）
    ↓
真伤累加到用户表
```

**核心逻辑**：
- 敌人 → 奖励经验
- 等级表 → 定义升级条件和升级奖励
- 用户表 → 存储当前等级、经验、累计真伤
- 称号/装备 → 额外提供真伤加成

---

## 数据同步策略

### Redis 缓存

| Key | 类型 | 说明 |
|-----|------|------|
| `user:{userId}` | Hash | 用户信息缓存 |
| `rank:damage` | ZSET | 总伤害排行榜 |
| `user:device:{deviceId}` | String | deviceId → userId 映射 |

### 缓存更新

- 用户登录：写入 Hash + ZSET
- 击杀敌人：更新 Hash + ZSET
- 定时刷新：每 5 分钟持久化到 MySQL
