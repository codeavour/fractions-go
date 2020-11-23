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
	ma, mb := lowestCommonDenominator(f.denominator, addend.denominator)

	res := Fraction{
		numerator:   f.numerator*ma + addend.numerator*mb,
		denominator: f.denominator * ma,
	}

	return res.reduce()
}

func lowestCommonDenominator(da, db int) (int, int) {
	ma := 1
	mb := 1

	for {
		ra := da * ma
		rb := db * mb

		if ra < rb {
			ma++
		} else if rb < ra {
			mb++
		} else {
			break
		}
	}

	return ma, mb
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
