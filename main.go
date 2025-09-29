package main

import (
	"wot-statistics-server/api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	route.SetupRouter(router)

	router.Run()

}
