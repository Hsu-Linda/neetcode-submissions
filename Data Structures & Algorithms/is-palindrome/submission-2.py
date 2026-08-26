class Solution:
    def isPalindrome(self, s: str) -> bool:
        left, right = 0, len(s)-1
        while left < right:
            print(left)
            while left < right and not s[left].isalnum():
                print(s[left])
                left +=1
                
            while left < right and not s[right].isalnum():
                right -=1
            print(s[left].lower())
            print(s[right].lower())
            if s[left].lower() != s[right].lower():
                return False
            left +=1
            right -=1
        return True