package main

import (
	"e-signature/app/config"
	database "e-signature/app/databases"
	routesV1 "e-signature/modules/v1/routes"
	"e-signature/pkg/html"
	error "e-signature/pkg/http-error"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func setup() (*mongo.Database, config.Conf, *gin.Engine) {
	conf, err := config.Init()
	gin.SetMode(conf.App.Mode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starter Database Monggo")
	db := database.Init(conf)

	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(conf.App.Secret_key))
	router.Use(sessions.Sessions("smartsign", cookieStore))
	router.HTMLRender = html.Render("./public/templates")
	//Error Handling for 404 Not Found Page and Method Not Allowed
	router.NoRoute(error.PageNotFound())
	router.NoMethod(error.NoMethod())
	return db, conf, router
}

func main() {
	conf, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	router := routesV1.Init(setup()) //Version 1

	//fmt.Println("Starter " + conf.App.Name)
	router.Run(":" + conf.App.Port)
}
