// [1,2,3,4,5,6,7,8]  8
// [8, 1,2,3,4,5,6,7,] // 8
// [7,8, 1,2,3,4,5,6,] // 7

func rotate(nums []int, k int) {
	k = k%len(nums)
	for i:=0; i<k;i++ {
		end := nums[len(nums)-1]
		// nums[1:(len(nums)-2)] = nums[0:(len(nums)-2)]
		for i:=len(nums)-2; i>=0; i-- {
			nums[i+1] = nums[i]
		}
		nums[0] = end
	}
}
