package data

import (
	domain "blog-platform-go/domain/users"
	"context"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.IUserRepository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser implements domain.IUserRepository.
func (s *UserRepository) CreateUser(c context.Context, u *domain.User) (rowAffect int, err error) {
	/*user := domain.SignUpRequest{
		UserName: u.Name,
		Email:    u.Email,
		Password: u.Password,
	}*/
	result := s.db.Create(&u)
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

// FetchUser implements domain.IUserRepository.
func (u *UserRepository) FetchUser(c context.Context) ([]domain.User, error) {
	var listUser []domain.User

	if err := u.db.Find(&listUser).Error; err != nil {
		log.Fatalf(err.Error())
	}
	return listUser, nil
}

// GetUserByEmail implements domain.IUserRepository.
func (u *UserRepository) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	err := u.db.Where("email=?", email).First(&user)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			log.Println("User not found")

		} else {
			log.Fatal("Error finding user:", err.Error)
		}
	}
	return user, nil
}

// GetUserById implements domain.IUserRepository.
func (u *UserRepository) GetUserById(c context.Context, id string) (domain.User, error) {
	var user domain.User
	if err := u.db.Where("id=?", id).First(&user).Error; err != nil {
		log.Fatalf("There is something wrong while get user by id %s", err.Error())
	}
	return user, nil
}
