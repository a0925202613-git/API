package handlers

// User 相關 API，Gin 用法與 item.go 相同：c.Param、c.ShouldBindJSON、c.JSON 等見 helpers.go 註解。

import (
	"net/http"

	"go-api-practice/models"
	"go-api-practice/services"

	"github.com/gin-gonic/gin"
)

// GetUsers GET /api/users
// c：這一筆請求的資訊包，回傳時用 c.JSON
func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID GET /api/users/:id
func GetUserByID(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	user, err := services.GetUserByID(id)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser POST /api/users
// 參數 c *gin.Context：代表「這一筆請求」的資訊包，拿參數、body、回傳回應都透過 c
func CreateUser(c *gin.Context) {
	var input models.User
	// ShouldBindJSON：把對方送來的 JSON body 自動轉成 input，並依 struct 的 binding 做驗證（如 required）
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.CreateUser(&input)
	if err != nil {
		respondError(c, err)
		return
	}
	// c.JSON(狀態碼, 資料)：把狀態碼和資料一起回給客戶端，Gin 會把資料轉成 JSON。201 表示「成功新增一筆」
	c.JSON(http.StatusCreated, user)
}

// UpdateUser PUT /api/users/:id
func UpdateUser(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	var input models.User
	// ShouldBindJSON：把 request body 的 JSON 轉成 input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.UpdateUser(id, &input)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser DELETE /api/users/:id
func DeleteUser(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	if err := services.DeleteUser(id); err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
