# [1,2,4,6]
# 246  146 126 124
#  642  64  6      -> [1, 6, 64,642]
#         1   12  124  
class Solution:
	def productExceptSelf(self, nums: List[int]) -> List[int]:
		pre :list[int] = [1,]  # reverse later
		suf :list[int] = [1,]
		for i in range(len(nums)-1):
			suf.append(suf[-1]*nums[i])
		for i in range(len(nums)-1,0,-1):
			pre.append(pre[-1]*nums[i])
		
		pre.reverse()

		return [x*y for x, y in zip(pre,suf)]