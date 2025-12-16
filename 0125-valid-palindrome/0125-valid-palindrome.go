func isPalindrome(s string) bool {

    left,right:= 0, len(s)-1

    for left < right{

        l:= rune(s[left])
        r:= rune(s[right])

        
		if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			left++
			continue
		}
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			right--
			continue
		}

		if unicode.ToLower(l) != unicode.ToLower(r) {
			return false
		}

		left++
		right--
    }
    return true
    
}