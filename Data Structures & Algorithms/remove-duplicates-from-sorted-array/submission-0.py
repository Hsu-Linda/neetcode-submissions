# [2,10,10,30,30,30]
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        unique, cur = 0, 1
        
        while cur < len(nums):
            if nums[cur] != nums[unique]:
                unique +=1
                nums[unique] = nums[cur]
                print(unique, nums[cur])
            cur += 1
        return unique+1