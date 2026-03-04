package handlers

/*
  Gin 框架常用 API（白話說明）：

  ● c *gin.Context 代表什麼？
    就是「這一筆 HTTP 請求」的資訊包。每次有人打你的 API，Gin 會給你一個 c，
    你要從「請求」拿東西（網址參數、body）或要「回傳」東西給對方，都透過 c 來做。
    可以想成：c = 這一通請求＋你要回覆的介面。

  ● c.Param("id") 的用途？
    URL 裡路徑參數的值。例如網址是 /api/users/123，路由有寫 :id，那 c.Param("id") 就是 "123"。

  ● c.ShouldBindJSON(&input) 的用途？
    把客戶端送來的「JSON  body」自動解析成你指定的 struct（例如 &input），
    並依照 struct 欄位上的 binding tag（如 `binding:"required"`）做基本驗證。
    成功的話 input 裡就有資料；失敗（格式錯、必填沒填）會回傳 err。

  ● c.JSON(狀態碼, 資料) 的用途？
    把「HTTP 狀態碼」和「你要回的資料」一起送給客戶端。Gin 會把資料轉成 JSON 寫進 response。
    例如 c.JSON(201, user) 就是：回傳狀態 201，body 是 user 轉成的 JSON。

  ● gin.H 是什麼？
    寫 JSON 物件用的簡寫。gin.H{"error": "not found"} 就等於回傳 {"error":"not found"}。
*/

import (
	"errors"
	"net/http"
	"strconv"

	"go-api-practice/services"

	"github.com/gin-gonic/gin"
)

// respondError 依 service 錯誤回傳對應 HTTP 狀態與訊息
// 使用 Gin：c.JSON(狀態碼, gin.H{...}) 回傳 JSON 與狀態碼
func respondError(c *gin.Context, err error) {
	if err == nil {
		return
	}
	switch {
	case errors.Is(err, services.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	case errors.Is(err, services.ErrInvalidInput):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// parseID 從 path param 解析 id，無效時回傳 0 與 false，並已寫入 400 回應
// 使用 Gin：c.Param(param) 取得路徑參數，例如路由 /api/items/:id 時 c.Param("id") 即 URL 中的 id
func parseID(c *gin.Context, param string) (int, bool) {
	id, err := strconv.Atoi(c.Param(param))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return 0, false
	}
	return id, true
}
