package services

import "database/sql"

type MyAppService struct {
	db *sql.DB
}

// NewMyAppService MyAppServiceのコンストラクタ
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
