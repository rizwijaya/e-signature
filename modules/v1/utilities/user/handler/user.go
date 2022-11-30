package user

import (
	api "e-signature/app/contracts"
	er "e-signature/modules/v1/utilities/signatures/repository"
	es "e-signature/modules/v1/utilities/signatures/service"
	"e-signature/modules/v1/utilities/user/repository"
	"e-signature/modules/v1/utilities/user/service"
	ss "e-signature/modules/v1/utilities/user/service"
	pw "e-signature/pkg/crypto"
	docs "e-signature/pkg/document"
	images "e-signature/pkg/images"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler interface {
	Login(c *gin.Context)
}

type userHandler struct {
	userService      ss.Service
	signatureService es.Service
}

func NewUserHandler(userService ss.Service, signatureService es.Service) *userHandler {
	return &userHandler{userService, signatureService}
}

func Handler(db *mongo.Database, blockhain *api.Api, client *ethclient.Client) *userHandler {
	userRepository := repository.NewRepository(db, blockhain, client)
	crypto := pw.NewCrypto()
	userService := service.NewService(userRepository, crypto)

	signatureRepository := er.NewRepository(db, blockhain, client)
	documents := docs.NewDocuments()
	images := images.NewImages()
	signatureService := es.NewService(signatureRepository, images, documents)
	userHandler := NewUserHandler(userService, signatureService)
	return userHandler
}
