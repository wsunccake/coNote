
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
        dummy1 = ListNode(0, None)
        head = dummy1
        cur = l1

        while cur:
            tmp_next = cur.next
            cur.next = head.next
            head.next = cur
            cur = tmp_next

        l1 = dummy1.next

        dummy2 = ListNode(0, None)
        head = dummy2
        cur = l2

        while cur:
            tmp_next = cur.next
            cur.next = head.next
            head.next = cur
            cur = tmp_next

        l2 = dummy2.next

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

            cur.next = ListNode(remainder, cur.next)

        if carry:
            cur.next = ListNode(carry, cur.next)

        return dummy.next

    def reversed(self, l:  Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0, None)
        head = dummy
        cur = l

        while cur:
            tmp_next = cur.next
            cur.next = head.next
            head.next = cur
            cur = tmp_next

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
    inputs1 = [[7, 2, 4, 3], [2, 4, 3], [0], [5]]
    inputs2 = [[5, 6, 4], [5, 6, 4], [0], [5]]
    outputs = [[7, 8, 0, 7], [8, 0, 7], [0], [1, 0]]

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
