package factorial

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnitFactorial(t *testing.T) {
	assert := assert.New(t)
	tests :=
		[]struct {
			n      int
			result int
		}{
			{0, 1},
			{1, 1},
			{2, 2},
			{3, 6},
			{4, 24},
			{5, 120},
			{6, 720},
			{7, 5040},
			{8, 40320},
			{9, 362880},
			{10, 3628800},
			{100, 93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000},
		}

	for _, test := range tests {
		/* act */
		v := Factorial(test.n)
		/* assert */
		assert.Equal(test.result, v)
	}
}
