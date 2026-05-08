// [1, 2, 4, 6]
// [2*4*6, 1*4*6, 1*2*6, 1*2*4]
// O(n*(n-1)) -> O(n^2)

// [1*2*4*6] n
// [all/1, all/2, all/4, all/6] n
// n+n -> 0(n)


// [-1,0,1,2,3]
// all = 0 we can know there are one or more 0 inside
// if there more than one 0, output each of the ele is 0
// but if just only one 0 in nums 
// except itself others are 0


func productExceptSelf(nums []int) []int {
	all := 1
	zeroCounts := 0
	firstZeroIndex := -1
	for i, v := range nums {
		if v == 0 {
			zeroCounts ++
			if zeroCounts == 1 {
				firstZeroIndex = i
			}
			continue
		}
		all = all *v
	}

	output := make([]int, 0, len(nums))
	if zeroCounts > 0 {
		for i:= 0; i<len(nums);i++ {
			if zeroCounts == 1 && i == firstZeroIndex{
				output = append(output, all)
				continue
			}
			output = append(output, 0)
		}
		return output
	}
	for _, v := range nums {
		if v == 0 {

		}
		output = append(output, all/v) 
	}

	return output
}
