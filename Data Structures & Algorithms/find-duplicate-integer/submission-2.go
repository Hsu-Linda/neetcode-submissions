// 3 1 3 4 2 3
// 3 3 2 1 4 3


// 0 4 2 3 4
func findDuplicate(nums []int) int {
	var fast, slow int
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0
	for {
		slow2 = nums[slow2]
		slow = nums[slow]
		if slow == slow2 {
			return slow
		}
	}
}


