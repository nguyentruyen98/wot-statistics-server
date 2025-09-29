package route

import (
	"wot-statistics-server/api/controller"
	"wot-statistics-server/usecase"

	"github.com/gin-gonic/gin"
)

func NewTankRouter(group *gin.RouterGroup) {
	tu := usecase.NewTankUseCase()
	tc := &controller.TankController{TankUseCase: tu}

	group.GET("/tanks", tc.GetTanks)
}
