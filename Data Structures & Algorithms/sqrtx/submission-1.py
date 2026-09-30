# 9 < 13 < 16
# 3

#  y^2 <= x <= (y+1)^2
# (x)-(x-l)/2 = x/2+l/2
class Solution:
    def mySqrt(self, x: int) -> int:
        l, h = 1, x
        
        while True:
            m = math.floor(h/2+l/2)
            print(m)
            if m*m <= x and (m+1)*(m+1) > x:
                return m
            elif m*m > x:
                h = m
            else:
                l = m
            