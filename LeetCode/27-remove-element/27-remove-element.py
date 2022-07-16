from dataclasses import dataclass


class Solution:
    def removeElement(self, nums: list[int], val: int) -> int:
        i = 0
        j = 0
        l = len(nums)
        while (i+j < l):
            if nums[i] == val:
                if nums[i+j] == val:
                    j += 1
                else:
                    nums[i], nums[i+j] = nums[i+j], nums[i]
                    i += 1
                    j = 0
            else:
                i += 1
        return i


#class Solution:
#    def removeElement(self, nums: 'list[int]', val: 'int') -> 'int':
#        l = []
#        for i in range(len(nums)):
#            if nums[i] != val:
#                l.append(nums[i])
#        for i in range(len(l)):
#            nums[i] = l[i]
#        return len(l)


# class Solution:
#     def removeElement(self, nums: 'list[int]', val: 'int') -> 'int':
#         for i in range(len(nums)-1, -1, -1):
#             if nums[i] == val:
#                 nums.pop(i)
#         return len(nums)


@dataclass
class Quest():
    inp_nums: list
    inp_val: int
    sol_nums: list
    sol_val: int


def checkAns(q: Quest) -> bool:
    sol = Solution()
    out = sol.removeElement(q.inp_nums, q.inp_val)
    res = (set(q.inp_nums[:out]) == set(q.sol_nums[:q.sol_val]))
    if not res:
        print(f'{set(q.inp_nums[:out])} != {set(q.sol_nums[:q.sol_val])}')
    return res


if __name__ == '__main__':
#    sol = Solution()
#    l1 = [3, 2, 2, 3]
#    assert sol.removeElement(l1, 3) == 2, 'Fail'
#    assert l1[0] == 2, 'Fail'
#    assert l1[1] == 2, 'Fail'
#
#    l2 = [0, 1, 2, 2, 3, 0, 4, 2]
#    assert sol.removeElement(l2, 2) == 5, 'Fail'
#    assert l2[0] == 0, 'Fail'
#    assert l2[1] == 1, 'Fail'
#    assert l2[2] == 3, 'Fail'
#    assert l2[3] == 0, 'Fail'
#    assert l2[4] == 4, 'Fail'

    q = Quest(inp_nums=[3, 2, 2, 3], inp_val=3,
              sol_nums=[2, 2, None, None], sol_val=2)
    checkAns(q)

    q = Quest(inp_nums=[0, 1, 2, 2, 3, 0, 4, 2], inp_val=2,
              sol_nums=[0, 1, 4, 0, 3, None, None, None], sol_val=5)
    checkAns(q)
