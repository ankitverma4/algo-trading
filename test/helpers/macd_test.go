package helpers

import (
	"server/indicators"
	"server/test/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMACDIndicator(t *testing.T) {
	series := utils.RandomTimeSeries(100)

	macd := indicators.NewMACDIndicator(indicators.NewClosePriceIndicator(series), 12, 26)

	assert.NotNil(t, macd)
}

func TestNewMACDHistogramIndicator(t *testing.T) {
	series := utils.RandomTimeSeries(100)

	macd := indicators.NewMACDIndicator(indicators.NewClosePriceIndicator(series), 12, 26)
	macdHistogram := indicators.NewMACDHistogramIndicator(macd, 9)

	assert.NotNil(t, macdHistogram)
}
