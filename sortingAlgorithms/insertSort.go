package sorting_algorithms


func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		j := i
		for (j > 0 && val < arr[j-1]) {
			arr[j] = arr[j-1]
			j = j-1
		}
		arr[j] = val
	}
	return arr
}