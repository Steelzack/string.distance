package algorithms

import ()

type SmiWat struct {
	gap            int
	missmatchscore int
	exactscore     int
	StringDistanceImpl
}

func NewSmiWat(gap int, missmatchscore int, exactscore int) *SmiWat {
	return &SmiWat{gap, missmatchscore, exactscore, *new(StringDistanceImpl)}
}

func (dist SmiWat) CalculateDistance(fromString string, toString string) int {
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
		matrix[0][j] = 0
		matrixroute[0][j] = int(Left)
	}

	for i := 1; i < lenFromString+1; i++ {
		matrix[i][0] = 0
		matrixroute[i][0] = int(Up)

		for j := 1; j < lenToString+1; j++ {
			var score int
			if []rune(toString)[j-1] == []rune(fromString)[i-1] {
				score = dist.exactscore
			} else {
				score = -1 * dist.missmatchscore
			}

			gdiag := matrix[i-1][j-1] + score
			gup := matrix[i-1][j] + dist.gap
			gleft := matrix[i][j-1] + dist.gap

			result, route := maximumwithdirection4(gdiag, gup, gleft, 0)
			matrix[i][j] = result
			matrixroute[i][j] = int(route)
		}

	}

	logArrayLine(matrix)
	logArrayLine(matrixroute)

	positions := dist.getAllMatrixPositions(matrix)

	var modifiedFrom string
	var modifiedTo string
	var distance int

	var minDistance int

	for i := 0; i < len(positions); i++ {
		bufferFrom, bufferTo, walk := dist.traceback(matrixroute, fromString, toString, positions[i][0], positions[i][1])
		modifiedFrom, modifiedTo, _ = dist.revertedStringsAndWalk(bufferFrom, bufferTo, walk)
		distance = compareSameSizeString(modifiedFrom, modifiedTo)
		if minDistance == 0 || distance < minDistance {
			minDistance = distance
		}
	}

	return minDistance
}

func (dist SmiWat) getAllMatrixPositions(matrix [][]int) [][]int {
	a := make([][]int, 0)
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			foundValue := matrix[i][j]
			if foundValue >= max {
				a = appendPosition(a, i, j, foundValue)
				max = foundValue
			}
		}
	}

	cleanArray := make([][]int, 0)
	for i := 0; i < len(a); i++ {
		if a[i][2] == max {
			cleanArray = appendPosition(cleanArray, a[i][0], a[i][1], a[i][2])
		}
	}

	return cleanArray
}
