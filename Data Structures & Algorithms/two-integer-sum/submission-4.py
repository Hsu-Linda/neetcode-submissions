class Solution:
	def twoSum(self, nums: List[int], target: int) -> List[int]:
		# put list into dict
		d = {}
		for i, v in enumerate(nums):
			d[v] = i
		
		for i, vi in enumerate(nums):
			diff = target - vi
			if diff in d and d[diff] != i:
				return [i, d[diff]]
		return []