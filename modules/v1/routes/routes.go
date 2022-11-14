package routes

import (
	"e-signature/app/blockhain"
	"e-signature/app/config"

	//basic "e-signature/pkg/basic_auth"

	signaturesHandlerV1 "e-signature/modules/v1/utilities/signatures/handler"
	signaturesViewV1 "e-signature/modules/v1/utilities/signatures/view"
	userHandlerV1 "e-signature/modules/v1/utilities/user/handler"
	userViewV1 "e-signature/modules/v1/utilities/user/view"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ParseTmpl(router *gin.Engine) *gin.Engine { //Load HTML Template
	router.Static("/landing/assets", "./public/assets/landing")
	router.Static("/landing/vendor", "./public/assets/landing/vendor")
	router.Static("/landing/swiper", "./public/assets/landing/vendor/swiper")
	router.Static("/landing/purecounter", "./public/assets/landing/vendor/purecounter")
	router.Static("/landing/img", "./public/assets/landing/img")
	router.Static("/landing/js", "./public/assets/landing/js")
	router.Static("/form/vendor", "./public/assets/form/vendor")
	router.Static("/form/css", "./public/assets/form/css")
	router.Static("/form/js", "./public/assets/form/js")
	router.Static("/form/img", "./public/assets/form/img")
	router.Static("/signatures", "./public/images/signatures")
	router.Static("/file/documents", "./public/temp/pdfsign")

	// router.Static("/dash", "./public/assets/dashboard")
	// router.Static("/dashcustom", "./public/assets/dashboard/assets/css")

	return router
}

func Init(db *mongo.Database, conf config.Conf, router *gin.Engine) *gin.Engine {
	blockchain, client := blockhain.Init(conf)
	signaturesHandlerV1 := signaturesHandlerV1.Handler(db, blockchain, client)
	signaturesViewV1 := signaturesViewV1.View(db, blockchain, client)
	userHandlerV1 := userHandlerV1.Handler(db, blockchain, client)
	userViewV1 := userViewV1.View(db, blockchain, client)
	// Routing Website Service
	user := router.Group("")
	user.GET("/", userViewV1.Index)
	user.GET("/dashboard", userViewV1.Dashboard)
	user.GET("/register", userViewV1.Register)
	user.POST("/register", userHandlerV1.Register)
	user.GET("/login", userViewV1.Login)
	user.POST("/login", userHandlerV1.Login)
	user.GET("/logout", userHandlerV1.Logout)

	signature := router.Group("")
	signature.GET("/my-signatures", signaturesViewV1.MySignatures)
	signature.POST("/add-signatures", signaturesHandlerV1.AddSignatures)
	signature.GET("/change-signatures/:sign_type", signaturesHandlerV1.ChangeSignatures)
	signature.GET("/sign-documents", signaturesViewV1.SignDocuments)
	signature.POST("/sign-documents", signaturesHandlerV1.SignDocuments)
	signature.GET("/invite-signatures", signaturesViewV1.InviteSignatures)
	signature.POST("/invite-signatures", signaturesHandlerV1.InviteSignatures)
	signature.GET("/request-signatures", signaturesViewV1.RequestSignatures)
	signature.GET("/document/:hash", signaturesViewV1.Document)

	//user.GET("/list-signature", signaturesViewV1.ListSignature)
	//signatures := router.Group("/", basic.Auth(conf))
	//Routing API Service
	//api := router.Group("/api/v1")
	//api.GET("/dashboard", signaturesHandlerV1.ListDashboard)

	router = ParseTmpl(router)
	return router
}
