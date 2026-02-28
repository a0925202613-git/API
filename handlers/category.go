package handlers

// Category 相關 API，Gin 用法與 item.go 相同：c.Param、c.ShouldBindJSON、c.JSON 等見 helpers.go 註解。

import (
	"net/http"

	"go-api-practice/models"
	"go-api-practice/services"

	"github.com/gin-gonic/gin"
)

// GetCategories GET /api/categories
func GetCategories(c *gin.Context) {
	list, err := services.GetCategories()
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, list)
}

// GetCategoryByID GET /api/categories/:id
func GetCategoryByID(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	cat, err := services.GetCategoryByID(id)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, cat)
}

// CreateCategory POST /api/categories
func CreateCategory(c *gin.Context) {
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cat, err := services.CreateCategory(&input)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusCreated, cat)
}

// UpdateCategory PUT /api/categories/:id
func UpdateCategory(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cat, err := services.UpdateCategory(id, &input)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, cat)
}

// DeleteCategory DELETE /api/categories/:id
func DeleteCategory(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	if err := services.DeleteCategory(id); err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
