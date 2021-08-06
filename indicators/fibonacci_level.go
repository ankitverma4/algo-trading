package indicators

import "github.com/sdcoffey/big"

type fibonacciIndicator struct {
	indicator Indicator
	minLevel  float64
	maxLevel  float64
	index int

}

func NewFibonacciLevelIndicator(indicator Indicator, minLevel float64, maxLevel float64,index int) Indicator {
	return fibonacciIndicator{indicator, minLevel,maxLevel,index}
}

// Return the value of fibonaccilevel at the given index position
func (fib fibonacciIndicator) Calculate(index int) big.Decimal {

	ratios := [8]float64{0,0.236, 0.382, 0.5 , 0.618, 0.786,1}
	result := make([]float64, 8)

sum := 0.0
for _,val := range ratios{

		sum = fib.maxLevel - ((fib.maxLevel - fib.minLevel)*val)
		result = append(result,sum)
	}

return (big.NewDecimal(result[index]))
}






