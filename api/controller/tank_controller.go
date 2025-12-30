package controller

import (
	"encoding/json"
	"os"
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

func (t *TankController) GetTanks(c *gin.Context) {
	// url := "https://api.worldoftanks.asia/wot/encyclopedia/vehicles/?application_id=d5c27b088716f6a2ca4d043e6fe2ba91&tier=1&language=vi&page_no=1&fields=description"

	// client := &http.Client{}
	// req, err := http.NewRequest(http.MethodGet, url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Bắt buộc thêm các header cơ bản
	// // req.Header.Set("Accept", "application/json")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Go-http-client)") // giả UA giống browser

	// res, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }

	// fmt.Printf("%s\n", body)
	tank, err := t.TankUseCase.GetTanksFromAPI()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tank)
}

// GetUSSRTanks trả về danh sách xe tăng Liên Xô từ file tanks.json
func (t *TankController) GetUSSRTanks(c *gin.Context) {
	// Đọc file tanks.json
	data, err := os.ReadFile("tanks.json")
	if err != nil {
		c.JSON(500, gin.H{"error": "Không thể đọc file tanks.json: " + err.Error()})
		return
	}

	// Parse JSON
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		c.JSON(500, gin.H{"error": "Không thể parse JSON: " + err.Error()})
		return
	}

	// Trả về kết quả
	c.JSON(200, result)
}
