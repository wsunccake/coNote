
from typing import List, Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def deleteDuplicatesUnsorted(self, head: ListNode) -> ListNode:
        dummy = ListNode(val=0, next=head)
        duplicated_dict = {}

        prev = dummy
        cur = dummy.next
        while cur != None:
            count = duplicated_dict.get(cur.val, 0) + 1
            duplicated_dict[cur.val] = count
            if count > 1:
                prev.next = cur.next
            else:
                prev = cur
            cur = cur.next

        prev = dummy
        cur = dummy.next
        while cur != None:
            if duplicated_dict[cur.val] > 1:
                prev.next = cur.next
            else:
                prev = cur
            cur = cur.next

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
    inputs = [[1, 2, 3, 2], [2, 1, 1, 2], [3, 2, 2, 1, 3, 2, 4]]
    outputs = [[1, 3], [], [1, 4]]

    for pos in range(len(outputs)):
        head = ListHepler.create(inputs[pos])
        h = ListHepler.traverse(head)
        ans = sol.deleteDuplicatesUnsorted(head)
        a1 = ListHepler.traverse(ans)
        a2 = ListHepler.traverse(ListHepler.create(outputs[pos]))

        if a1 != a2:
            print(h, a1, a2, a1 == a2)
