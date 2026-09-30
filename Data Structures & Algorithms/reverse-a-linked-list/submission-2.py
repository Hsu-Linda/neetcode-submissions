# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

# 0 , 1,  2 , 3
# cur nxt
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None:
            return None
        cur, nxt = head, head.next
        
        while nxt is not None:
            new_nxt = nxt.next
            nxt.next = cur
            cur = nxt
            nxt = new_nxt

        head.next = None
        return cur
