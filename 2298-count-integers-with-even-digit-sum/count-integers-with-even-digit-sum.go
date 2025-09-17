func countEven(num int) int {

	var count int

	for i := 2; i <= num; i++ {
		strI := strconv.Itoa(i)
		var total int
		for _, j := range strI {
			total += int(j)

		}
		if total%2 == 0 {
			fmt.Println("str,", strI)
			count++
		}

	}

	return count

}