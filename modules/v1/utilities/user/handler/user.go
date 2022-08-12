package user

import (
	"e-signature/modules/v1/utilities/user/repository"
	"e-signature/modules/v1/utilities/user/service"

	api "e-signature/app/contracts"
	ss "e-signature/modules/v1/utilities/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler interface {
	Login(c *gin.Context)
}

type userHandler struct {
	userService ss.Service
}

func NewUserHandler(userService ss.Service) *userHandler {
	return &userHandler{userService}
}

func Handler(db *gorm.DB, blockhain *api.Api) *userHandler {
	userRepository := repository.NewRepository(db)
	userService := service.NewService(userRepository, blockhain)
	userHandler := NewUserHandler(userService)
	return userHandler
}
