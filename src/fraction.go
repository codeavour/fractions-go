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
	res := Fraction{
		numerator:   f.numerator + addend.numerator,
		denominator: f.denominator,
	}

	return res.reduce()
}

func (f Fraction) reduce() Fraction {
	if f.numerator%f.denominator == 0 {
		return Fraction{
			numerator:   f.numerator / f.denominator,
			denominator: 1,
		}
	}

	return f
}
