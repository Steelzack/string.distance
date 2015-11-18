package algorithms

type OptAli struct {
	StringDistance
}

func (dist OptAli) CalculateDistance(fromString string, toString string) int {

	lenStr1 := len(fromString)
	lenStr2 := len(toString)
	d := make([][]int, lenStr1)
	for i := range d {
		d[i] = make([]int, lenStr2)
	}
	i := 0
	j := 0
	cost := 0

	for i < lenStr1 {
		d[i][0] = i
		i++
	}

	for j < lenStr2 {
		d[0][j] = j
		j++
	}

	distance := 0

	for i = 1; i < lenStr1; i++ {
		for j = 1; j < lenStr2; j++ {

			if []byte(fromString)[i] == []byte(toString)[j] {
				cost = 0
			} else {
				cost = 1
			}
			d[i][j] = minimum3(d[i-1][j]+1, // deletion
				d[i][j-1]+1,      // insertion
				d[i-1][j-1]+cost, // substitution
			)
			if i > 1 &&
				j > 1 &&
				[]byte(fromString)[i] == []byte(toString)[j-1] &&
				[]byte(fromString)[i-1] == []byte(toString)[j] {
				d[i][j] = minimum2(
					d[i][j],
					d[i-2][j-2]+cost, // transposition
				)
			}
			if distance < d[i][j] {
				distance = d[i][j]
			}
		}
	}
	return d[lenStr1-1][lenStr2-1]
}
