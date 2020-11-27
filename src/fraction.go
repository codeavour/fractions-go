// Package fraction contains code to add mathematical fractions.
package fraction

// Fraction represents a mathematical fraction.
type Fraction struct {
	numerator   int
	denominator int
}

// NewFraction create a Fraction object from numerator and denominator.
func NewFraction(numerator, denominator int) Fraction {
	numeratorSign := signOf(numerator)
	denominatorSign := signOf(denominator)

	positiveDenominator := denominator * denominatorSign
	highestCommonDivisor := highestCommonDivisor(numerator*numeratorSign, positiveDenominator)

	return Fraction{
		numerator:   numerator * denominatorSign / highestCommonDivisor,
		denominator: positiveDenominator / highestCommonDivisor,
	}
}

// Add adds addend to fraction.
func (f Fraction) Add(addend Fraction) Fraction {
	return NewFraction(f.numerator*addend.denominator+addend.numerator*f.denominator, f.denominator*addend.denominator)
}

func highestCommonDivisor(da, db int) int {
	for db != 0 {
		t := db
		db = da % db
		da = t
	}

	return da
}

func signOf(x int) int {
	if x < 0 {
		return -1
	}

	return 1
}
