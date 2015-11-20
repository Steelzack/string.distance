package algorithms

import (

)

type SmiWat struct {
	StringDistance
}

func (dist SmiWat) CalculateDistance(fromString string, toString string) int {
	return 1
}