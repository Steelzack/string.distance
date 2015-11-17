package optimalalinment

import (
	"github.com/steelzack/string.distance/stringinterface"
)

type OptAli struct {
	stringinterface.StringDistance
}

func (dist OptAli) minimum3(element1 int, element2 int, element3 int) int {
	min := element1
	if element2 < element1 {
		min = element2
	} else if element3 < element1 {
		min = element3
	}
	return min
}

func (dist OptAli) minimum2(element1 int, element2 int) int {
	min := element1
	if element2 < element1 {
		min = element2
	}
	
	return min
}

func (dist OptAli) calculateDistance(fromString string, toString string) int {

	// d is a table with lenStr1+1 rows and lenStr2+1 columns
	lenStr1 := len(fromString)
	lenStr2 := len(toString)
	d := make([]int, lenStr1, lenStr2)
	// i and j are used to iterate over str1 and str2
	i := 0
	j := 0
	cost := 0

	// for loop is inclusive, need table 1 row/column larger than string length
	for i < lenStr1 {
		d[:i][0] = i
		i++
	}

	for j < lenStr2 {
		d[:0][j] = j
		j++
	}

	// pseudo-code assumes string indices start at 1, not 0
	// if implemented, make sure to start comparing at 1st letter of strings
	for i < lenStr1 {
		for j < lenStr2 {

			if []byte(fromString)[i] == []byte(toString)[j] {
				cost = 0
			} else {
				cost = 1
			}
			d[:i][j] = dist.minimum3(d[:i-1][j]+1, // deletion
				d[:i][j-1]+1,      // insertion
				d[:i-1][j-1]+cost, // substitution
			)
			if i > 1 &&
				j > 1 &&
				[]byte(fromString)[i] == []byte(toString)[j-1] &&
				[]byte(fromString)[i-1] == []byte(toString)[j] {
				d[:i][j] = dist.minimum2(
					d[:i][j],
					d[:i-2][j-2]+cost, // transposition
				)
			}
		}

	}
	return d[:lenStr1][lenStr2]
}
