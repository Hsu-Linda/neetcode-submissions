func threeSum(nums []int) [][]int {
	// value can't be the same, output in any order
	// a + b + c = 0
	// a and b are fixed, then we can predict c is fixed
	// a is fixed, maybe b and c can find other options, b and c both needs to change and also b can not change to the c
	sort.Slice(nums, func(i, j int) bool{
		return nums[i] < nums[j]
	})
	fmt.Printf("The sorted is %+v\n", nums)
    // i, i+1, i+2 < len(nums)-2
	output := make([][]int, 0, len(nums)*len(nums))
	for i:=0; i<len(nums)-2; i++ {
		fmt.Printf("The i nums[%v] is %v\n", i, nums[i])
        
        if i>0 && nums[i] == nums[i-1] {
			continue
		}
		
		j, k := i+1, len(nums)-1
        
		for j < k {
            if  j > i+1 && nums[j] == nums[j-1] {
                j++
                continue
            }
            sum := nums[i] +nums[j] + nums[k]
			if sum == 0 {
				output = append(output, []int{nums[i],nums[j],nums[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return output
}
