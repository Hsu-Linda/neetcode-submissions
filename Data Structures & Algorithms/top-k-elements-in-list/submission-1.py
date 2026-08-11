from collections import Counter
class Solution:
	def topKFrequent(self, nums: List[int], k: int) -> List[int]:
		counter = Counter(nums)
		
		arr = []
		for char, freq in counter.items():
			arr.append([freq, char])
		arr.sort()

		result = []
		while len(result) < k:
			result.append(arr.pop()[1])
		return result