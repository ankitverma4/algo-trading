package indicators

import "github.com/sdcoffey/big"

// TODO: Add Logic For Exponential Moving Average
 /*


  */

type emaIndicator struct {
	indicator   Indicator
	window      int
	alpha       big.Decimal
	resultCache resultCache
}


func NewEMAIndicator(indicator Indicator, window int) Indicator {
	return &emaIndicator{
		indicator:   indicator,
		window:      window,
		alpha:       big.ONE.Frac(2).Div(big.NewFromInt(window + 1)),
		resultCache: make([]*big.Decimal, 1000),
	}
}

func (ema *emaIndicator) Calculate(index int) big.Decimal {
	if cachedValue := returnIfCached(ema, index, func(i int) big.Decimal {
		return NewSimpleMovingAverage(ema.indicator, ema.window).Calculate(i)
	}); cachedValue != nil {
		return *cachedValue
	}

	todayVal := ema.indicator.Calculate(index).Mul(ema.alpha)
	result := todayVal.Add(ema.Calculate(index - 1).Mul(big.ONE.Sub(ema.alpha)))

	cacheResult(ema, index, result)

	return result
}

func (ema emaIndicator) cache() resultCache { return ema.resultCache }

func (ema *emaIndicator) setCache(newCache resultCache) {
	ema.resultCache = newCache
}

func (ema emaIndicator) windowSize() int { return ema.window }

