package services

import (
	"database/sql"
	"fmt"
	"go-api-practice/database"
	"go-api-practice/models"
)

/*
  Items：在此實作從 DB 查、寫、改、刪。使用 database.DB（Query / QueryRow / Exec、Scan、RowsAffected）。
*/

// GetItems 取得所有項目
func GetItems() ([]models.Item, error) {
	// 請實作：查 items 表，回傳列表

	// 1. SQL語法
	query := "SELECT id, name, description, created_at, updated_at FROM items"

	// 2. 執行SQL
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("查詢 items 失敗：%w", err)
	}
	//查詢結束時關閉連線
	defer rows.Close()

	// 3. 解析結果
	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, fmt.Errorf("解析 items 失敗：%w", err)
		}
		items = append(items, item)
	}

	// id name description created_at updated_at
	// 1  物件1 描述        2024-01-01 2024-01-02 <-
	// 2  物件2 描述        2024-01-03 2024-01-04 <-

	return items, nil
	//return nil, fmt.Errorf("請實作：查詢 items 並回傳列表")
}

// GetItemByID 依 ID 取得單一項目
func GetItemByID(id int) (*models.Item, error) {
	// 請實作：依 id 查單筆，找不到回傳 ErrNotFound
	// 1
	var ErrNotFound = fmt.Errorf("record not found")

	query := "SELECT id, name, description, created_at, updated_at FROM items WHERE id = ?"

	var item models.Item

	// id name  description created_at updated_at
	// 1  物件1 描述  2024-01-01 2024-01-02
	err := database.DB.QueryRow(query, id).Scan(
		&item.ID,
		&item.Name,
		&item.Description,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	// 假設今天有錯誤
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("查詢單筆item失敗:%w", err)
	}

	return &item, nil
}

// CreateItem 新增項目
func CreateItem(input *models.Item) error {
	// 請實作：INSERT 進 items，回傳新增的資料
	// Body: {"name": "物件1", "description": "描述"}
	// JSON -> struct -> SQL

	query := "INSERT INTO Items (ID, Name, Description) VALUES (?,?,?)"

	if _, err := database.DB.Exec(query, input.ID, input.Name, input.Description); err != nil {
		return fmt.Errorf("新增 item 失敗：%w", err)
	}

	return nil
}

// UpdateItem 更新項目
func UpdateItem(id int, input *models.Item) error {
	// 請實作：UPDATE 該 id，回傳更新後資料；不存在回傳 ErrNotFound
	query := "UPDATE Items SET ID = ? WHERE = "
	return fmt.Errorf("請實作：UPDATE item 並回傳（id=%d）", id)
}

// DeleteItem 刪除項目
func DeleteItem(id int) error {
	// 請實作：DELETE 該 id，不存在回傳 ErrNotFound
	query := "DELETE FROM Items WHERE = id"

	return fmt.Errorf("請實作：DELETE item（id=%d）", id)
}
