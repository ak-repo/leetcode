func removeOuterParentheses(s string) string {

	res := []string{}
	c := 0
	for _, v := range s {
		if v == '(' {
			if c > 0 {
				res = append(res, string(v))
			}
			c += 1
		} else {
			c -= 1
			if c > 0 {
				res = append(res, string(v))
			}
		}
	}

	return strings.Join(res, "")

}