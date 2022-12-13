package routes

import (
	"e-signature/app/blockhain"
	"e-signature/app/config"
	mid "e-signature/app/middlewares"

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
	router.Static("/landing/css", "./public/assets/landing/css")
	router.Static("/landing/js", "./public/assets/landing/js")
	router.Static("/form/vendor", "./public/assets/form/vendor")
	router.Static("/form/css", "./public/assets/form/css")
	router.Static("/form/js", "./public/assets/form/js")
	router.Static("/form/img", "./public/assets/form/img")
	router.Static("/signatures", "./public/images/signatures")
	router.Static("/file/documents", "./public/temp/pdfsign")

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
	user.GET("/dashboard", mid.Permission(), userViewV1.Dashboard)
	user.GET("/register", userViewV1.Register)
	user.POST("/register", userHandlerV1.Register)
	user.GET("/login", userViewV1.Login)
	user.POST("/login", userHandlerV1.Login)
	user.GET("/logout", mid.Permission(), userHandlerV1.Logout)
	user.GET("/log-user", mid.Permission(), userViewV1.Logg)

	//Routing Signature Service
	signature := router.Group("")
	signature.GET("/my-signatures", mid.Permission(), signaturesViewV1.MySignatures)
	signature.POST("/add-signatures", mid.Permission(), signaturesHandlerV1.AddSignatures)
	signature.GET("/change-signatures/:sign_type", mid.Permission(), signaturesHandlerV1.ChangeSignatures)
	signature.GET("/sign-documents", mid.Permission(), signaturesViewV1.SignDocuments)
	signature.POST("/sign-documents", mid.Permission(), signaturesHandlerV1.SignDocuments)
	signature.GET("/invite-signatures", mid.Permission(), signaturesViewV1.InviteSignatures)
	signature.POST("/invite-signatures", mid.Permission(), signaturesHandlerV1.InviteSignatures)
	signature.GET("/request-signatures", mid.Permission(), signaturesViewV1.RequestSignatures)
	signature.GET("/document/:hash", mid.Permission(), signaturesViewV1.Document)
	signature.POST("/document/:hash", mid.Permission(), signaturesHandlerV1.Document)
	signature.GET("/verification", signaturesViewV1.Verification)
	signature.POST("/verification", signaturesHandlerV1.Verification)
	signature.GET("/download", mid.Permission(), signaturesViewV1.Download)
	signature.GET("/download/:hash", mid.Permission(), signaturesHandlerV1.Download)
	signature.GET("/history", mid.Permission(), signaturesViewV1.History)
	signature.GET("/transactions", signaturesViewV1.Transactions)

	//Testing and Checking Data
	//signature.GET("/verification_result", signaturesViewV1.VerificationResult)
	//signature.GET("/docs/:hash/:id", signaturesHandlerV1.GetDocs)
	//signature.GET("/verif/:hash", signaturesHandlerV1.Verif)

	//Create JWT Token Authentication
	apiV1 := router.Group("/api/v1")
	apiV1.POST("/create-token", userHandlerV1.CreateToken)

	//Analysis Integrity Data
	analysis := apiV1.Group("/analysis")
	analysis.POST("/integrity-document", mid.AuthAPI(db), signaturesHandlerV1.Integrity)
	analysis.GET("/download-document/:hash", mid.AuthAPI(db), signaturesHandlerV1.Download)

	router = ParseTmpl(router)
	return router
}
