// Package fraction contains fraction stuff.
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
	d := minInts(da, db)

	for d > 1 && (da%d != 0 || db%d != 0) {
		d--
	}

	return clampInt(d, 1, d)
}

func signOf(x int) int {
	if x < 0 {
		return -1
	}

	return 1
}

func minInts(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func clampInt(x, min, max int) int {
	if x < min {
		return min
	} else if x > max {
		return max
	}

	return x
}
