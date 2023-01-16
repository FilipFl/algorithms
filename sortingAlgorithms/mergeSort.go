package sorting_algorithms

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	
	return arr
}


func merge(arr []int, start, mid, stop int) []int {
	l := arr[start:mid]
	r := arr[mid:stop]
	for i, j, k := 0, 0, 0; k<len(arr); k++ {
		if l[i] < r[j] && i < len(l) {
			arr[k] = l[i]
			i++
		} else {
			arr[k] = r[j]
			j++ 
		}
	}
	return arr
}