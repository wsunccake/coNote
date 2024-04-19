from typing import List


class Solution:
    # # normal
    # def containsDuplicate(self, nums: List[int]) -> bool:
    #     num_dict = {}

    #     for n in nums:
    #         if n in num_dict:
    #             return True
    #         num_dict.update({n: 1})

    #     return False

    # fast
    def containsDuplicate(self, nums: List[int]) -> bool:
        if len(set(nums)) == len(nums):
            return False
        return True


if __name__ == '__main__':
    sol = Solution()
    inputs = [
        [1, 2, 3, 1],
        [1, 2, 3, 4],
        [1, 1, 1, 3, 3, 4, 3, 2, 4, 2],
    ]
    outputs = [True, False, True]

    for idx in range(len(outputs)):
        ans = sol.containsDuplicate(inputs[idx])
        if ans != outputs[idx]:
            print(f'{inputs[idx]}, {ans}, {outputs[idx]}')
