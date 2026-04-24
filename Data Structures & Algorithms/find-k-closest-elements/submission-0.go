//  k = 2 -> return two integer
//  x = 6 -> interger close to 6

// |5-6| < |8-6|
// choose 5

// |4-6| == |8-6|
// since 4(a) < (8)b -> choose 4

// [2,4,5]

func findClosestElements(arr []int, k int, x int) []int {
	if k >len(arr) {
		return arr
	}

	output := make([]int, 0, k)
	for i:=0; i<len(arr); i++ {
		if i < k {
			output = append(output, arr[i])
			continue
		}
		if arr[i] > x &&
			arr[i]-x >= x-output[0] {
			return output
		}
		output = output[1:]
		output = append(output, arr[i])
	}
	return output

}