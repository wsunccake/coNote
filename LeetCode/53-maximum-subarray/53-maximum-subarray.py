from dataclasses import dataclass


class Solution:
    # def maxSubArray(self, nums: list[int]) -> int:
    #     # i  j  sum  max    1,  2,  5,  -1,  5,
    #     # 1  1    1    1
    #     # 1  2    3    3
    #     # 1  3    8    8
    #     # 1  4    7    8
    #     # 1  5   12   12
    #     # 2  2    2   12
    #     # 2  3    5   12
    #     l = len(nums)
    #     i = 0
    #     j = 0
    #     k = 0
    #     max_total = nums[0]
    #     for i in range(l):
    #         for j in range(i, l):
    #             sub_total = 0
    #             # print(f'{i} -> {j}')
    #             for k in range(i, j+1):
    #                 # print(f'{i} -> {j} => {k}')
    #                 sub_total += nums[k]
    #             max_total = max(max_total, sub_total)
    #     return max_total

    # def maxSubArray(self, nums: list[int]) -> int:
    #     l = len(nums)
    #     i = 0
    #     j = 0
    #     sub_arr = {}
    #     max_total = nums[0]
    #     for i in range(l):
    #         sub_arr[f'{i}_{i}'] = nums[i]
    #         sub_total = 0
    #         for j in range(i, l):
    #             sub_total += nums[j]
    #             print(f'{nums}: {i} -> {j} {sub_total} {max_total}')
    #             max_total = max(max_total, sub_total)
    #     return max_total

    def maxSubArray(self, nums: list[int]) -> int:
        l = len(nums)
        i = 0
        max_total = nums[0]
        sub_total = 0
        for i in range(l):
            if sub_total > 0:
                sub_total += nums[i]
            else:
                sub_total = nums[i]
            if sub_total > max_total:
                max_total = sub_total
        return max_total

    # def maxSubArray(self, nums: list[int]) -> int:
    #     l = len(nums)
    #     i = 0
    #     max_total = nums[0]
    #     sub_total = 0
    #     index_start = 0
    #     index_end = 0
    #     tmp_start = 0
    #     tmp_end = 0
    #     for i in range(l):
    #         if sub_total > 0:
    #             sub_total += nums[i]
    #             tmp_end = i
    #         else:
    #             sub_total = nums[i]
    #             tmp_start = i
    #             tmp_end = i
    #         if sub_total > max_total:
    #             max_total = sub_total
    #             index_start = tmp_start
    #             index_end = tmp_end
    #     print(f"{index_start} -> {index_end}: {max_total}")
    #     return max_total


@dataclass
class Quest():
    inp: list
    sol: int


def checkAns(q: Quest) -> bool:
    sol = Solution()
    out = sol.maxSubArray(q.inp)
    res = (q.sol == out)
    if not res:
        print(f'{q.inp} -> {q.sol} != {out}')
    return res


if __name__ == '__main__':
    q = Quest(inp=[-2, 1, -3, 4, -1, 2, 1, -5, 4], sol=6)
    checkAns(q)

    q = Quest(inp=[1], sol=1)
    checkAns(q)

    q = Quest(inp=[5, 4, -1, 7, 8], sol=23)
    checkAns(q)

    q = Quest(inp=[-10, -1, -2, -3], sol=-1)
    checkAns(q)
