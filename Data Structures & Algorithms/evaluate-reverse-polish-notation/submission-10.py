class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = tokens[0:2]
        i = 2

            
        while i < len(tokens):
            if not self.isInt(tokens[i]):
                v1 = int(stack[-2])
                v2 = int(stack[-1])
                tempResult = 0
                if tokens[i] == '+':
                    tempResult = v1+ v2
                elif tokens[i] == '-':
                    tempResult = v1- v2
                elif tokens[i] == '*':
                    tempResult = v1* v2
                elif tokens[i] == '/':
                    tempResult = int(v1/ v2)
                
                stack.pop()
                stack.pop()
                
                stack.append(str(tempResult))
            else:
                stack.append(tokens[i])
            
            i+=1
        
        return int(stack[0])
    
    def isInt(self, v: str) -> bool:
        try:
            int(v)
            return True
        except:
            return False
            