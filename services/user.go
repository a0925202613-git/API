package services

import (
	"fmt"

	"go-api-practice/models"
)

/*
  Users：在此實作從 DB 查、寫、改、刪。使用 database.DB（Query / QueryRow / Exec、Scan、RowsAffected）。
*/

// GetUsers 取得所有使用者
func GetUsers() ([]models.User, error) {
	// 請實作：查 users 表，回傳列表
	return nil, fmt.Errorf("請實作：查詢 users 並回傳列表")
}

// GetUserByID 依 ID 取得單一使用者
func GetUserByID(id int) (*models.User, error) {
	// 請實作：依 id 查單筆，找不到回傳 ErrNotFound
	return nil, fmt.Errorf("請實作：依 id 查詢單筆 user（id=%d）", id)
}

// CreateUser 新增使用者
func CreateUser(input *models.User) (*models.User, error) {
	// 請實作：INSERT 進 users，回傳新增的資料
	return nil, fmt.Errorf("請實作：INSERT users 並回傳（name=%s）", input.Name)
}

// UpdateUser 更新使用者
func UpdateUser(id int, input *models.User) (*models.User, error) {
	// 請實作：UPDATE 該 id，回傳更新後資料；不存在回傳 ErrNotFound
	return nil, fmt.Errorf("請實作：UPDATE user 並回傳（id=%d）", id)
}

// DeleteUser 刪除使用者
func DeleteUser(id int) error {
	// 請實作：DELETE 該 id，不存在回傳 ErrNotFound
	return fmt.Errorf("請實作：DELETE user（id=%d）", id)
}
