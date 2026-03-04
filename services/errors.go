package services

import "errors"

var (
	// ErrNotFound 資源不存在（handler 會回傳 HTTP 404）
	ErrNotFound = errors.New("not found")
	// ErrInvalidInput 輸入不合法（handler 會回傳 HTTP 400）
	ErrInvalidInput = errors.New("invalid input")
)
