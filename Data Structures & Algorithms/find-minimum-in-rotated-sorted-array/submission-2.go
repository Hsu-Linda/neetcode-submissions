// the min one

// [1,2,3,4,5,6] -> [1,2,3] [4] [5,6] -> [1][2][3] -> [1]
// [6,1,2,3,4,5] -> [6,1,2] [3] [4,5] -> [6][1][2] -> [1]
// [5,6,1,2,3,4] 
// [4,5,6,1,2,3]
// if mid bigger than both side, then it is the biggest one, and also the min is mid+1
// [3,4,5,6,1,2] -> [3,4,5] [6] [1,2] -> [1][2] -> [1]
// [2,3,4,5,6,1]
// until the len -> 1
// which left side and right side both bigger than itself -> mid is min
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	
	current := nums

	for len(current) > 1 {
		if len(current) == 2 {
			if current[0] < current[1] {
				return current[0]	
			}
			return current[1]
		}

		mid := len(current)/2
		left := mid -1
		right := mid +1

		if current[mid] < current[right] &&
			current[mid] < current[left] {
			return current[mid]
		} else if current[mid] > current[right] &&
			current[mid] > current[left] {
			return current[right]
		} 
		
		if current[0] < current[mid] &&
			current[mid] < current[len(current)-1] {
				return current[0]
		} else if current[0] < current[mid] &&
			current[mid] > current[len(current)-1]  {
			current = current[mid+1:]
		} else {
			current = current[:mid]
		}
	}

	return current[0]
}
