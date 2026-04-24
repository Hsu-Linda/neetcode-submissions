
// foreach row use binary search  -> m logn 
// as we know the second condition first integer if greater than the last ele in previous row 
// seem it as big slice [1,2,4,8, 10,11,12,13,.....]
// slice is m*n 
// log(m*n)

type Matrix [][]int

func searchMatrix(matrix [][]int, target int) bool {
	l := 0
	h := len(matrix) * len(matrix[0]) -1
	for l<=h{
		mid := (l+h)/2
		if getValue(matrix,mid) == target {
			return true
		} else if getValue(matrix, mid) < target {
			if mid+1 > h {
				return false
			} 
			l = mid+1
		} else {
			if mid-1 <l {
				return false
			}
			h = mid-1
		}
	}
	return true
}
func getValue(matrix [][]int,nth int) int{
	eleCountsPerRow := len((matrix)[0])
	row := nth /eleCountsPerRow
	col := nth % eleCountsPerRow
	return (matrix)[row][col]
}
