func firstUniqChar(s string) int {

	queue := map[rune]int{}
	for _, v := range s {
		queue[v]++
	}

	for i, v := range s {

		if queue[v] == 1 {
			return i
		}
	}

	return -1
}