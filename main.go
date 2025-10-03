package main

import (
	"wot-statistics-server/api/route"
	"wot-statistics-server/config"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	appConfig := config.GetAppConfig()

	route.SetupRouter(router, appConfig)

	router.Run()

}
