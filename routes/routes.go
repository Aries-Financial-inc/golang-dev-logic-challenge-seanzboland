package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"aries-technical-challenge/services"
	"aries-technical-challenge/models"
)

// AnalysisResult structure for the response body
type AnalysisResponse struct {
	GraphData       []services.XYValue `json:"graph_data"`
	MaxProfit       float64      `json:"max_profit"`
	MaxLoss         float64      `json:"max_loss"`
	BreakEvenPoints []float64    `json:"break_even_points"`
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/analyze", func(c *gin.Context) {
		var contracts []models.OptionsContract

		if err := c.ShouldBindJSON(&contracts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := AnalysisResponse{
			GraphData: services.CalculateXYValues(contracts),
			MaxProfit: services.CalculateMaxProfit(contracts),
			MaxLoss: services.CalculateMaxLoss(contracts),
			BreakEvenPoints: services.CalculateBreakEvenPoints(contracts),
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	})

	return router
}
