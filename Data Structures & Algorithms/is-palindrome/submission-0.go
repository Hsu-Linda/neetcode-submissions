func isPalindrome(s string) bool {
	l, r :=0, len(s)-1
	for l < r {
		for l<r &&  !isAlphaNum(s[l]) {
			l++
		}
		for l<r && !isAlphaNum(s[r]) {
			r --
		}

		// type diff
		if (isDigit(s[l]) != isDigit(s[r])) {
			return false
		}

		if isDigit(s[l]) && s[l]!=s[r] {
			return false
		}
		
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		
		l++
		r--		
	}

	return true
}

func isAlphaNum(b byte) bool {
	return isDigit(b) || (b >='A' && b <='Z') || (b >='a' && b <='z')
}

func isDigit (i byte) bool{
	return i >= '0' && i <= '9'
}

func toLower (b byte) byte {
	if b >= 'A' && b <='Z' {
		return b+32
	}
	return b
}