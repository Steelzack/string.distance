package algorithms

import (
	"bytes"
	"log"
	"math"
	"strconv"
)

type Direction int

const (
	Diag Direction = 0
	Up   Direction = 1
	Left Direction = 2
	Stop Direction = 3 // Only used for S-W algorithm
)

func minimum4(element1 int, element2 int, element3 int, element4 int) int {
	return int(math.Min(float64(element4), float64(minimum3(element1, element2, element3))))
}

func minimum3(element1 int, element2 int, element3 int) int {
	return int(math.Min(float64(element3), float64(minimum2(element1, element2))))
}

func minimum2(element1 int, element2 int) int {
	return int(math.Min(float64(element1), float64(element2)))
}

func maximumwithdirection3(diagonal int, up int, left int) (value int, direction Direction) {
	value = int(math.Max(float64(left), float64(maximum2(diagonal, up))))
	if value == diagonal {
		direction = Diag
	} else if value == up {
		direction = Up
	} else if value == left {
		direction = Left
	}
	return
}
func maximumwithdirection4(diagonal int, up int, left int, control int) (value int, direction Direction) {
	value = maximum2(int(math.Max(float64(left), float64(maximum2(diagonal, up)))), control)
	if value == diagonal {
		direction = Diag
	} else if value == up {
		direction = Up
	} else if value == left {
		direction = Left
	} else if value == control {
		direction = Stop
	}
	return
}

func maximum2(element1 int, element2 int) int {
	return int(math.Max(float64(element1), float64(element2)))
}

func logArrayLine(array [][]int) {
	lenCol := len(array[0])
	lenRow := len(array)

	for i := 0; i < lenRow; i++ {
		buffer := bytes.NewBufferString("")
		for j := 0; j < lenCol; j++ {
			buffer.WriteString(strconv.Itoa(array[i][j]))
			buffer.WriteString(" ")
		}
		log.Println(buffer.String())
	}
}
