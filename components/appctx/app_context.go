package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetDatabaseConnection() *gorm.DB
	GetSecretKey() string
}

type appCtx struct {
	db *gorm.DB
	secretKey string
}

func NewAppContext(db *gorm.DB, secretKey string) *appCtx {
	return &appCtx{db : db}
}

func (ctx *appCtx) GetDatabaseConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}