package authservice

import (
	"abug-store/common"
	"abug-store/components/tokenprovider"
	"abug-store/modules/auth/authmodel"
	"context"
)

// 1. Find user, email
// 2. Hash password from input and compare with password in db
// 3. Provider: issue JWT Token for client
// 4. Access Token and Refresh Token
// 5. Return tokens.

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*authmodel.User, error) 
}

type Hasher interface {
	Hash(data string) string
}

type loginService struct {
	store         LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginService(tokenProvider tokenprovider.Provider, storeUser LoginStorage, hasher Hasher, expiry int) *loginService {
	return &loginService{
		tokenProvider: tokenProvider,
		store: storeUser,
		hasher: hasher,
		expiry: expiry,
	}
}

func (service *loginService) Login(ctx context.Context, data *authmodel.UserLogin) (*authmodel.Account, error) {
	user, err := service.store.FindUser(ctx, map[string]interface{}{"username": data.Username })

	if err != nil {
		return nil, authmodel.ErrUsernameOrPasswordInvalid
	}

	passwordHashed := service.hasher.Hash(data.Password + user.Salt)
	
	if passwordHashed != user.Password {
		return nil, authmodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := service.tokenProvider.Generate(payload, service.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := service.tokenProvider.Generate(payload, service.expiry*2)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := authmodel.NewAccount(accessToken, refreshToken)

	return account, nil
}