package common

import "math/big"

type Fraction struct {
	rat *big.Rat
}

func NewFractionFromInt64(numerator, denominator int64) *Fraction {
	return &Fraction{
		rat: big.NewRat(numerator, denominator),
	}
}

func NewFraction(numerator, denominator *big.Int) *Fraction {
	if denominator == nil {
		denominator = big.NewInt(1)
	}

	if denominator.Cmp(big.NewInt(0)) == 0 {
		panic("denominator is zero")
	}

	return &Fraction{
		rat: new(big.Rat).SetFrac(numerator, denominator),
	}
}

func (f *Fraction) Quotient() *big.Int {
	return new(big.Int).Div(f.rat.Num(), f.rat.Denom())
}

// Remainder returns the remainder after floor division
func (f *Fraction) Remainder() *Fraction {

	remainder := new(big.Int)
	remainder.Rem(f.rat.Num(), f.rat.Denom())
	denom := new(big.Int).Set(f.rat.Denom())
	return &Fraction{
		rat: new(big.Rat).SetFrac(remainder, denom),
	}

}

func (f *Fraction) Invert() *Fraction {

	return &Fraction{
		rat: new(big.Rat).Inv(f.rat),
	}

}

func (f *Fraction) Add(other *Fraction) *Fraction {
	//if other == nil {
	//	panic("other fraction is nil")
	//}

	//will be panic if denominator is nil, no need to check unless we want to have a clearer panic message.
	return &Fraction{
		rat: new(big.Rat).Add(f.rat, other.rat),
	}

}

func (f *Fraction) Subtract(other *Fraction) *Fraction {

	return &Fraction{
		rat: new(big.Rat).Sub(f.rat, other.rat),
	}

}

func (f *Fraction) LessThan(other *Fraction) bool {

	return f.rat.Cmp(other.rat) == -1

}

func (f *Fraction) EqualTo(other *Fraction) bool {

	return f.rat.Cmp(other.rat) == 0

}

func (f *Fraction) GreaterThan(other *Fraction) bool {
	return f.rat.Cmp(other.rat) == 1

}

func (f *Fraction) Multiply(other *Fraction) *Fraction {

	return &Fraction{
		rat: new(big.Rat).Mul(f.rat, other.rat),
	}

}

func (f *Fraction) Divide(other *Fraction) *Fraction {

	return &Fraction{
		rat: new(big.Rat).Quo(f.rat, other.rat),
	}

}

func (f *Fraction) ToFloat() *big.Float {
	return new(big.Float).SetRat(f.rat)
}

func (f *Fraction) ToFloat64() (fl float64, exact bool) {
	return f.rat.Float64()
}
