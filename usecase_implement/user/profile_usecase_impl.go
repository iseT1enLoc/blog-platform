package user

import (
	domain "blog-platform-go/domain/users_repo"
	"context"
	"time"
)

type ProfileUseCase struct {
	user_repo      domain.IUserRepository
	contextTimeOut time.Duration
}

// GetProfileByID implements domain.IProfileUsecase.
func (p *ProfileUseCase) GetProfileByID(c context.Context, userID string) (*domain.Profile, error) {
	panic("unimplemented")
}

func NewProfileUseCase(user_repo domain.IUserRepository, contextTimeOut time.Duration) domain.IProfileUsecase {
	return &ProfileUseCase{
		user_repo:      user_repo,
		contextTimeOut: contextTimeOut,
	}
}
