# -4 -1 -1 0 1 2 
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        print(nums)
        value_to_index = {}
        for index, value in enumerate(nums):
            value_to_index[value] = index
        
        result = []

        used_i_stack = []
        for i in range(len(nums)):
            if used_i_stack and nums[i] == used_i_stack[-1]:
                continue
            used_i_stack.append(nums[i])
            used_j_stack = []
            for j in range(i+1,len(nums)):
                if used_j_stack and nums[j] == used_j_stack[-1]:
                    continue
                used_j_stack.append(nums[j])
                if 0-(nums[i]+nums[j]) in value_to_index :
                    k = value_to_index[0-(nums[i]+nums[j])]
                    if k > j:
                        result.append([nums[i], nums[j], nums[k]])
        
        return result
        
