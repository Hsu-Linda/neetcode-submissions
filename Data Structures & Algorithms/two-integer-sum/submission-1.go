func twoSum(nums []int, target int) []int {
	
	numsMap := make(map[int]int)
	for i, v := range nums {
		numsMap[v] = i
	}

	for i := 0 ; i < len(nums) ; i++ {
		otherHalf := target - nums[i]
		if v, ok := numsMap[otherHalf]; ok && v != i {
			return []int {i, v}
		}
	}
	return []int {}
}
