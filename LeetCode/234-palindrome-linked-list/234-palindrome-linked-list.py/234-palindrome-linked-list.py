
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
    # slow
    # def isPalindrome(self, head: Optional[ListNode]) -> bool:
    #     # if head.next == None:
    #     if not head.next:
    #         return True

    #     middle = head
    #     tail = head

    #     while tail.next:
    #         middle = middle.next
    #         tail = tail.next
    #         if tail.next:
    #             tail = tail.next

    #     dummy = ListNode(0, None)
    #     reversed = dummy

    #     while middle:
    #         next1 = reversed.next
    #         next2 = middle.next

    #         reversed.next = middle
    #         reversed.next.next = next1

    #         middle = next2

    #     reversed = dummy.next
    #     cur = head

    #     while reversed:
    #         if cur.val != reversed.val:
    #             return False
    #         reversed = reversed.next
    #         cur = cur.next
    #     return True

    # fast
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        if not head.next:
            return True

        cur = head
        l = []
        while cur:
            l.append(cur.val)
            cur = cur.next

        return l == l[::-1]


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
    inputs = [
        [1, 2, 2, 1],
        [1, 2],
        [1, 1, 2, 2],
        [1, 2, 3, 1, 2, 3],
        [1, 2, 1, 1, 2, 1],
        [1],
        [1, 2, 3],
        [1, 2, 3, 4],
    ]
    outputs = [
        True,
        False,
        False,
        False,
        True,
        True,
        False,
        False
    ]

    for pos in range(len(outputs)):
        head = ListHepler.create(inputs[pos])
        h = ListHepler.traverse(head)
        ans = sol.isPalindrome(head)

        if ans != outputs[pos]:
            print(h, ans, outputs[pos], ans == outputs[pos])
