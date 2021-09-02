package indicators

import "github.com/sdcoffey/big"

type plusDiIndicator struct {
	avgPlusDMIndicator   emaIndicator
	window      int
	alpha       big.Decimal
	resultCache resultCache
}



