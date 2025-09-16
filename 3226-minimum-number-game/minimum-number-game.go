func numberGame(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	var arr []int
	sort.Ints(nums)
	for i := 0; i < n; i++ {

		alice := nums[i]
		i++
		bob := nums[i]
		// nums = nums[i+1:]
		arr = append(arr, bob, alice)

	}
	return arr
}