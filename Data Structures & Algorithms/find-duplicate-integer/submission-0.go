// [1,2,3,2,2] temp 1
// [1, 2, 3, 2, 2] temp 2

func findDuplicate(nums []int) int {
	cur := nums[0]
	for nums[cur] != 0 {
		temp :=  nums[cur]
		nums[cur] = 0
		cur = temp
	}
	return cur
}
