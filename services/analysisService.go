package services

import (
	"aries-technical-challenge/models"
)

// XYValue represents a pair of X and Y values
type XYValue struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// calculateXYValues calculates the X & Y values for the risk & reward graph for multiple contracts
func CalculateXYValues(contracts []models.OptionsContract) []XYValue {
	points := []XYValue{}

	minPrice, maxPrice := getPriceRange(contracts)
	for price := minPrice; price <= maxPrice; price += 1.0 {
		totalProfitLoss := 0.0
		for _, option := range contracts {
			totalProfitLoss += option.CalculateProfitLoss(price)
		}
		points = append(points, XYValue{
			X: price,
			Y: totalProfitLoss,
		})
	}

	return points
}

// calculateMaxProfit calculates the maximum possible profit for multiple contracts
func CalculateMaxProfit(contracts []models.OptionsContract) float64 {
	totalProfit := 0.0
	unlimitedProfit := false
	for _, option := range contracts {
		if option.LongShort == models.Long {
			if option.Type == models.Call {
				unlimitedProfit = true
			} else {
				totalProfit += option.StrikePrice - option.Ask
			}
		} else {
			totalProfit += option.Bid
		}
	}
	if unlimitedProfit {
		// Returns -1 meaning unlimited profit
		return -1
	}
	return totalProfit
}

// calculateMaxLoss calculates the maximum possible loss for multiple contracts
func CalculateMaxLoss(contracts []models.OptionsContract) float64 {
	totalLoss := 0.0
	unlimitedLoss := false
	for _, option := range contracts {
		if option.LongShort == models.Long {
			totalLoss += option.Ask
		} else {
			if option.Type == models.Call {
				unlimitedLoss = true
			} else {
				totalLoss += option.StrikePrice - option.Bid
			}
		}
	}
	if unlimitedLoss {
		// Returns -1 meaning unlimited loss
		return -1
	}
	return totalLoss
}

// calculateBreakEvenPoints calculates all breakeven points for multiple options
func CalculateBreakEvenPoints(contracts []models.OptionsContract) []float64 {
	// Create a breakeven points as a map architecture to prevent duplicated insertion
	breakevenPoints := make(map[float64]bool)
	for _, option := range contracts {
		if option.Type == models.Call {
			if option.LongShort == models.Long {
				breakevenPoints[option.StrikePrice+option.Ask] = true
			} else {
				breakevenPoints[option.StrikePrice+option.Bid] = true
			}
		} else {
			if option.LongShort == models.Long {
				breakevenPoints[option.StrikePrice-option.Ask] = true
			} else {
				breakevenPoints[option.StrikePrice-option.Bid] = true
			}
		}
	}
	keys := make([]float64, 0, len(breakevenPoints))
	for k := range breakevenPoints {
		keys = append(keys, k)
	}
	return keys
}

func getPriceRange(options []models.OptionsContract) (minPrice float64, maxPrice float64) {
	// Initialize minPrice & maxPrice with the first option contract's strike_price
	minPrice = options[0].StrikePrice
	maxPrice = options[0].StrikePrice
	for _, option := range options {
		if option.StrikePrice < minPrice {
			minPrice = option.StrikePrice
		}
		if option.StrikePrice > maxPrice {
			maxPrice = option.StrikePrice
		}
	}
	// Extend range by a reasonable margin
	return minPrice - 20, maxPrice + 20
}
