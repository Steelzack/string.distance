package algorithms

import (
	"bytes"
	"log"
)

type NeeWun struct {
	gap            int
	missmatchscore int
	exactscore     int
	StringDistance
}

func NewNeeWun(gap int, missmatchscore int, exactscore int) *NeeWun {
	return &NeeWun{gap, missmatchscore, exactscore, nil}
}

func (dist NeeWun) CalculateDistance(fromString string, toString string) int {
	lenFromString := len(fromString)
	lenToString := len(toString)

	matrix := make([][]int, lenFromString+1)
	matrixroute := make([][]int, lenFromString+1)
	for i := 0; i < lenFromString+1; i++ {
		matrix[i] = make([]int, lenToString+1)
		matrixroute[i] = make([]int, lenToString+1)
	}

	matrix[0][0] = 0
	for j := 1; j < lenToString+1; j++ {
		matrix[0][j] = matrix[0][j-1] + dist.gap
		matrixroute[0][j] = int(Left)
	}

	for i := 1; i < lenFromString+1; i++ {
		matrix[i][0] = matrix[i-1][0] + dist.gap
		matrixroute[i][0] = int(Up)

		for j := 1; j < lenToString+1; j++ {
			var score int
			if []byte(toString)[j-1] == []byte(fromString)[i-1] {
				score = dist.exactscore
			} else {
				score = -1 * dist.missmatchscore
			}

			gdiag := matrix[i-1][j-1] + score
			gup := matrix[i-1][j] + dist.gap
			gleft := matrix[i][j-1] + dist.gap

			result, route := maximumwithdirection3(gdiag, gup, gleft)
			matrix[i][j] = result
			matrixroute[i][j] = int(route)
		}

	}

	logArrayLine(matrix)
	logArrayLine(matrixroute)

	modifiedFrom, modifiedTo, walk := dist.traceback(matrixroute, fromString, toString)

	log.Println("Number of steps: ", walk)
	log.Println(modifiedFrom)
	log.Println(modifiedTo)

	return compareSameSizeString(modifiedFrom, modifiedTo)
}
func (dist NeeWun) traceback(matrixroute [][]int, fromString string, toString string) (string, string, int) {

	i := len(matrixroute) - 1
	j := len(matrixroute[i]) - 1
	var bufferFrom bytes.Buffer
	var bufferTo bytes.Buffer
	walk := 0
	for !(i == 0 && j == 0) {
		currentarrow := matrixroute[i][j]
		switch currentarrow {
		case int(Diag):
			bufferFrom.WriteByte([]byte(fromString)[i-1])
			bufferTo.WriteByte([]byte(toString)[j-1])
			i--
			j--
		case int(Left):
			bufferFrom.WriteByte('-')
			bufferTo.WriteByte([]byte(toString)[j-1])
			j--
		case int(Up):
			bufferFrom.WriteByte([]byte(fromString)[i-1])
			bufferTo.WriteByte('-')
			i--
		}
		walk++
	}
	return revertString(bufferFrom.String()), revertString(bufferTo.String()), walk
}
