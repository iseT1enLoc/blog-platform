package domain

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name     string    `gorm:"column:name" json:"name"`
	Email    string    `gorm:"column:email;unique" json:"email"`
	Password string    `gorm:"column:password" json:"password"`
}

type IUserRepository interface {
	CreateUser(c context.Context, u *User) (rowAffect int, err error)
	FetchUser(c context.Context) ([]User, error)
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserById(c context.Context, id string) (User, error)
}
