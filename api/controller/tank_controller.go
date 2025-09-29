package controller

import (
	"wot-statistics-server/domain"

	"github.com/gin-gonic/gin"
)

type TankController struct {
	TankUseCase domain.TankUseCase
}

func (t *TankController) GetTanks(c *gin.Context) {

	tank, err := t.TankUseCase.GetTanks()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tank)
}
