# GoNav 後端

版本：1.0.0

這是 GoNav 專案的後端部分，使用 Go 和 Gin 框架構建。

## 環境要求

- Go（建議版本 1.16 或更高）
- SQLite3

## 安裝與運行

1. 克隆專案後，進入後端目錄：

   ```bash
   cd backend
   ```
2. 安裝依賴：

   ```bash
   go mod tidy
   ```
3. 配置應用：

   - 打開 `config.yaml` 並根據需要調整設置。
4. 運行伺服器：

   ```bash
   go run main.go
   ```
5. 伺服器預設將在 `http://localhost:8080` 運行

## 配置

`config.yaml` 文件包含以下設置：

- `server`：
  - `port`: 伺服器運行的端口（預設：8080）
  - `mode`: Gin 模式，可以是 "debug" 或 "release"
- `database`：
  - `type`: 資料庫類型（目前僅支援 "sqlite"）
  - `name`: 資料庫文件名
  - `path`: 資料庫文件路徑
- `cors`：
  - `allowOrigins`: CORS 允許的來源列表
- `jwt`：
  - `secret`: JWT 令牌生成的密鑰
  - `expirationHours`: JWT 令牌過期時間（小時）

## 主要功能

- 用戶認證
- 書籤 CRUD 操作
- 分類管理
- 設置管理
- 公開書籤分享
