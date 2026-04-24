// zxyzxyz
// zxy

// azxyxyz

// xxxx
// x -> curL = 1
// xx -> left =0 last =0
// left from 2-> 5   right = 6
// 23456 -> 456

func lengthOfLongestSubstring(s string) int {
	existE := [128]int {}
	for i:=0; i<len(existE); i++ {
		existE[i] = -1
	}


	var maxL, curL, left int
	for right :=0 ; right<len(s);right ++ {
		last := existE[s[right]]
		if last != -1 && left <= last {
			originalLeft := left
			left = last+1
			curL -= (left - originalLeft)
		}
		curL ++
		existE[s[right]] = right

		if curL > maxL {
			maxL = curL
		}
	}

	return maxL
}
