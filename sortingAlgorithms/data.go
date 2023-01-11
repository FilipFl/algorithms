package sorting_algorithms

import (
	"math/rand"
	"time"
)

func NewData(length int) []int {
	rand.Seed(time.Now().UnixNano())
	result := make([]int, length)
	elementsRange := length * 10
	for i := 0; i < length; i++ {
		result[i] = rand.Intn(elementsRange)
	}
	return result
}
