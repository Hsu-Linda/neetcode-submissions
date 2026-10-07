"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""

# 3
# 7  3-> 7
# 3 7 4 5 as a key, value we can keep random X BUT value can be the same so we can't use value as a key

#  key might be the liked list address
#  key orignial address, value newe address
#  3XASDSA, 3NJSOIDJF
#




class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        OldNewMap = {}
        dummy = Node(0)
        newPre = dummy
        cur = head
        
        while cur:
            newCur = Node(cur.val)
            OldNewMap[cur] = newCur
            
            newPre.next = newCur
            newPre = newPre.next
            cur = cur.next

        for key, val in OldNewMap.items():
            if key.random:
                val.random = OldNewMap[key.random]

        return dummy.next

            
            







