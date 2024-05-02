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

    @staticmethod
    def search(head: Optional[ListNode], val: int) -> Optional[ListNode]:
        cur = head
        while (cur.val != val) and (cur.next != None):
            cur = cur.next
        return cur


class Solution:
    def deleteNode(self, node):
        """
        :type node: ListNode
        :rtype: void Do not return anything, modify node in-place instead.
        """
        cur = node.next
        node.val = cur.val
        node.next = cur.next


if __name__ == '__main__':
    sol = Solution()
    inputs1 = [[4, 5, 1, 9], [4, 5, 1, 9], [4, 5]]
    inputs2 = [5, 1, 4]
    outputs = [[4, 1, 9], [4, 5, 9], [5]]

    for pos in range(len(outputs)):
        head = ListHepler.create(inputs1[pos])
        h = ListHepler.traverse(head)
        node = ListHepler.search(head, inputs2[pos])
        n = ListHepler.traverse(node)
        sol.deleteNode(node)
        a1 = ListHepler.traverse(head)
        a2 = ListHepler.traverse(ListHepler.create(outputs[pos]))

        if a1 != a2:
            print(h, a1, a2, a1 == a2, n)
