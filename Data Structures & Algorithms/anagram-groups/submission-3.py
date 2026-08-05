from collections import Counter
from collections import defaultdict
class Solution:
	def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
		# key Counter value []str
		result = defaultdict(list)
		for i in strs:
			result[tuple(sorted(Counter(i).items()))].append(i)
		return list(result.values())