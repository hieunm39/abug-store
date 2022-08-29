package authstorage

import (
	"abug-store/common"
	"abug-store/modules/auth/authmodel"
	"context"

	"gorm.io/gorm"
)


func (s *SQLStorage) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*authmodel.User, error) {

	db := s.db

	db.Table(authmodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user authmodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}