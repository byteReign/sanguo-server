# 三国砍王 - 后端 API 文档

## 基础信息

- **基础URL**：`http://服务器IP:8081`
- **响应格式**：统一 JSON 格式
- **认证方式**：请求头 `X-Device-Id` 携带设备ID

---

## 统一响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

### 错误码

| code | 说明 |
|------|------|
| 0 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未登录 |
| 500 | 服务器错误 |
| 10001 | 用户不存在 |

---

## 接口列表

### 1. 登录/注册

**POST** `/api/auth/login`

**请求头**：
```
X-Device-Id: abc123def456
```

**请求体**：
```json
{
  "deviceId": "abc123def456"
}
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "userId": 1,
    "deviceId": "abc123def456",
    "nickname": "乡勇_456",
    "level": 1,
    "exp": 0,
    "clickTrueDamage": 1,
    "dpsTrueDamage": 1,
    "title": "乡勇",
    "totalDamage": 0,
    "totalKills": 0,
    "chapter": 1,
    "stage": 1,
    "lastOnlineTime": "2026-06-08T15:00:00Z"
  }
}
```

**说明**：
- 设备ID不存在则自动注册
- 新玩家初始等级=1，点击真伤=1，DPS真伤=1，称号=乡勇
- 从第1章第1关开始
- 全部真伤，无防御/减伤

---

### 2. 获取玩家信息

**GET** `/api/user/info`

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "userId": 1,
    "nickname": "乡勇_456",
    "level": 5,
    "exp": 120,
    "expToNext": 200,
    "clickTrueDamage": 15,
    "dpsTrueDamage": 8,
    "title": "伍长",
    "totalDamage": 12500,
    "totalKills": 52,
    "chapter": 1,
    "stage": 8,
    "unlockedIllustrations": 7,
    "totalIllustrations": 13
  }
}
```

**字段说明**：
- `clickTrueDamage`：点击真伤（升级+装备+称号累计）
- `dpsTrueDamage`：每秒真伤（升级+装备+称号累计）
- `expToNext`：升到下一级所需经验

---

### 3. 获取当前敌人信息

**GET** `/api/battle/current`

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "enemyId": 8,
    "enemyName": "八级黄巾兵",
    "enemyHp": 1200,
    "enemyMaxHp": 2500,
    "chapter": 1,
    "stage": 8,
    "isBoss": false
  }
}
```

---

### 4. 点击攻击

**POST** `/api/battle/attack`

**请求体**：
```json
{
  "damage": 15
}
```

**说明**：
- `damage` 由前端计算：点击真伤 + 临时 Buff
- 服务器会校验伤害是否合理（不能超过点击真伤 × 1.2）

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "enemyHp": 1185,
    "enemyMaxHp": 2500,
    "killed": false,
    "rewards": null
  }
}
```

**敌人被击杀时**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "enemyHp": 0,
    "enemyMaxHp": 2500,
    "killed": true,
    "rewards": {
      "powerGain": 4,
      "dpsGain": 4,
      "totalDamageGain": 2500,
      "newTitle": null,
      "newIllustration": {
        "id": 8,
        "name": "八级黄巾兵",
        "description": "黄巾起义军中的精锐士卒。"
      }
    },
    "nextEnemy": {
      "enemyId": 9,
      "enemyName": "九级黄巾兵",
      "enemyHp": 3500,
      "enemyMaxHp": 3500,
      "chapter": 1,
      "stage": 9,
      "isBoss": false
    }
  }
}
```

---

### 5. 离线收益结算

**POST** `/api/battle/offline`

**说明**：玩家上线时调用，计算离线期间的收益

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "offlineSeconds": 28800,
    "damageDealt": 86400,
    "enemiesKilled": 12,
    "powerGain": 36,
    "dpsGain": 36,
    "newTitle": "什长",
    "newIllustrations": [
      {
        "id": 10,
        "name": "十级黄巾兵",
        "description": "黄巾军中的老兵。"
      }
    ]
  }
}
```

**离线收益计算规则**：
- 离线 DPS = 在线 DPS真伤 × 50%（离线效率减半）
- 离线不获得点击伤害
- 离线收益最多计算 8 小时
- 离线期间获得的经验正常计算

---

### 6. 获取图鉴列表

**GET** `/api/illustration/list`

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 13,
    "unlocked": 7,
    "list": [
      {
        "id": 1,
        "name": "一级黄巾兵",
        "description": "东汉末年黄巾起义军士卒。",
        "unlocked": true,
        "unlockTime": "2026-06-02T10:30:00Z"
      },
      {
        "id": 2,
        "name": "二级黄巾兵",
        "description": "略通武艺的黄巾士卒。",
        "unlocked": true,
        "unlockTime": "2026-06-02T11:00:00Z"
      },
      {
        "id": 8,
        "name": "八级黄巾兵",
        "description": "黄巾起义军中的精锐士卒。",
        "unlocked": false,
        "unlockTime": null
      }
    ]
  }
}
```

---

### 7. 获取称号列表

**GET** `/api/title/list`

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "currentTitle": "伍长",
    "list": [
      {
        "name": "乡勇",
        "condition": "击败10名敌人",
        "requiredKills": 10,
        "dpsBonus": 1,
        "unlocked": true
      },
      {
        "name": "伍长",
        "condition": "击败50名敌人",
        "requiredKills": 50,
        "dpsBonus": 2,
        "unlocked": true
      },
      {
        "name": "什长",
        "condition": "击败100名敌人",
        "requiredKills": 100,
        "dpsBonus": 5,
        "unlocked": false
      }
    ]
  }
}
```

---

### 8. 排行榜

**GET** `/api/rank/damage`

**请求参数**：
- `page`：页码，默认 1
- `pageSize`：每页数量，默认 20

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 156,
    "page": 1,
    "pageSize": 20,
    "myRank": 23,
    "list": [
      {
        "rank": 1,
        "nickname": "将军_789",
        "totalDamage": 9999999
      },
      {
        "rank": 2,
        "nickname": "校尉_234",
        "totalDamage": 8888888
      }
    ]
  }
}
```

---

## 数据同步策略

### 实时同步
- 击杀敌人时同步
- 获得称号时同步
- 解锁图鉴时同步

### 定时同步
- 每 30 秒同步一次 DPS 和战力
- 每 60 秒同步一次总伤害

### 防作弊
- 服务器校验每次攻击伤害是否合理
- 攻击伤害不能超过（点击真伤 × 1.2）
- 异常数据直接拒绝并记录日志
- 关键数值（等级、真伤、关卡进度）以服务器为准
