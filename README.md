# Go API Practice

Golang 後端 API 練習專案，使用 Gin + PostgreSQL + plain SQL。  
**路由與框架相關程式已寫好。商業邏輯 = 如何從 DB 查、寫、改、刪，實作在 `services/`，每個函式上方都有「實作要點」註解說明。**

## 環境需求

- Go 1.21+
- PostgreSQL

## 設定

1. 建立 PostgreSQL 資料庫（例如名稱 `practice`）。
2. 複製並編輯環境變數：
   ```bash
   cp .env.example .env
   ```
3. `.env` 可設定：
   - `PORT`：API 監聽 port（預設 `8080`）
   - `DATABASE_URL`：PostgreSQL 連線字串

## 執行

```bash
go mod tidy
go run main.go
```

伺服器預設在 `http://localhost:8080`。

## 專案結構

| 目錄／檔案 | 說明 |
|------------|------|
| `handlers/` | 只處理 HTTP：解析參數、綁定 JSON、呼叫 service、回傳狀態與 JSON |
| `services/` | **商業邏輯實作處**：對 DB 的查詢與寫入（Query / QueryRow / Exec、Scan、錯誤處理），每個函式有註解說明要實作的內容 |
| `models/` | 請求／回應用的 struct |
| `database/` | 連線與建表 |
| `routes/` | 註冊路由 |
| `postman/` | Postman collection 匯入後可測試所有 API |

## API 路由（路徑簡單、一致）

### Items

| 方法 | 路徑 | 說明 |
|------|------|------|
| GET | `/api/items` | 取得所有項目 |
| GET | `/api/items/:id` | 依 ID 取得單一項目 |
| POST | `/api/items` | 新增項目 |
| PUT | `/api/items/:id` | 更新項目 |
| DELETE | `/api/items/:id` | 刪除項目 |

### Users

| 方法 | 路徑 | 說明 |
|------|------|------|
| GET | `/api/users` | 取得所有使用者 |
| GET | `/api/users/:id` | 依 ID 取得單一使用者 |
| POST | `/api/users` | 新增使用者 |
| PUT | `/api/users/:id` | 更新使用者 |
| DELETE | `/api/users/:id` | 刪除使用者 |

### Categories

| 方法 | 路徑 | 說明 |
|------|------|------|
| GET | `/api/categories` | 取得所有分類 |
| GET | `/api/categories/:id` | 依 ID 取得單一分類 |
| POST | `/api/categories` | 新增分類 |
| PUT | `/api/categories/:id` | 更新分類 |
| DELETE | `/api/categories/:id` | 刪除分類 |

## 練習重點：商業邏輯（怎麼查 DB、寫 DB）

在 `services/` 各檔案中，每個函式上方都有 **「實作要點」** 註解，說明該函式要做的 DB 操作，例如：

- **查多筆**：`database.DB.Query(...)` → `rows.Next()` + `rows.Scan(&欄位...)` → 組成 slice 回傳
- **查單筆**：`database.DB.QueryRow("... WHERE id = $1", id).Scan(...)`；若 `err == sql.ErrNoRows` 則回傳 `ErrNotFound`
- **新增**：`INSERT ... RETURNING ...` 搭配 `QueryRow(...).Scan(...)` 取得寫入後的資料
- **更新／刪除**：`DB.Exec(...)`，用 `result.RowsAffected()` 若為 0 則回傳 `ErrNotFound`

錯誤回傳 `services.ErrNotFound` 或 `services.ErrInvalidInput` 時，handler 會自動對應 HTTP 404 / 400。
