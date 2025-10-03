package route

import (
	"wot-statistics-server/config"
	"wot-statistics-server/wire"

	"github.com/gin-gonic/gin"
)

func NewTankRouter(group *gin.RouterGroup, appConfig *config.AppConfig) {
	tc := wire.InitializeTankController(appConfig)
	group.GET("/tanks", tc.GetTanks)
}
