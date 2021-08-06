package indicators

import "github.com/sdcoffey/big"

type PPSIndicator struct {
	indicator Indicator
	ppsType string
	high  float64
	low  float64
	close float64
	index int


}

func NewPPSIndicator(indicator Indicator,ppsType string, high float64, low float64,close float64,index int) Indicator {
	return PPSIndicator{indicator,ppsType, high,low,close,index}
}




func (pps PPSIndicator) Calculate(index int) big.Decimal {


	pp := (pps.high + pps.low + pps.close)/3
	return (big.NewDecimal(pp))

}

func(pps PPSIndicator) getSupportValue() []float64{

	pp := (pps.high + pps.low + pps.close)/3
	r1 := 2*pp - pps.low
	s1 := 2 * pp - pps.high
	r2 := pp + (pps.high - pps.low)
	s2 := pp - (pps.high- pps.low)
	r3 := pp + 2 * (pps.high - pps.low)
	s3 := pp - 2 * (pps.high - pps.low)

	result := []float64{r1,s1,r2,s2,r3,s3}

	return result
}