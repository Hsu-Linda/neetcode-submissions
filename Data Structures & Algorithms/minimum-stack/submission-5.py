# 1,   2,   
# 1:1  2:1  

# -2  0  -3
# -2  -2 -3
class MinStack:

    def __init__(self):
        self.stack = []

    def push(self, val: int) -> None:
        if not self.stack:
            self.stack.append((val,val))
        else:
            _, min_v = self.stack[-1]
            if val < min_v:
                min_v = val
            self.stack.append((val, min_v))
            
    def pop(self) -> None:
        self.stack.pop()
        
    def top(self) -> int:
        val, _ = self.stack[-1]
        return val
    
    def getMin(self) -> int:
        _, min_v = self.stack[-1]
        return min_v
