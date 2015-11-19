package algorithms

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

	for i := 0; i < lenFromString+1; i++ {
		matrix[i] = make([]int, lenToString+1)
	}

	matrix[0][0] = 0
	for j := 1; j < lenToString +1; j++ {
		matrix[0][j] = matrix[0][j-1] + dist.gap
	}

	for i := 1; i < lenFromString+1; i++ {
		matrix[i][0] = matrix[i-1][0] + dist.gap
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

			result := maximum3(gdiag, gup, gleft)
			matrix[i][j] = result
		}

	}

	logArrayLine(matrix)

	return 0
}
