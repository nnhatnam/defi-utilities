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

// DivideBigInt performs division of two big.Int numbers and returns the result as a float64.
// If an error occurs during the conversion, it returns 0 and the error.
func DivideBigInt(dividend *big.Int, divisor *big.Int) float64 {
	dividendFloat := new(big.Float).SetInt(dividend)
	divisorFloat := new(big.Float).SetInt(divisor)
	quotient := new(big.Float).Quo(dividendFloat, divisorFloat)
	quotientFloat, _ := quotient.Float64()
	return quotientFloat
}
