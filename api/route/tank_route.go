package route

import (
	"wot-statistics-server/api/controller"
	"wot-statistics-server/repository"
	"wot-statistics-server/usecase"

	"github.com/gin-gonic/gin"
)

func NewTankRouter(group *gin.RouterGroup) {
	tr := repository.NewTankRepository()
	tu := usecase.NewTankUseCase(tr)
	tc := &controller.TankController{TankUseCase: tu}

	group.GET("/tanks", tc.GetTanks)
}
