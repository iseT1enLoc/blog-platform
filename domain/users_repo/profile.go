package domain

import "context"

type Profile struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type IProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*Profile, error)
}
