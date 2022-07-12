class Solution:
    def removeElement(self, nums: 'list[int]', val: 'int') -> 'int':
        l = []
        for i in range(len(nums)):
            if nums[i] != val:
                l.append(nums[i])

        for i in range(len(l)):
            nums[i] = l[i]
        return len(l)


# class Solution:
#     def removeElement(self, nums: 'list[int]', val: 'int') -> 'int':
#         for i in range(len(nums)-1, -1, -1):
#             if nums[i] == val:
#                 nums.pop(i)
#         return len(nums)


if __name__ == '__main__':
    sol = Solution()
    l1 = [3, 2, 2, 3]
    assert sol.removeElement(l1, 3) == 2, 'Fail'
    assert l1[0] == 2, 'Fail'
    assert l1[1] == 2, 'Fail'

    l2 = [0, 1, 2, 2, 3, 0, 4, 2]
    assert sol.removeElement(l2, 2) == 5, 'Fail'
    assert l2[0] == 0, 'Fail'
    assert l2[1] == 1, 'Fail'
    assert l2[2] == 3, 'Fail'
    assert l2[3] == 0, 'Fail'
    assert l2[4] == 4, 'Fail'
