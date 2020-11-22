package fraction_test

import (
	"testing"

	fraction "github.com/codeavour/fractions-go/src"

	"github.com/stretchr/testify/assert"
)

func TestAddingFractions(t *testing.T) {
	testCases := []struct {
		name        string
		leftAddend  fraction.Fraction
		rightAddend fraction.Fraction
		sum         fraction.Fraction
	}{
		{
			name:        "zero whole numbers",
			leftAddend:  fraction.NewFraction(0, 1),
			rightAddend: fraction.NewFraction(0, 1),
			sum:         fraction.NewFraction(0, 1),
		},
		{
			name:        "non-zero and zero whole numbers",
			leftAddend:  fraction.NewFraction(5, 1),
			rightAddend: fraction.NewFraction(0, 1),
			sum:         fraction.NewFraction(5, 1),
		},
		{
			name:        "zero and non-zero whole numbers",
			leftAddend:  fraction.NewFraction(0, 1),
			rightAddend: fraction.NewFraction(8, 1),
			sum:         fraction.NewFraction(8, 1),
		},
		{
			name:        "non-zero whole numbers",
			leftAddend:  fraction.NewFraction(3, 1),
			rightAddend: fraction.NewFraction(4, 1),
			sum:         fraction.NewFraction(7, 1),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.sum, testCase.leftAddend.Add(testCase.rightAddend), testCase.name)
		})
	}
}
