package sort

import (
	"math/rand"
)

func NewData(length int) []int {
	result := make([]int, length)
	elementsRange := length * 10
	for i := 0; i < length; i++ {
		result[i] = rand.Intn(elementsRange)
	}
	return result
}
