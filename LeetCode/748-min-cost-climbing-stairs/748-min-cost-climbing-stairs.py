# dp(n) = min(dp(n-1), dp(n-2) + cost(n-1))

class Solution:
    def minCostClimbingStairs(self, cost: list[int]) -> int:
        dp = [cost[0], cost[1]]
        n = len(cost)
        for i in range(2, n):
            dp.append(min(dp[i-1], dp[i-2]) + cost[i])
        return min(dp[n-1], dp[n-2])

# faster
# class Solution:
#     def minCostClimbingStairs(self, cost: list[int]) -> int:
#         n = len(cost)
#         for i in range(2, n):
#             cost[i] = (min(cost[i-1], cost[i-2]) + cost[i])
#         return min(cost[n-1], cost[n-2])


if __name__ == '__main__':
    sol = Solution()
    a = [
        ([10, 15, 20], 15),
        ([1, 100, 1, 1, 1, 100, 1, 1, 100, 1], 6),
    ]
    for k in a:
        assert sol.minCostClimbingStairs(k[0]) == k[1], f'{k[0]} ans != {k[1]}'
