package algorithms

type StringDistance interface {
	CalculateDistance(fromString string, toString string) int
}
