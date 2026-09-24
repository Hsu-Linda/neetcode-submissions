class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        i = 0
        result = ""
        while i < len(word1) and i < len(word2):
            result += word1[i]
            result += word2[i]
            i+=1
        

        if i < len(word1):
            result += word1[i:]
        elif i < len(word2):
            result += word2[i:]

        return result