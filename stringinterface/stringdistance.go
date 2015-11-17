package stringinterface

import ()

type StringDistance interface {
	calculateDistance(fromString string, toString string) int
}
