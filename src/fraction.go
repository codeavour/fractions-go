// Package fraction contains fraction stuff.
package fraction

type Fraction struct {
	numerator   int
	denominator int
}

func NewFraction(numerator, denominator int) Fraction {
	return Fraction{
		numerator:   numerator,
		denominator: denominator,
	}
}

func (f Fraction) Add(addend Fraction) Fraction {
	return Fraction{
		numerator:   0,
		denominator: 1,
	}
}
