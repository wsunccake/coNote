
from typing import List, Optional


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        carry = 0
        dummy = ListNode(0, None)
        cur = dummy

        while l1 or l2:
            v1 = 0
            v2 = 0
            if l1:
                v1 = l1.val
                l1 = l1.next
            if l2:
                v2 = l2.val
                l2 = l2.next

            sum = v1 + v2 + carry
            remainder = sum % 10
            carry = sum // 10

            cur.next = ListNode(remainder, None)
            cur = cur.next

        if carry:
            cur.next = ListNode(carry, None)

        return dummy.next


class ListHepler:
    @staticmethod
    def create(l: List[int]) -> Optional[ListNode]:
        if not l:
            return None
        head = ListNode(l[0])
        cur = head
        for v in l[1:]:
            cur.next = ListNode(v)
            cur = cur.next
        return head

    @staticmethod
    def traverse(head: Optional[ListNode]) -> str:
        s = ""
        if head == None:
            return s

        cur = head
        while cur.next != None:
            s += f"{cur.val}->"
            cur = cur.next
        s += f"{cur.val}"
        return s


if __name__ == '__main__':
    sol = Solution()
    inputs1 = [[2, 4, 3], [0], [9, 9, 9, 9, 9, 9, 9]]
    inputs2 = [[5, 6, 4], [0], [9, 9, 9, 9]]
    outputs = [[7, 0, 8], [0], [8, 9, 9, 9, 0, 0, 0, 1]]

    for pos in range(len(outputs)):
        head1 = ListHepler.create(inputs1[pos])
        head2 = ListHepler.create(inputs2[pos])
        h1 = ListHepler.traverse(head1)
        h2 = ListHepler.traverse(head2)
        ans = sol.addTwoNumbers(head1, head2)
        a1 = ListHepler.traverse(ans)
        a2 = ListHepler.traverse(ListHepler.create(outputs[pos]))

        if a1 != a2:
            print(a1, a2, a1 == a2)
