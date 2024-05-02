
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
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        dummy = ListNode(val=0, next=head)
        prev = dummy
        cur = dummy.next
        pres_head = None
        pres_tail = None

        while cur != None:
            if cur.val >= x:
                if pres_head == None:
                    pres_head = prev
                pres_tail = cur
            else:
                if pres_head != None:
                    tmp_head = pres_head.next
                    tmp_tail = cur.next
                    pres_head.next = cur
                    cur.next = tmp_head
                    pres_tail.next = tmp_tail
                    pres_head = None

            prev = cur
            cur = cur.next

        return dummy.next

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
    inputs1 = [[1, 4, 3, 2, 5, 2], [2, 1]]
    inputs2 = [3, 2]
    outputs = [[1, 2, 2, 4, 3, 5], [1, 2]]

    for pos in range(len(outputs)):
        head = ListHepler.create(inputs1[pos])
        h = ListHepler.traverse(head)
        ans = sol.partition(head, inputs2[pos])
        a1 = ListHepler.traverse(ans)
        a2 = ListHepler.traverse(ListHepler.create(outputs[pos]))

        if a1 != a2:
            print(h, a1, a2, a1 == a2)
