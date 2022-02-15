//go:build !jesperancinhastringdistance
// +build !jesperancinhastringdistance

package algorithms

import (
	"github.com/jesperancinha/string.distance/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateDistanceDamLevWithAlphabet(t *testing.T) {
	damleu1 := new(algorithms.DamLev)
	result1 := damleu1.CalculateDistanceWithAlphabet("SimpleTest", "SimplTest", 20)
	assert.Equal(t, 1, result1, "Calculation for DamLeu is not correct!")

	damleu2 := new(algorithms.DamLev)
	result2 := damleu2.CalculateDistanceWithAlphabet("ATCGT", "TGGTG", 20)
	assert.Equal(t, 3, result2, "Calculation for DamLeu is not correct!")
}
