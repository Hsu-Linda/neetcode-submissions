// [1,2,3,4,5,6,7,_]  8
// [8, 1,2,3,4,5,6,7,] // 8
// [7,8, 1,2,3,4,5,6,] // 7

func rotate(nums []int, k int) {
	for i:=0; i<k;i++ {
		end := nums[len(nums)-1]
		for i:=len(nums)-2; i>=0; i-- {
			nums[i+1] = nums[i]
		}
		nums[0] = end
	}
}
