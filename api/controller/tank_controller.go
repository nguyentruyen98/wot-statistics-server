package controller

import (
	"net/http"
	"strconv"

	"wot-statistics-server/domain"

	"github.com/gin-gonic/gin"
)

type TankController struct {
	TankUseCase domain.TankUseCase
}

func NewTankController(tu domain.TankUseCase) *TankController {
	return &TankController{
		TankUseCase: tu,
	}
}

// GetTanks retrieves all tanks from database
// @Summary Get all tanks
// @Description Get all tanks from database with optional filters
// @Tags tanks
// @Accept json
// @Produce json
// @Param nation query string false "Filter by nation (e.g., ussr, usa, germany)"
// @Param tier query int false "Filter by tier (1-10)"
// @Param type query string false "Filter by type (heavyTank, mediumTank, lightTank, AT-SPG, SPG)"
// @Param premium query bool false "Filter by premium status"
// @Success 200 {object} map[string]interface{} "List of tanks"
// @Failure 400 {object} map[string]interface{} "Invalid request parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/tanks [get]
func (t *TankController) GetTanks(c *gin.Context) {
	// Get filter parameters
	nation := c.Query("nation")
	tierStr := c.Query("tier")
	tankType := c.Query("type")
	premiumStr := c.Query("premium")

	var tier int
	var isPremium *bool

	// Parse tier
	if tierStr != "" {
		parsedTier, err := strconv.Atoi(tierStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid tier parameter",
				"message": "Tier must be a number between 1 and 10",
			})
			return
		}
		tier = parsedTier
	}

	// Parse premium
	if premiumStr != "" {
		parsedPremium, err := strconv.ParseBool(premiumStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid premium parameter",
				"message": "Premium must be true or false",
			})
			return
		}
		isPremium = &parsedPremium
	}

	// Get tanks with filters
	var tanks []domain.Tank
	var err error

	if nation != "" || tier > 0 || tankType != "" || isPremium != nil {
		tanks, err = t.TankUseCase.GetTanksByFilter(nation, tier, tankType, isPremium)
	} else {
		tanks, err = t.TankUseCase.GetTanks()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get tanks",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"count":  len(tanks),
		"data":   tanks,
	})
}

// GetTankByID retrieves a single tank by ID
// @Summary Get tank by ID
// @Description Get a single tank by its tank_id
// @Tags tanks
// @Accept json
// @Produce json
// @Param id path int true "Tank ID"
// @Success 200 {object} map[string]interface{} "Tank details"
// @Failure 400 {object} map[string]interface{} "Invalid tank ID"
// @Failure 404 {object} map[string]interface{} "Tank not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/tanks/{id} [get]
func (t *TankController) GetTankByID(c *gin.Context) {
	idStr := c.Param("id")
	tankID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid tank ID",
			"message": "Tank ID must be a valid number",
		})
		return
	}

	tank, err := t.TankUseCase.GetTankByID(tankID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Tank not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   tank,
	})
}

// ImportTanksFromJSON imports tanks from JSON file
// @Summary Import tanks from JSON
// @Description Import tanks from a JSON file into the database
// @Tags tanks
// @Accept json
// @Produce json
// @Param request body map[string]string true "File path"
// @Success 200 {object} map[string]interface{} "Import successful"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Import failed"
// @Router /api/tanks/import [post]
func (t *TankController) ImportTanksFromJSON(c *gin.Context) {
	var req struct {
		FilePath string `json:"file_path" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "file_path is required",
		})
		return
	}

	count, err := t.TankUseCase.ImportTanksFromJSON(req.FilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to import tanks",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Tanks imported successfully",
		"count":   count,
	})
}

// CreateTank creates a new tank
// @Summary Create a new tank
// @Description Create a new tank in the database
// @Tags tanks
// @Accept json
// @Produce json
// @Param tank body domain.Tank true "Tank data"
// @Success 201 {object} map[string]interface{} "Tank created"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Failed to create tank"
// @Router /api/tanks [post]
func (t *TankController) CreateTank(c *gin.Context) {
	var tank domain.Tank

	if err := c.ShouldBindJSON(&tank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": err.Error(),
		})
		return
	}

	if err := t.TankUseCase.CreateTank(&tank); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create tank",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Tank created successfully",
		"data":    tank,
	})
}

// UpdateTank updates an existing tank
// @Summary Update tank
// @Description Update an existing tank in the database
// @Tags tanks
// @Accept json
// @Produce json
// @Param id path int true "Tank ID"
// @Param tank body domain.Tank true "Tank data"
// @Success 200 {object} map[string]interface{} "Tank updated"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "Tank not found"
// @Failure 500 {object} map[string]interface{} "Failed to update tank"
// @Router /api/tanks/{id} [put]
func (t *TankController) UpdateTank(c *gin.Context) {
	idStr := c.Param("id")
	tankID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid tank ID",
			"message": "Tank ID must be a valid number",
		})
		return
	}

	var tank domain.Tank
	if err := c.ShouldBindJSON(&tank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": err.Error(),
		})
		return
	}

	// Set tank ID from URL parameter
	tank.TankID = tankID

	if err := t.TankUseCase.UpdateTank(&tank); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update tank",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Tank updated successfully",
		"data":    tank,
	})
}

// DeleteTank deletes a tank
// @Summary Delete tank
// @Description Delete a tank from the database
// @Tags tanks
// @Accept json
// @Produce json
// @Param id path int true "Tank ID"
// @Success 200 {object} map[string]interface{} "Tank deleted"
// @Failure 400 {object} map[string]interface{} "Invalid tank ID"
// @Failure 404 {object} map[string]interface{} "Tank not found"
// @Failure 500 {object} map[string]interface{} "Failed to delete tank"
// @Router /api/tanks/{id} [delete]
func (t *TankController) DeleteTank(c *gin.Context) {
	idStr := c.Param("id")
	tankID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid tank ID",
			"message": "Tank ID must be a valid number",
		})
		return
	}

	if err := t.TankUseCase.DeleteTank(tankID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete tank",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Tank deleted successfully",
	})
}

// GetTanksFromAPI fetches tanks from Wargaming API
// @Summary Get tanks from Wargaming API
// @Description Fetch tanks data from Wargaming API
// @Tags tanks
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "API response"
// @Failure 500 {object} map[string]interface{} "Failed to fetch from API"
// @Router /api/tanks/wargaming [get]
func (t *TankController) GetTanksFromAPI(c *gin.Context) {
	apiResponse, err := t.TankUseCase.GetTanksFromAPI()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch tanks from API",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   apiResponse,
	})
}
