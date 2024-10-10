# GoNav 后端

版本：1.0.0

这是 GoNav 项目的后端部分，使用 Go 和 Gin 框架构建。

## 环境要求

- Go（建议版本 1.16 或更高）
- SQLite3

## 安装和运行

1. 克隆项目后，进入后端目录：
   ```bash
   cd backend
   ```

2. 安装依赖：
   ```bash
   go mod tidy
   ```

3. 配置应用：
   - 打开 `config.yaml` 并根据需要调整设置。

4. 运行服务器：
   ```bash
   go run main.go
   ```

5. 服务器默认将在 `http://localhost:8080` 运行

## 配置

`config.yaml` 文件包含以下设置：

- `server`：
  - `port`: 服务器运行的端口（默认：8080）
  - `mode`: Gin 模式，可以是 "debug" 或 "release"
- `database`：
  - `type`: 数据库类型（目前仅支持 "sqlite")
  - `name`: 数据库文件名
  - `path`: 数据库文件路径
- `cors`：
  - `allowOrigins`: CORS 允许的源列表
- `jwt`：
  - `secret`: JWT 令牌生成的密钥
  - `expirationHours`: JWT 令牌过期时间（小时）

## 主要功能

- 用户认证
- 书签 CRUD 操作
- 分类管理
- 设置管理
- 公开书签分享
