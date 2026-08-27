# 1, 2, 4, 5, 6
# 1, 2, 2, 3, 3 
class Solution:
    def numRescueBoats(self, people: List[int], limit: int) -> int:
        people.sort()
        left, right = 0, len(people)-1
        boat_count = 0
        while left < right:
            if people[left]+people[right] > limit:
                right -=1
            else:
                left +=1
                right -=1
            boat_count +=1
        if left == right:
            boat_count+=1
        return boat_count