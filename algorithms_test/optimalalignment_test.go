// +build !steelzackstringdistance
package algorithms

import (
	"github.com/steelzack/string.distance/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateDistanceOptAli(t *testing.T) {
	optaln1 := new(algorithms.OptAli)
	result1 := optaln1.CalculateDistance("SimpleTest", "SimplTest")
	assert.Equal(t, 1, result1, "Calculation for Optimal Alignment is not correct!")

	optaln2 := new(algorithms.OptAli)
	result2 := optaln2.CalculateDistance("ATCGT", "TGGTG")
	assert.Equal(t, 3, result2, "Calculation for Optimal Alignment is not correct!")

	optaln3 := new(algorithms.OptAli)
	result3 := optaln3.CalculateDistance("GACTTAC", "CGTGAATTCAT")
	assert.Equal(t, 6, result3, "Calculation for Optimal Alignment is not correct!")
}
