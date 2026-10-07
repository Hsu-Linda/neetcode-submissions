# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=, NoDefault0, next=None):
#         self.val = val
#         self.next = next


# two pointer 
# find the 1th node from the end
# find n+1th node from the end
# 1,   2,    3,    4,
#            2th  1th
#      n+1th        



class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode()
        dummy.next = head
        n1Node, firstNode = dummy, dummy
        for i in range(n):
            if firstNode.next is not None:
                firstNode = firstNode.next
            else:
                return None

        while firstNode.next is not None:
            if n1Node is None:
                return None
            n1Node = n1Node.next
            firstNode = firstNode.next

        if n1Node is not None and n1Node.next is not None:
            n1Node.next = n1Node.next.next

        return dummy.next

        

