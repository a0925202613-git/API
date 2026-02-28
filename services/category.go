package services

import (
	"fmt"

	"go-api-practice/models"
)

/*
  Categories：在此實作從 DB 查、寫、改、刪。使用 database.DB（Query / QueryRow / Exec、Scan、RowsAffected）。
*/

// GetCategories 取得所有分類
func GetCategories() ([]models.Category, error) {
	// 請實作：查 categories 表，回傳列表
	return nil, fmt.Errorf("請實作：查詢 categories 並回傳列表")
}

// GetCategoryByID 依 ID 取得單一分類
func GetCategoryByID(id int) (*models.Category, error) {
	// 請實作：依 id 查單筆，找不到回傳 ErrNotFound
	return nil, fmt.Errorf("請實作：依 id 查詢單筆 category（id=%d）", id)
}

// CreateCategory 新增分類
func CreateCategory(input *models.Category) (*models.Category, error) {
	// 請實作：INSERT 進 categories，回傳新增的資料
	return nil, fmt.Errorf("請實作：INSERT categories 並回傳（name=%s）", input.Name)
}

// UpdateCategory 更新分類
func UpdateCategory(id int, input *models.Category) (*models.Category, error) {
	// 請實作：UPDATE 該 id，回傳更新後資料；不存在回傳 ErrNotFound
	return nil, fmt.Errorf("請實作：UPDATE category 並回傳（id=%d）", id)
}

// DeleteCategory 刪除分類
func DeleteCategory(id int) error {
	// 請實作：DELETE 該 id，不存在回傳 ErrNotFound
	return fmt.Errorf("請實作：DELETE category（id=%d）", id)
}
