func isValid(s string) bool {
	store := make([]rune, 0 , len(s))
	for _, e := range s {
		if isOpen(e) {
			store = append(store, e)
			continue
		}
		ok, last := getLast(store)
		if !ok {
			// current is close, but before that don't have any open -> no corresponding
			return false
		}
		if isSame(last, e) {
			store = popupLast(store)
		} else {
			return false
		}
	}

	return 0 == len(store)
}

func isOpen(s rune) bool {
	return '('==s || '{'==s || '['==s
}

func getLast(s []rune) (ok bool, value rune) {
	if len(s) == 0 {
		return false, rune(0)
	}
	return true, s[len(s)-1]
}

func isSame(open, close rune) bool {
	switch 
	{
		case open == '(':
			return close == ')'
		case open == '[':
			return close == ']'
		case open == '{':
			return close == '}'
		default:
			return false
	}
}

// [0, 1,2] // len = 3-1
func popupLast(s []rune) []rune {
	return s[: len(s)-1]
}
