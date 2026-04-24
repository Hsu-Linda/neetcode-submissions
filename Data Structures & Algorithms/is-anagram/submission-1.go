func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	usingCha := make(map[rune]int)
	for _, c := range s {
		usingCha[c] += 1
	}

	for _, c := range t {
		if v, ok := usingCha[c]; !ok{
			return false
		} else if v < 1 {
			return false
		}
		usingCha[c] -= 1
	}
	return true
}
