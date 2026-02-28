package handlers

import (
	"net/http"

	"go-api-practice/models"
	"go-api-practice/services"

	"github.com/gin-gonic/gin"
)

// GetItems GET /api/items
// c *gin.Context：代表「這一筆請求」的資訊包，拿參數、body、回傳回應都透過 c
func GetItems(c *gin.Context) {
	items, err := services.GetItems()
	if err != nil {
		respondError(c, err)
		return
	}
	// c.JSON(狀態碼, 資料)：把狀態碼和資料回給客戶端，Gin 會把資料轉成 JSON
	c.JSON(http.StatusOK, items)
}

// GetItemByID GET /api/items/:id
// 路徑中的 :id 由 parseID(c, "id") 透過 c.Param("id") 取得
func GetItemByID(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	item, err := services.GetItemByID(id)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, item)
}

// CreateItem POST /api/items
func CreateItem(c *gin.Context) {
	var input models.Item
	// ShouldBindJSON：把對方送來的 JSON body 自動轉成 input，並依 binding tag 做驗證
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := services.CreateItem(&input)
	if err != nil {
		respondError(c, err)
		return
	}
	// c.JSON(201, item)：回傳狀態 201（成功新增）與 item 轉成的 JSON
	c.JSON(http.StatusCreated, item)
}

// UpdateItem PUT /api/items/:id
func UpdateItem(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := services.UpdateItem(id, &input)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, item)
}

// DeleteItem DELETE /api/items/:id
func DeleteItem(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	if err := services.DeleteItem(id); err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
