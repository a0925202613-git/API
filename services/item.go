package services

import (
	"fmt"

	"go-api-practice/models"
)

/*
  Items：在此實作從 DB 查、寫、改、刪。使用 database.DB（Query / QueryRow / Exec、Scan、RowsAffected）。
*/

// GetItems 取得所有項目
func GetItems() ([]models.Item, error) {
	// 請實作：查 items 表，回傳列表
	return nil, fmt.Errorf("請實作：查詢 items 並回傳列表")
}

// GetItemByID 依 ID 取得單一項目
func GetItemByID(id int) (*models.Item, error) {
	// 請實作：依 id 查單筆，找不到回傳 ErrNotFound
	return nil, fmt.Errorf("請實作：依 id 查詢單筆 item（id=%d）", id)
}

// CreateItem 新增項目
func CreateItem(input *models.Item) (*models.Item, error) {
	// 請實作：INSERT 進 items，回傳新增的資料
	return nil, fmt.Errorf("請實作：INSERT items 並回傳（name=%s）", input.Name)
}

// UpdateItem 更新項目
func UpdateItem(id int, input *models.Item) (*models.Item, error) {
	// 請實作：UPDATE 該 id，回傳更新後資料；不存在回傳 ErrNotFound
	return nil, fmt.Errorf("請實作：UPDATE item 並回傳（id=%d）", id)
}

// DeleteItem 刪除項目
func DeleteItem(id int) error {
	// 請實作：DELETE 該 id，不存在回傳 ErrNotFound
	return fmt.Errorf("請實作：DELETE item（id=%d）", id)
}
