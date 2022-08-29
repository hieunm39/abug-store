package authstorage

import (
	"abug-store/common"
	"abug-store/modules/auth/authmodel"
	"context"
)


func (s *SQLStorage) CreateUser(ctx context.Context, data *authmodel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(authmodel.User{}.TableName()).Create(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)

	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}