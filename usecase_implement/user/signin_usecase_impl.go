package user

import (
	domain "blog-platform-go/domain/users_repo"
	"context"
	"time"
)

type SignInUseCase struct {
	user_repo      domain.IUserRepository
	contextTimeOut time.Duration
}

func NewSignInUseCase(user_repo domain.IUserRepository, contextTimeOut time.Duration) domain.ISignInUsecase {
	return &SignInUseCase{
		user_repo:      user_repo,
		contextTimeOut: contextTimeOut,
	}
}

// CreateAccessToken implements domain.ISignInUsecase.
func (s *SignInUseCase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	panic("unimplemented")
}

// CreateRefreshToken implements domain.ISignInUsecase.
func (s *SignInUseCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	panic("unimplemented")
}

// GetUserByEmail implements domain.ISignInUsecase.
func (s *SignInUseCase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	panic("unimplemented")
}
