package userhandlers

import (
	domain "blog-platform-go/domain/users"
	"blog-platform-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignInHandler struct {
	SignInUseCase domain.ISignInUsecase
}

func (s *SignInHandler) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signin_req domain.SignInRequest

		err := c.ShouldBind(&signin_req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error binding request"})
			return
		}

		user, err := s.SignInUseCase.GetUserByEmail(c, signin_req.Email)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Not found user with the given email"})
			return
		}

		if err := utils.CheckPasswordHash(user.Password, signin_req.Password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user cretidentials"})
			return
		}

		accessToken, err := s.SignInUseCase.CreateAccessToken(&user, "secret access token", 3600)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error create access token"})
			return
		}

		refreshToken, err := s.SignInUseCase.CreateRefreshToken(&user, "secret refresh token", 3600)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error create refresh token"})
			return
		}

		loginResponse := domain.SignInResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		c.JSON(http.StatusOK, loginResponse)
	}
}
