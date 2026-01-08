





func isHappy(n int) bool {
	str := strconv.Itoa(n)
	runes := []rune(str)
    coll:= make(map[int]bool)

	for {

		ng := len(runes)
		var res int
		for i := 0; i < ng; i++ {
			val, _ := strconv.Atoi(string(runes[i]))
			val = val * val
			res += val
		}
		if res != 1 {
			str := strconv.Itoa(res)
			runes = []rune(str)
		}
        if coll[res]{
            return false
        }
		if res <= 0 {
			return false
		}
		if res > 2147483648 {
			return false
		}
		if res == 1 {
			return true
		}
        coll[res] = true
	}

}