// a [b] c befe b c a
// a b c efe c b [c] a


// l =0 r =1
func validPalindrome(s string) bool {
	if len(s)<= 2 {
		return true
	}
	
	isDel := false
	
	l, r :=0, len(s)-1
	for l < r {
		if s[l] == s[r] {
			l ++
			r --
			continue
		}
		
		if isDel {
			return false
		}

		if s[l+1] == s[r] {
			if l+2 > r-1 ||
				(l+2 <= r-1 && s[l+2] == s[r-1]) {
					isDel = true
					l += 2
					r --
					continue
				}
		}
		
		if s[l] == s[r-1] {
			if l+1 > r-2 ||
				(l+1 <= r-2 && s[l+1] == s[r-2]) {
				isDel = true
				l ++
				r -= 2
				continue
			}
			
			
		}
		
		return false
	}

	return true
}
