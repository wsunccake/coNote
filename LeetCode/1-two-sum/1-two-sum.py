class Solution:
    def twoSum(self, nums: 'List[int]', target: 'int') -> 'List[int]':
        result = [0, 0]
        for num in nums:
            goal = target - num
            num_index = nums.index(num)
            nums[num_index] = None
            if goal in nums:
                result = [num_index, nums.index(goal)]
                break
            nums[num_index] = num
        return result


# class Solution:
#     def twoSum(self, nums: 'List[int]', target: 'int') -> 'List[int]':
#         result = [0, 0]
#         for num_index, num in enumerate(nums[:-1]):
#             goal = target - num
#             shift_index = num_index + 1
#             if goal in nums[shift_index:]:
#                 result = [num_index, nums[shift_index:].index(goal) + shift_index]
#                 break
#         return result


# class Solution:
#     def twoSum(self, nums: 'List[int]', target: 'int') -> 'List[int]':
#         tmp_dict = {}
#         for index, num in enumerate(nums):
#             goal = target - num
#             if goal in tmp_dict:
#                 return [tmp_dict[goal], index]
#             tmp_dict[num] = index
#         return [0, 0]


if __name__ == '__main__':
    sol = Solution()

    q = [1, 2, 3]
    target = 6
    ans1 = 0
    ans2 = 0
    assert sol.twoSum(q, target) == [ans1, ans2], 'Fail'

    q = [3, 2, 4]
    target = 6
    ans1 = 1
    ans2 = 2
    assert sol.twoSum(q, target) == [ans1, ans2], 'Fail'

    q = [3, 3]
    target = 6
    ans1 = 0
    ans2 = 1
    assert sol.twoSum([3, 3], 6) == [0, 1], 'Fail'
