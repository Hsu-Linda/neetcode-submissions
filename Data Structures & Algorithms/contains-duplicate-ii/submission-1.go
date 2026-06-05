func containsNearbyDuplicate(nums []int, k int) bool {
    window := make([]int, 0 , k+1)
    windowSet := make(map[int]struct{})
    
    for i, _  := range nums {
        if len(window) == k+1 {
            delete(windowSet, window[0])
            
            newWindow := make([]int,k, k+1)
            copy(newWindow, window[1:]) 
            window = newWindow
        }
        window = append(window, nums[i])
        windowSet[nums[i]] = struct{}{}
        if len(window) > len(windowSet) {
            return true
        }
    }
    return false
}
