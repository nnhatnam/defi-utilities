package common

import (
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
)

func BigPow(x, y int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(x), big.NewInt(y), nil)
}

func Exp(base, exponent *big.Int) *big.Int {
	return math.Exp(base, exponent)
}
