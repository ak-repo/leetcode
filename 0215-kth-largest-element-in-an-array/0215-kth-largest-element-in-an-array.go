func findKthLargest(nums []int, k int) int {

arr:= QuickSort(nums)

return arr[k-1]
}


func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var less, equal, greater []int

	for _, v := range arr {
		if v < pivot {
			less = append(less, v)
		} else if v > pivot {
			greater = append(greater, v)
		} else {
			equal = append(equal, v)
		}
	}

	return append(append(QuickSort(greater), equal...), QuickSort(less)...)
}