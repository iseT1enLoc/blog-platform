package userhandlers

import (
	domain "blog-platform-go/domain/users"
	"blog-platform-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SignUpHandler struct {
	SignupUseCase domain.ISignUpUseCase
}

func (su *SignUpHandler) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		//declare a signup variable
		var sign_up_req domain.SignUpRequest

		//Binding Json request
		err := c.ShouldBind(&sign_up_req)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.ErrInvalidRequest(err))
			print("line 25\n")
			return
		}
		//Check if the email has already been in db
		_, err = su.SignupUseCase.GetUserByEmail(c, sign_up_req.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Email is already registerd in database"})
			print("line 32\n")
			return
		}
		//encrypt password
		encryptedPassword, _ := utils.HashPassword(sign_up_req.Password)

		//create DB Object for insertion
		user := domain.User{
			Id:       uuid.New(),
			Name:     sign_up_req.Email,
			Email:    sign_up_req.Password,
			Password: encryptedPassword,
		}

		//insertion
		err = su.SignupUseCase.CreateUser(c, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while create user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "create user successfully"})
	}
}
