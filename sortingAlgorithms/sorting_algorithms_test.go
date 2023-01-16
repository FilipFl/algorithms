package sorting_algorithms

import "testing"

type sortingTestData struct {
	have []int 
}

var sortingTests = []sortingTestData {
	sortingTestData{have: NewData(11)},
	sortingTestData{have: NewData(100)},
	sortingTestData{have: NewData(1000)},
	sortingTestData{have: NewData(10000)},
	sortingTestData{have: NewData(100000)},
}


func TestInsertSort(t *testing.T) {
	for _, test := range sortingTests {
		toTest := make([]int, len(test.have))
		copy(toTest, test.have)
        sorted := InsertSort(toTest)
		for i := 1; i < len(sorted); i++ {
			if (sorted[i-1] > sorted[i]) {
				t.Errorf("Array is not sorted properly")
				break
			}
		} 
    }
}


func TestMergeSort(t *testing.T) {
	for _, test := range sortingTests {
		toTest := make([]int, len(test.have))
		copy(toTest, test.have)
        sorted := MergeSort(toTest)
		for i := 1; i < len(sorted); i++ {
			if (sorted[i-1] > sorted[i]) {
				t.Errorf("Array is not sorted properly")
				break
			}
		} 
    }
}