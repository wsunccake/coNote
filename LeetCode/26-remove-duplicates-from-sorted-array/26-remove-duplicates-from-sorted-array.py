from dataclasses import dataclass


# class Solution:
#     def removeDuplicates(self, nums: 'list[int]') -> 'int':
#         count = 0
#         current = None
#         for num in nums:
#             if current != num:
#                 nums[count] = num
#                 count += 1
#                 current = num
#
#         return count


# class Solution:
#     def removeDuplicates(self, nums: 'list[int]') -> 'int':
#         count = 0
#         current = None
#         for _, num in enumerate(nums):
#             if current != num:
#                 nums[count] = num
#                 count += 1
#                 current = num
#         return count


# class Solution:
#     def removeDuplicates(self, nums: 'list[int]') -> 'int':
#         count = 0
#         current = None
#         i = 0
#         while i < len(nums):
#             if current != nums[i]:
#                 nums[count] = nums[i]
#                 count += 1
#                 current = nums[i]
#             i += 1
#         return count

class Solution:
    def removeDuplicates(self, nums: 'list[int]') -> 'int':
        current = nums[0]
        i = 1
        j = 1
        while i < len(nums):
            if current != nums[i]:
                nums[j] = nums[i]
                j += 1
                current = nums[i]
            i += 1
        return j


@dataclass
class Quest():
    inp: list
    sol: list


def checkAns(q: Quest) -> bool:
    sol = Solution()
    out = sol.removeDuplicates(q.inp)
    res = (q.sol == q.inp[:out])
    if not res:
        print(f'{q.sol} != {q.inp}')
    return res


if __name__ == '__main__':
    q = Quest(inp=[1, 1, 2], sol=[1, 2])
    checkAns(q)

    q = Quest(inp=[0, 0, 1, 1, 1, 2, 2, 3, 3, 4], sol=[0, 1, 2, 3, 4])
    checkAns(q)

    q = Quest(inp=[0], sol=[0])
    checkAns(q)
