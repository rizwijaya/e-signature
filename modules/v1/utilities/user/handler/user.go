package user

import (
	"e-signature/modules/v1/utilities/user/repository"
	"e-signature/modules/v1/utilities/user/service"

	api "e-signature/app/contracts"
	ss "e-signature/modules/v1/utilities/user/service"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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

func Handler(db *mongo.Database, blockhain *api.Api, client *ethclient.Client) *userHandler {
	userRepository := repository.NewRepository(db, blockhain, client)
	userService := service.NewService(userRepository)
	userHandler := NewUserHandler(userService)
	return userHandler
}
