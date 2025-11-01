func lengthOfLongestSubstring(s string) int {

	var maxValue, minValue int
	coll := map[byte]int{}

	for i := 0; i < len(s); i++ {

		key := s[i]
		index, ok := coll[key]
		if ok {
			clear(coll)
			if maxValue < minValue {
				maxValue = minValue
			}
			minValue = 0
			i = index
			continue
		}
		coll[key] = i
		minValue++
	}
	if maxValue < minValue {
		maxValue = minValue
	}

	return maxValue

}
