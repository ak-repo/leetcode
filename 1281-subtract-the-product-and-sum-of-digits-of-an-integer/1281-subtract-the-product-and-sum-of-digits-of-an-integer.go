func subtractProductAndSum(n int) int {
	sum:= 0
	product := 1

	strN := strconv.Itoa(n)
	for _, v := range strN {
		sum += int(v - '0')
		product *= int(v - '0')
	}

	return product - sum
}