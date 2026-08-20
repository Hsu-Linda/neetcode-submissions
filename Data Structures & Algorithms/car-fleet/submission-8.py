# 10
#  (0,1) (1,2) (4,2) (7,1)
#   1  3  6  8
#   2  5  8  9
#   3  7  10 10
#      9
#      10
class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        posiAndSpeed = []
        for i in range(len(position)):
            posiAndSpeed.append((position[i], speed[i]))
        posiAndSpeed.sort()

        fleetCount = 1
        fleetUsedTime = (target-posiAndSpeed[-1][0])/posiAndSpeed[-1][1]
        while posiAndSpeed:
            curPos, curSpe = posiAndSpeed.pop()
            if curPos+curSpe*fleetUsedTime < target:
                fleetCount += 1
                fleetUsedTime = (target-curPos)/curSpe
        return fleetCount