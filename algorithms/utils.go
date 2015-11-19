package algorithms

import (
	"bytes"
	"log"
	"math"
	"strconv"
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

func maximum3(element1 int, element2 int, element3 int) int {
	return int(math.Max(float64(element3), float64(maximum2(element1, element2))))
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
