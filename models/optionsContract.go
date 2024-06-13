package models

import (
	"math"
	"time"
)

// OptionType defines the type of the option
type OptionType string

const (
	Call OptionType = "call"
	Put  OptionType = "put"
)

// LongShortType defines if the option is long or short
type LongShortType string

const (
	Long  LongShortType = "long"
	Short LongShortType = "short"
)

// OptionsContract represents an options contract
type OptionsContract struct {
	Type          OptionType	`json:"type"`
	StrikePrice   float64		`json:"strike_price"`
	Bid           float64		`json:"bid"`
	Ask           float64		`json:"ask"`
	ExpirationDate time.Time	`json:"expiration_date"`
	LongShort     LongShortType	`json:"long_short"`
}

// CalculateProfitLoss calculates the profit or loss at a given underlying price
func (o OptionsContract) CalculateProfitLoss(underlyingPrice float64) float64 {
	var intrinsicValue float64
	if o.Type == Call {
		intrinsicValue = math.Max(0, underlyingPrice-o.StrikePrice)
	} else {
		intrinsicValue = math.Max(0, o.StrikePrice-underlyingPrice)
	}

	if o.LongShort == Long {
		return intrinsicValue - o.Ask
	} else {
		return o.Bid - intrinsicValue
	}
}
