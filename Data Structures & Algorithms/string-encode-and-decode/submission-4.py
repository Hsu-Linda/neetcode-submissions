class Solution:

	def encode(self, strs: List[str]) -> str:
		result = ""
		
		for s in strs:
			result += str(len(s))
			result += s
			result += ","
		
		return result

	
	# 5Hello,5World,
	# 10Hiiiiiiiii,5World
	def decode(self, s: str) -> List[str]:
		result : List[str] = []
		wordCount = ""
		i = 0
		while i < len(s):
			print(str(i)+":"+s[i])
			if s[i].isdigit():
				wordCount += s[i]
			if s[i+int(wordCount)+1] == ",":
				newI = i+int(wordCount)+1
				result.append(s[i+1:newI])
				i = newI+1
				wordCount = ""
				continue
			else:
				i+=1
			
		return result
