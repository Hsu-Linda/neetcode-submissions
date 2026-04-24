func search(nums []int, target int) int {
	// divided by 2 is 0
	if len(nums) == 0 ||
		( len(nums) == 1 && target != nums[0] ){
		return -1
	}
	
	// out of range
	if nums[len(nums)-1] < target &&
		target < nums[0] {
		return -1
	}

	start := 0
	end := len(nums)-1
	
	for end-start > 1 {
		// is mid
		mid := (end-start)/2 + start
		if nums[mid] == target { return mid }
		
		// [0,1,2] -> [0,1][1,2]
		// [1,2,0] -> [1,2][2,0]
		// [2,0,1] -> [2,0][0,1]

		if nums[start] < nums[mid] && nums[mid] < nums[end] {
			if nums[start] <= target && target < nums[mid] {
				end = mid-1
				continue
			} else if nums[mid] < target && target <= nums[end] {
				start = mid+1
				continue
			} else {
				return -1
			}
		} else if nums[start] > nums[mid] && nums[mid] < nums[end] {
			if nums[mid] < target && target <= nums[end] {
				start = mid+1
				continue
			} else {
				end = mid-1
				continue
			} 
		} else if nums[start] < nums[mid] && nums[mid] > nums[end] {
			if nums[start] <= target && target < nums[mid] {
				end = mid-1
				continue
			} else {
				start = mid+1
				continue
			}
		}

	}
	
	if end - start == 1 {
		if nums[end] == target {
			return end
		} else if nums[start] == target {
			return start
		} else {
			return -1
		}
	} else if nums[start] == target {
		// start == end == target
		return start
	} else {
		// start == end != target
		return -1
	}
}
