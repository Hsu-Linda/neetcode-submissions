// 0, 1, 2, [3], 4, 5, 6  len=7
// 0, 6
// 1, 5
// 2, 4
// 0, 1, [2], 3, 4, 5, len=6
// 0, 5
// 1, 4
// 2, 3

func reverseString(s []byte) {
	for i:=0; i<= len(s)/2-1 ; i++ {
		shift(s, i, len(s)-1-i)
	}
}

func shift (a []byte, first, second int) {
	a[first], a[second] = a[second], a[first]
}