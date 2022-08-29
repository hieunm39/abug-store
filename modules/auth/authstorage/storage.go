package authstorage

import "gorm.io/gorm"

type SQLStorage struct {
	db *gorm.DB
}

func NewSQLStorage(db *gorm.DB) *SQLStorage {
	return &SQLStorage{db: db}
}