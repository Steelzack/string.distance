package algorithms

import (
	"math"
)

func minimum4(element1 int, element2 int, element3 int, element4 int) int {
	return int(math.Min(float64(element4), float64(minimum3(element1,element2,element3))))
}

func minimum3(element1 int, element2 int, element3 int) int {
	return int(math.Min(float64(element3), float64(minimum2(element1,element2))))
}

func minimum2(element1 int, element2 int) int {
	return int(math.Min(float64(element1), float64(element2)))
}