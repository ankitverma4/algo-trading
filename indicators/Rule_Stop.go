package indicators

import (
	"github.com/sdcoffey/big"
	"server/helpers"
)

type stopLossRule struct {
	Indicator
	tolerance big.Decimal
}

// NewStopLossRule returns a new rule that is satisfied when the given loss tolerance (a percentage) is met or exceeded.
// Loss tolerance should be a value between -1 and 1.
func NewStopLossRule(series *helpers.TimeSeries, lossTolerance float64) Rule {
	return stopLossRule{
		Indicator: NewClosePriceIndicator(series),
		tolerance: big.NewDecimal(lossTolerance),
	}
}

func (slr stopLossRule) IsSatisfied(index int, record *helpers.TradingRecord) bool {
	if !record.CurrentPosition().IsOpen() {
		return false
	}

	openPrice := record.CurrentPosition().CostBasis()
	loss := slr.Indicator.Calculate(index).Div(openPrice).Sub(big.ONE)
	return loss.LTE(slr.tolerance)
}

