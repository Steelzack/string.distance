//go:build !jesperancinhastringdistance
// +build !jesperancinhastringdistance

package algorithms

import (
	"github.com/jesperancinha/string.distance/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateDistanceSmiWat(t *testing.T) {
	smiwat := algorithms.NewSmiWat(-4, 3, 5)
	result := smiwat.CalculateDistance("GACTTAC", "CGTGAATTCAT")
	assert.Equal(t, 6, result, "Calculation for Smith Waterman is not correct!")
}
