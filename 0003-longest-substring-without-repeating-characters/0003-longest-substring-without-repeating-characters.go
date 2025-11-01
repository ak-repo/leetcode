func lengthOfLongestSubstring(s string) int {

	var maxValue, minValue int
	coll := map[string]int{}

	for i := 0; i < len(s); i++ {
		str := string(s[i])

		index, ok := coll[str]
		if ok {

			clear(coll)
			if maxValue < minValue {
				maxValue = minValue
			}
			minValue = 0
			i = index
			continue
		}
		coll[str] = i
		minValue++
	}
	if maxValue < minValue {
		maxValue = minValue
	}

	return maxValue

}