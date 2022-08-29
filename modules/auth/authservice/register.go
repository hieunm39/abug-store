package authservice

import (
	"abug-store/common"
	"abug-store/modules/auth/authmodel"
	"context"
)


type RegisterStorage interface {
	FindUser(
		ctx context.Context, 
		conditions map[string]interface{}, 
		moreInfo ...string) (*authmodel.User, error)

	CreateUser(ctx context.Context, data *authmodel.UserCreate) error
}

type registerService struct {
	store RegisterStorage
	hasher Hasher
}

func NewRegisterService(registerStorage RegisterStorage, hasher Hasher) *registerService {
	return &registerService{
		store: registerStorage,
		hasher: hasher,
	}
}

func (service *registerService) Register(ctx context.Context, data *authmodel.UserCreate) error {
	user, _ := service.store.FindUser(ctx, map[string]interface{}{
		"username": data.Username,
	})

	if user != nil {
		return authmodel.ErrUsernameExisted
	}

	salt := common.GenSalt(50)
	data.Password = service.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = common.UserRole
	data.Status = 1

	if err := service.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(authmodel.EntityName, err)
	}

	return nil
}

