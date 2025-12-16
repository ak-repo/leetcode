func isAnagram(s string, t string) bool {

	a := map[rune]int{}
	b := map[rune]int{}

	if len(s) != len(t) {
		return false
	}

	for _, v := range s {
		a[v]++
	}

	for _, v := range t {
		b[v]++
	}

	for k, v := range a {
		if b[k] != v {
			return false
		}
	}

	return true
}