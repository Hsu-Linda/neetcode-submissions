// [1,2,3,4,_,6,7,8]  //5
// [5,2,3,4,1,6,7,8]  //5


func rotate(nums []int, k int) {
	for i:=0; i<k;i++ {
		end := nums[len(nums)-1]
		// nums[1:(len(nums)-2)] = nums[0:(len(nums)-2)]
		for i:=len(nums)-2; i>=0; i-- {
			nums[i+1] = nums[i]
		}
		nums[0] = end
	}
}
