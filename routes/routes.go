package routes

import (
	"go-api-practice/handlers"

	"github.com/gin-gonic/gin"
)

// Setup 註冊所有 API 路由（路徑保持簡單：/api/資源名、/api/資源名/:id）
//
// Gin 路由說明：
// - r.Group("/api")：建立路徑前綴為 /api 的群組，底下的 GET/POST 等都會加上此前綴
// - api.GET("/items", handler)：註冊 GET /api/items，第二參數為 handler 函式（簽名為 func(*gin.Context)）
// - "/items/:id"：冒號表示路徑參數，:id 可在 handler 內用 c.Param("id") 取得
func Setup(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Items
		api.GET("/items", handlers.GetItems)
		api.GET("/items/:id", handlers.GetItemByID)
		api.POST("/items", handlers.CreateItem)
		api.PUT("/items/:id", handlers.UpdateItem)
		api.DELETE("/items/:id", handlers.DeleteItem)

		// Users
		api.GET("/users", handlers.GetUsers)
		api.GET("/users/:id", handlers.GetUserByID)
		api.POST("/users", handlers.CreateUser)
		api.PUT("/users/:id", handlers.UpdateUser)
		api.DELETE("/users/:id", handlers.DeleteUser)

		// Categories
		api.GET("/categories", handlers.GetCategories)
		api.GET("/categories/:id", handlers.GetCategoryByID)
		api.POST("/categories", handlers.CreateCategory)
		api.PUT("/categories/:id", handlers.UpdateCategory)
		api.DELETE("/categories/:id", handlers.DeleteCategory)
	}
}
