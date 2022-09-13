package routes

import (
	"e-signature/app/blockhain"
	"e-signature/app/config"

	//basic "e-signature/pkg/basic_auth"

	//signaturesHandlerV1 "e-signature/modules/v1/utilities/signatures/handler"

	userHandlerV1 "e-signature/modules/v1/utilities/user/handler"
	userViewV1 "e-signature/modules/v1/utilities/user/view"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParseTmpl(router *gin.Engine) *gin.Engine { //Load HTML Template
	router.Static("/landing/assets", "./public/assets/landing")
	router.Static("/landing/vendor", "./public/assets/landing/vendor")
	router.Static("/landing/swiper", "./public/assets/landing/vendor/swiper")
	router.Static("/landing/purecounter", "./public/assets/landing/vendor/purecounter")
	router.Static("/dash", "./public/assets/dashboard")
	router.Static("/dashcustom", "./public/assets/dashboard/assets/css")

	return router
}

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	blockchain, client := blockhain.Init(conf)
	//signaturesHandlerV1 := signaturesHandlerV1.Handler(db)
	//signaturesViewV1 := signaturesViewV1.View(db)
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
	//user.GET("/add-signature", signaturesViewV1.AddSignature)
	//user.GET("/list-signature", signaturesViewV1.ListSignature)
	//signatures := router.Group("/", basic.Auth(conf))
	//Routing API Service
	//api := router.Group("/api/v1")
	//api.GET("/dashboard", signaturesHandlerV1.ListDashboard)

	router = ParseTmpl(router)
	return router
}
