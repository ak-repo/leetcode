func findGCD(nums []int) int {

	// sort the array
	sort.Ints(nums)
	smallest := nums[0]
	greatest := nums[len(nums)-1]

	for smallest != greatest {
		sub := greatest - smallest
		if smallest == sub {
			return sub
		}
		if sub > smallest {
			greatest = sub
		} else {
			greatest = smallest
			smallest = sub
		}

	}

	return smallest

}