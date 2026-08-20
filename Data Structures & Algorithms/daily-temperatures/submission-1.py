# (30, 0) 38
class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack = []
        result = [0]*len(temperatures)
        i = 0
        while i < len(temperatures):
            if not stack:
                stack.append((temperatures[i], i))
                i +=1
                continue
            lastTem, lastIndex = stack[-1]
            if lastTem >= temperatures[i]:
                stack.append((temperatures[i], i))
                i +=1
                continue
            while lastTem < temperatures[i]:
                result[lastIndex] = i-lastIndex
                stack.pop()
                if stack:
                    lastTem, lastIndex = stack[-1]
                else:
                    break;
            stack.append((temperatures[i], i))

        return result