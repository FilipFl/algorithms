package sorting_algorithms

import "testing"

type insertSortTestData struct {
	have []int 
}

var insertTests = []insertSortTestData {
	insertSortTestData{have: NewData(11)},
	insertSortTestData{have: NewData(100)},
	insertSortTestData{have: NewData(1000)},
	insertSortTestData{have: NewData(10000)},
	insertSortTestData{have: NewData(100000)},
}


func TestInsertSort(t *testing.T) {
	for _, test := range insertTests {
        sorted := InsertSort(test.have)
		for i := 1; i < len(sorted); i++ {
			if (sorted[i-1] > sorted[i]) {
				t.Errorf("Array is not sorted properly")
				break
			}
		} 
    }
}