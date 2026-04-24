// A A B A B B A
// 0 1 2 3 4 5 6
// left:   2
// right:  6
// maxFreq 3
// maxL    4

func characterReplacement(s string, k int) int {
	left := -1
	count := [26]int{}
	maxFreq := 0
	maxL := 0
	
	for right:=0; right<len(s); right++ {
		// right +1
		count[s[right]-'A'] ++ 
		
		// count max freq
		if count[s[right]-'A'] > maxFreq {
			maxFreq = count[s[right]-'A']
		}

		for right-left-maxFreq > k {
			left ++
			count[s[left]-'A'] --
		}

		curL := right - left
		if maxL < curL {
			maxL = curL
		}
		
	}

	return maxL
}
