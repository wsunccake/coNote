
from typing import List, Optional


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def getIntersectionNode1(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        lenA = 0
        lenB = 0
        A = headA
        B = headB
        while A:
            lenA = lenA+1
            A = A.next
        while B:
            lenB = lenB+1
            B = B.next

        i = abs(lenA-lenB)
        if lenB > lenA:
            for j in range(i):
                headB = headB.next
        else:
            for j in range(i):
                headA = headA.next

        while headA and headB:
            if headA == headB:
                return headA
            headA = headA.next
            headB = headB.next
        return None

    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        curA = headA
        curB = headB
        while curA != curB:
            curA = curA.next if curA else headB
            curB = curB.next if curB else headA

        return curA


class ListHepler:
    @ staticmethod
    def create(l: List[int]) -> Optional[ListNode]:
        if not l:
            return None
        head = ListNode(l[0])
        cur = head
        for v in l[1:]:
            cur.next = ListNode(v)
            cur = cur.next
        return head

    @ staticmethod
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

# 0
# [2,6,4]
# [1,5]
# 3
# 2
    inputs1 = [
        [4, 1, 8, 4, 5],
        [1, 9, 1, 2, 4],
        [2, 6, 4],
    ]
    inputs2 = [
        [5, 6, 1, 8, 4, 5],
        [3, 2, 4],
        [1, 5],
    ]
    outputs = [
        (8, 2, 3),
        (2, 3, 1),
        (0, 3, 2),

    ]

    for pos in range(len(outputs)):
        head1 = ListHepler.create(inputs1[pos])
        h1 = ListHepler.traverse(head1)
        head2 = ListHepler.create(inputs2[pos])
        h2 = ListHepler.traverse(head2)
        ans = sol.getIntersectionNode(head1, head2)

        if ans != outputs[pos]:
            print(h1, h2, ans, outputs[pos], ans == outputs[pos])
