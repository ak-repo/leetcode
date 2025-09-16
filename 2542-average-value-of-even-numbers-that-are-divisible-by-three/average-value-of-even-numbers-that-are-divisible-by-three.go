func averageValue(nums []int) int {

	var avg,count int

	for _, v := range nums {

		if v%3 == 0 && v%2 == 0 {
			avg += v
            count++

		}

	}
    if count == 0{
        return 0
    }

	return avg/count

}