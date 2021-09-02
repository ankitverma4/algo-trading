package strategies

import (
	"server/helpers"
	"server/indicators"
	"time"

	"github.com/sdcoffey/big"
)

func Strategy1() {

	series := helpers.NewTimeSeries()

	threshold := 10
	threshold2 := 20

	// fetch this from your preferred exchange
	dataset := helpers.ReadCsvFile("/Users/abhiagr2/Documents/AlgoTrading/algo-trading/data/BAJAJFINSV.csv")

	for _, datum := range dataset {

		start, _ := time.Parse(datum[0], time.StampMilli)
		period := helpers.NewTimePeriod(time.Time(start), time.Hour*24)

		candle := helpers.NewCandle(period)
		candle.OpenPrice = big.NewFromString(datum[1])
		candle.ClosePrice = big.NewFromString(datum[4])
		candle.MaxPrice = big.NewFromString(datum[2])
		candle.MinPrice = big.NewFromString(datum[3])

		series.AddCandle(candle)
	}

	closePrices := indicators.NewClosePriceIndicator(series)
	s1 := indicators.NewEMAIndicator(closePrices, 4) // Create an exponential moving average with a window of 10
	s2 := indicators.NewEMAIndicator(closePrices, 9)
	s3 := indicators.NewEMAIndicator(closePrices, 18)

	closeIndicator := indicators.NewClosePriceIndicator(series) // returns closing price of the series

	// record trades on this object
	record := helpers.NewTradingRecord()

	long := ((s1.Calculate(4).Sub(s2.Calculate(9))).GT(big.NewDecimal(
		float64(threshold)))) && (s2.Calculate(9).Sub(s3.Calculate(10))).GT(big.NewDecimal(float64(threshold2)))

	short := ((s3.Calculate(4).Sub(s2.Calculate(9))).GT(big.NewDecimal(
		float64(threshold2)))) && (s2.Calculate(9).Sub(s3.Calculate(10))).GT(big.NewDecimal(float64(threshold)))

	// Is satisfied when the price ema moves above 30 and the current position is new
	entryRule := And(
		indicators.NewCrossUpIndicatorRule(closeIndicator, s3),
		long)

	// Is satisfied when the price ema moves below 10 and the current position is open
	exitRule := And(
		indicators.NewCrossDownIndicatorRule(closeIndicator, s3),
		short)

	strategy := RuleStrategy{
		UnstablePeriod: 10, // Period before which ShouldEnter and ShouldExit will always return false
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}

	strategy.ShouldEnter(0, record) // returns false

}
