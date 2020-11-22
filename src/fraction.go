// Package fraction contains fraction stuff.
package fraction

// Fraction represents a mathematical fraction.
type Fraction struct {
	numerator   int
	denominator int
}

// NewFraction create a Fraction object from numerator and denominator.
func NewFraction(numerator, denominator int) Fraction {
	return Fraction{
		numerator:   numerator,
		denominator: denominator,
	}
}

// Add adds addend to fraction.
func (f Fraction) Add(addend Fraction) Fraction {
	return Fraction{
		numerator:   f.numerator + addend.numerator,
		denominator: 1,
	}
}
