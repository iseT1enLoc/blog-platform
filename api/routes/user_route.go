package routes

import (
	userhandlers "blog-platform-go/api/handlers/user_handlers"
	data "blog-platform-go/repository/user"
	"blog-platform-go/usecase_impl/user"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewSignupRouter(timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := data.NewUserRepository(db)

	sc := userhandlers.SignUpHandler{SignupUseCase: user.NewSignUpUseCase(ur, timeout)}
	group.POST("/signup", sc.SignUp())
}
