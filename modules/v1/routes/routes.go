package routes

import (
	"e-signature/app/config"
	//basic "e-signature/pkg/basic_auth"

	//signaturesHandlerV1 "e-signature/modules/v1/utilities/signatures/handler"
	//signaturesViewV1 "e-signature/modules/v1/utilities/signatures/view"
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
	return router
}

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	//signaturesHandlerV1 := signaturesHandlerV1.Handler(db)
	//signaturesViewV1 := signaturesViewV1.View(db)
	userHandlerV1 := userHandlerV1.Handler(db)
	userViewV1 := userViewV1.View(db)
	// Routing Website Service
	user := router.Group("")
	user.GET("/", userViewV1.Index)
	user.GET("/dashboard", userViewV1.Dashboard)
	user.GET("/register", userViewV1.Register)
	user.GET("/login", userViewV1.Login)
	user.POST("/login", userHandlerV1.Login)
	//signatures := router.Group("/", basic.Auth(conf))
	//Routing API Service
	//api := router.Group("/api/v1")
	//api.GET("/dashboard", signaturesHandlerV1.ListDashboard)

	router = ParseTmpl(router)
	return router
}
