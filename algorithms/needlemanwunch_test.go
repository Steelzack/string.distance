//go:build !jesperancinhastringdistance
// +build !jesperancinhastringdistance

package algorithms

import (
	"github.com/jesperancinha/string.distance/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateDistanceNeeWun(t *testing.T) {
	neewun := algorithms.NewNeeWun(-2, 1, 1)
	result := neewun.CalculateDistance("ATCGT", "TGGTG")
	assert.Equal(t, 3, result, "Calculation for Needlman Wunsch is not correct!")
}
