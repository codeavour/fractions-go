package fraction_test

import (
	"testing"

	fraction "github.com/codeavour/fractions-go/src"
	"github.com/stretchr/testify/assert"
)

//nolint: scopelint, funlen
func TestAddingFractions(t *testing.T) {
	testCases := []struct {
		name   string
		augend fraction.Fraction
		addend fraction.Fraction
		sum    fraction.Fraction
	}{
		{
			name:   "zero whole numbers",
			augend: fraction.NewFraction(0, 1),
			addend: fraction.NewFraction(0, 1),
			sum:    fraction.NewFraction(0, 1),
		},
		{
			name:   "non-zero and zero whole numbers",
			augend: fraction.NewFraction(5, 1),
			addend: fraction.NewFraction(0, 1),
			sum:    fraction.NewFraction(5, 1),
		},
		{
			name:   "zero and non-zero whole numbers",
			augend: fraction.NewFraction(0, 1),
			addend: fraction.NewFraction(8, 1),
			sum:    fraction.NewFraction(8, 1),
		},
		{
			name:   "non-zero whole numbers",
			augend: fraction.NewFraction(3, 1),
			addend: fraction.NewFraction(4, 1),
			sum:    fraction.NewFraction(7, 1),
		},
		{
			name:   "negative left whole number addend",
			augend: fraction.NewFraction(-3, 1),
			addend: fraction.NewFraction(1, 1),
			sum:    fraction.NewFraction(-2, 1),
		},
		{
			name:   "negative right whole number addend",
			augend: fraction.NewFraction(8, 1),
			addend: fraction.NewFraction(-5, 1),
			sum:    fraction.NewFraction(3, 1),
		},
		{
			name:   "fractions adding to 1",
			augend: fraction.NewFraction(1, 3),
			addend: fraction.NewFraction(2, 3),
			sum:    fraction.NewFraction(1, 1),
		},
		{
			name:   "different denominators where one is multiple of the other",
			augend: fraction.NewFraction(3, 4),
			addend: fraction.NewFraction(5, 8),
			sum:    fraction.NewFraction(11, 8),
		},
		{
			name:   "different denominators where one isn't multiple of the other",
			augend: fraction.NewFraction(1, 6),
			addend: fraction.NewFraction(4, 9),
			sum:    fraction.NewFraction(11, 18),
		},
		{
			name:   "resulting fraction requiring reduction",
			augend: fraction.NewFraction(3, 4),
			addend: fraction.NewFraction(3, 4),
			sum:    fraction.NewFraction(3, 2),
		},
		{
			name:   "same denominators result requiring reduction",
			augend: fraction.NewFraction(3, 8),
			addend: fraction.NewFraction(1, 8),
			sum:    fraction.NewFraction(1, 2),
		},
		{
			name:   "one fraction with negative numerator",
			augend: fraction.NewFraction(-1, 4),
			addend: fraction.NewFraction(3, 4),
			sum:    fraction.NewFraction(1, 2),
		},
		{
			name:   "one fraction with negative denominator",
			augend: fraction.NewFraction(1, -4),
			addend: fraction.NewFraction(3, 4),
			sum:    fraction.NewFraction(1, 2),
		},
		{
			name:   "one fraction with negative numerator resulting in negative fraction",
			augend: fraction.NewFraction(3, 8),
			addend: fraction.NewFraction(-1, 2),
			sum:    fraction.NewFraction(-1, 8),
		},
		{
			name:   "one fraction with negative denominator resulting in negative fraction",
			augend: fraction.NewFraction(3, 8),
			addend: fraction.NewFraction(1, -2),
			sum:    fraction.NewFraction(-1, 8),
		},
		{
			name:   "fraction with negative numerator and denominator",
			augend: fraction.NewFraction(1, -4),
			addend: fraction.NewFraction(-3, -4),
			sum:    fraction.NewFraction(1, 2),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.sum, testCase.augend.Add(testCase.addend), testCase.name)
		})
	}
}
