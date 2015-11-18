package algorithms

import (
	"bytes"
)

type DamLev struct {
}

func (dist DamLev) CalculateDistance(fromString string, toString string, alphabetSize int) int {
	// "infinite" distance is just the max possible distance
	fullLength := len(fromString) + len(toString)

	// make and initialize the character array indices
	da := make([]byte, alphabetSize)
	for k := 0; k < alphabetSize; k++ {
		da[k] = 0
	}

	// make the distance matrix H[-1..a.length][-1..b.length]
	matrix := make([][]int, len(fromString)+2)
	for k := 0; k < len(fromString)+2; k++ {
		matrix[k] = make([]int, len(toString)+2)
	}

	// initialize the left and top edges of H
	matrix[1][0] = fullLength
	for i := 1; i <= len(fromString)+1; i++ {
		matrix[i][0] = fullLength
		matrix[i][1] = i
	}
	for j := 1; j <= len(toString)+1; j++ {
		matrix[0][j] = fullLength
		matrix[1][j] = j
	}
	// fill in the distance matrix H
	// look at each character in a
	for i := 1; i <= len(fromString); i++ {
		chari := 1
		// look at each character in b
		for j := 1; j <= len(toString); j++ {
			i1 := bytes.IndexByte(da, []byte(toString)[j-1])
			j1 := chari
			var cost int
			if []byte(fromString)[i-1] == []byte(toString)[j-1] {
				cost = 0
				chari = j
			} else {
				cost = 1
			}
			var matrixi1j1 int
			if i1 >= 0 {
				matrixi1j1 = matrix[i1][j1]
			} else {
				matrixi1j1 = 0
			}
			matrix[i+1][j+1] = minimum4( //
				matrix[i][j]+cost, // substitution
				matrix[i+1][j]+1,  // insertion
				matrix[i][j+1]+1,  // deletion
				matrixi1j1+(i-i1)+1+(j-j1),
			)
		}
		da[i] = []byte(fromString)[i-1]
	}
	return matrix[len(fromString)][len(toString)] - 1
}
