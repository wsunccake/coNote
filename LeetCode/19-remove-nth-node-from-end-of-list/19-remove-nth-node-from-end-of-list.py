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


class ListHelper:
    def create(l: List[int]) -> Optional[ListNode]:
        if not l:
            return None

        head = ListNode(l[0])

        cur = head
        for v in l[1:]:
            cur.next = ListNode(v)
            cur = cur.next

        return head

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


class Solution:
    # # normal
    # def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
    #     total = 0
    #     cur = head
    #     while cur != None:
    #         total += 1
    #         cur = cur.next

    #     prev = None
    #     cur = head
    #     for _ in range(1, total+1-n):
    #         prev = cur
    #         cur = cur.next

    #     if cur == head:
    #         if cur.next is None:
    #             head = None
    #         else:
    #             head = cur.next
    #     else:
    #         prev.next = cur.next

    #     return head

    # dummy node
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(0)
        dummy.next = head

        total = 0
        cur = head
        while cur != None:
            total += 1
            cur = cur.next

        prev = dummy
        cur = dummy.next
        for _ in range(1, total+1-n):
            prev = cur
            cur = cur.next

        prev.next = cur.next
        return dummy.next

    # two pointer
    # def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
    #     dummy = ListNode(0)
    #     dummy.next = head
    #     first = dummy
    #     second = dummy

    #     for _ in range(n+1):
    #         first = first.next

    #     while first != None:
    #         first = first.next
    #         second = second.next

    #     second.next = second.next.next

    #     return dummy.next


if __name__ == '__main__':
    sol = Solution()
    inputs1 = [[1, 2, 3, 4, 5], [1], [1, 2], [1, 2]]
    inputs2 = [2, 1, 1, 2]
    outputs = [[1, 2, 3, 5], [], [1], [2]]

    for pos in range(len(outputs)):
        head = ListHelper.create(inputs1[pos])
        h = ListHelper.traverse(head)
        ans = sol.removeNthFromEnd(head, inputs2[pos])
        a1 = ListHelper.traverse(ans)
        a2 = ListHelper.traverse(ListHelper.create(outputs[pos]))
        if a1 != a2:
            print(h, a1, a2, a1 == a2)
        # if ans != outputs[pos]:
        #     print(f'{inputs1[pos]}, {inputs2[pos]}, {ans} {outputs[pos]}')
