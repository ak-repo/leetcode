
func findGCD(nums []int) int {

	// sort the array
	sort.Ints(nums)
	smallest := nums[0]
	greatest := nums[len(nums)-1]
	var gcd int

	//
	for greatest >= 0 {
		reminder := greatest % smallest
		if reminder == 0 {
			gcd = smallest
			break
		}
		greatest = smallest
		smallest = reminder

	}

	return gcd

}