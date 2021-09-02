package indicators

import (
	"github.com/sdcoffey/big"
	"server/helpers"
)

type volumeIndicator struct {
	*helpers.TimeSeries
}

// NewVolumeIndicator returns an indicator which returns the volume of a candle for a given index
func NewVolumeIndicator(series *helpers.TimeSeries) Indicator {
	return volumeIndicator{series}
}

func (vi volumeIndicator) Calculate(index int) big.Decimal {
	return vi.Candles[index].Volume
}

type closePriceIndicator struct {
	*helpers.TimeSeries
}

// NewClosePriceIndicator returns an Indicator which returns the close price of a candle for a given index
func NewClosePriceIndicator(series *helpers.TimeSeries) Indicator {
	return closePriceIndicator{series}
}

func (cpi closePriceIndicator) Calculate(index int) big.Decimal {
	return cpi.Candles[index].ClosePrice
}

type highPriceIndicator struct {
	*helpers.TimeSeries
}

// NewHighPriceIndicator returns an Indicator which returns the high price of a candle for a given index
func NewHighPriceIndicator(series *helpers.TimeSeries) Indicator {
	return highPriceIndicator{
		series,
	}
}

func (hpi highPriceIndicator) Calculate(index int) big.Decimal {
	return hpi.Candles[index].MaxPrice
}

type lowPriceIndicator struct {
	*helpers.TimeSeries
}

// NewLowPriceIndicator returns an Indicator which returns the low price of a candle for a given index
func NewLowPriceIndicator(series *helpers.TimeSeries) Indicator {
	return lowPriceIndicator{
		series,
	}
}

func (lpi lowPriceIndicator) Calculate(index int) big.Decimal {
	return lpi.Candles[index].MinPrice
}

type openPriceIndicator struct {
	*helpers.TimeSeries
}

// NewOpenPriceIndicator returns an Indicator which returns the open price of a candle for a given index
func NewOpenPriceIndicator(series *helpers.TimeSeries) Indicator {
	return openPriceIndicator{
		series,
	}
}

func (opi openPriceIndicator) Calculate(index int) big.Decimal {
	return opi.Candles[index].OpenPrice
}

type typicalPriceIndicator struct {
	*helpers.TimeSeries
}

// NewTypicalPriceIndicator returns an Indicator which returns the typical price of a candle for a given index.
// The typical price is an average of the high, low, and close prices for a given candle.
func NewTypicalPriceIndicator(series *helpers.TimeSeries) Indicator {
	return typicalPriceIndicator{series}
}

func (tpi typicalPriceIndicator) Calculate(index int) big.Decimal {
	numerator := tpi.Candles[index].MaxPrice.Add(tpi.Candles[index].MinPrice).Add(tpi.Candles[index].ClosePrice)
	return numerator.Div(big.NewFromString("3"))
}