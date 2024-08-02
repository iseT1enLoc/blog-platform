package user

import (
	domain "blog-platform-go/domain/users"
	"context"
	"time"
)

type SignUpUseCase struct {
	user_repo      domain.IUserRepository
	contextTimeOut time.Duration
}

func NewSignUpUseCase(user_repo domain.IUserRepository, contextTimeOut time.Duration) domain.ISignUpUseCase {
	return &SignUpUseCase{
		user_repo:      user_repo,
		contextTimeOut: contextTimeOut,
	}
}

// CreateAccessToken implements domain.ISignUpUseCase.
func (s *SignUpUseCase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	panic("unimplemented")
}

// CreateRefreshToken implements domain.ISignUpUseCase.
func (s *SignUpUseCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	panic("unimplemented")
}

// CreateUser implements domain.ISignUpUseCase.
func (s *SignUpUseCase) CreateUser(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeOut)
	defer cancel()
	_, err := s.user_repo.CreateUser(ctx, user)
	return err
}

// GetUserByEmail implements domain.ISignUpUseCase.
func (s *SignUpUseCase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeOut)
	defer cancel()
	user, err := s.user_repo.GetUserByEmail(ctx, email)
	return user, err
}
