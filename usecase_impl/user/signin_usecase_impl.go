package user

import (
	domain "blog-platform-go/domain/users"
	"blog-platform-go/utils"
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
	return utils.CreateAccessToken(user, secret, expiry)
}

// CreateRefreshToken implements domain.ISignInUsecase.
func (s *SignInUseCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}

// GetUserByEmail implements domain.ISignInUsecase.
func (s *SignInUseCase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeOut)
	defer cancel()
	return s.user_repo.GetUserByEmail(ctx, email)
}
