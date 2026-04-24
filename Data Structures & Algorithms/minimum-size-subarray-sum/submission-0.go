// [2,] 2 
// [2,1] 3
// [2,1,5] 8
// [2,1,5,1] 9
// [2,1,5,1,5] 14
// [1,5,1,5] 12
// [5,1,5] 11

func minSubArrayLen(target int, nums []int) int {
	
	window := make([]int, 0, len(nums))
	windowSum := 0
	minLeng := len(nums)
	
	for i:=0; i<len(nums); i++ {
		window = append(window, nums[i])
		windowSum += nums[i]
		
		if windowSum < target {	
			continue
		}
		for windowSum - window[0] >= target {
			windowSum -= window[0]
			window = window[1:]
		}
		if len(window) < minLeng {
			minLeng = len(window)
		}
		
	}
	if windowSum <target {
		return 0
	}
	return minLeng
}
