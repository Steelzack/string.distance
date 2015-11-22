// +build !steelzackstringdistance
package algorithms

import (
	"github.com/steelzack/string.distance/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateDistanceWithAlphabet(t *testing.T) {
	damleu1 := new(algorithms.DamLev)
	result1 := damleu1.CalculateDistanceWithAlphabet("SimpleTest", "SimplTest", 20)
	assert.Equal(t, 1, result1, "Calculation for DamLeu is not correct!")

	damleu2 := new(algorithms.DamLev)
	result2 := damleu2.CalculateDistanceWithAlphabet("ATCGT", "TGGTG", 20)
	assert.Equal(t, 3, result2, "Calculation for DamLeu is not correct!")
}
