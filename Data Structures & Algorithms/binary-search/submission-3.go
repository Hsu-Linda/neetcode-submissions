func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	
	l := 0
	h := len(nums)-1
	var mid int 
	for l <= h {
		mid = (l+h)/2
		if nums[mid] == target {
			return mid
		}
		
		if l == h {
			return -1
		}

		if target < nums[mid] {
			h = mid-1
		} else {
			l = mid+1
			
		}
	}
	return -1
}
