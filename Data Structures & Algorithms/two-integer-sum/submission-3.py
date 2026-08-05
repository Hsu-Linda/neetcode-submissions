class Solution:
	def twoSum(self, nums: List[int], target: int) -> List[int]:
		# put list into dict
		d = {}
		for i, v in enumerate(nums):
			d[v] = i
		
		for i, vi in enumerate(nums):
			j = d.get((target-vi),-1)
			if -1 == j or i == j:
				continue
			else:
				return [i, j]
		return []