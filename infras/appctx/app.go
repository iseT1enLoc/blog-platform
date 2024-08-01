package infras

import "gorm.io/gorm"

type AppContext interface {
	GetConnectionToDB() *gorm.DB
	GetSecretKeyString() string
}

type appctx struct {
	db        *gorm.DB
	SecretKey string
}

func NewAppContext(db *gorm.DB, secretKey string) *appctx {
	return &appctx{
		db:        db,
		SecretKey: secretKey,
	}
}
func (ctx *appctx) GetConnectionToDB() *gorm.DB {
	return ctx.db
}

func (ctx *appctx) GetSecretKeyString() string {
	return ctx.SecretKey
}
