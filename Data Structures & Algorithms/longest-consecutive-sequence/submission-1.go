// 2,20,4,10,3,4,5

// start 2 20 4 10  
func longestConsecutive(nums []int) int {
	numSet := set(nums)
	longest := 0
	for _, num := range nums {
		if _, ok := numSet[num-1]; ok {
			// not sequence start
			continue
		}
		
		length := 0
		for {
			if _, ok := numSet[num+length]; ok {
				length ++
			} else {
				break
			}
		}
		if length > longest  {
			longest = length
		}
	}
	return longest
}

func set(nums []int) map[int]struct{} {
	output := make(map[int]struct{})
	for _, num := range nums {
		output[num] = struct{}{}
	}
	return output
}