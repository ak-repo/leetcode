
func countEven(num int) int {

	var count int

	for i := 2; i <= num; i++ {

		if digitSum(i)%2 == 0 {
			count++
		}
	}

	return count

}

func digitSum(n int) (sum int) {

	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum

}
