// Package fraction contains fraction stuff.
package fraction

// Fraction represents a mathematical fraction.
type Fraction struct {
	numerator   int
	denominator int
}

// NewFraction create a Fraction object from numerator and denominator.
func NewFraction(numerator, denominator int) Fraction {
	nm := 1
	if denominator < 0 {
		nm = -1
	}

	num := numerator * nm
	denom := denominator * nm

	highCommonDenom := highestCommonDenominator(num, denom)

	return Fraction{
		numerator:   num / highCommonDenom,
		denominator: denom / highCommonDenom,
	}
}

// Add adds addend to fraction.
func (f Fraction) Add(addend Fraction) Fraction {
	ma, mb := lowestCommonDenominator(f.denominator, addend.denominator)

	return NewFraction(f.numerator*ma+addend.numerator*mb, f.denominator*ma)
}

func lowestCommonDenominator(da, db int) (int, int) {
	ra := da
	rb := db

	for ra != rb {
		switch {
		case ra < rb:
			ra += da
		case rb < ra:
			rb += db
		}
	}

	return ra / da, rb / db
}

func highestCommonDenominator(da, db int) int {
	d := minInts(da, db)

	for d > 1 && (da%d != 0 || db%d != 0) {
		d--
	}

	return clampInt(d, 1, d)
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
