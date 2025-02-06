from typing import List


class Solution:
    # if (word1[i – 1] == word2[j – 1]) dp[i][j] = dp[i – 1][j – 1];
    # if (word1[i – 1] != word2[j – 1]) dp[i][j] = 1 + min(dp[i][j – 1], dp[i – 1][j], dp[i – 1][j – 1]);
    # iterative
    def minDistance(self, word1: str, word2: str) -> int:
        m = len(word1)
        n = len(word2)

        dp = [[0] * (n + 1) for _ in range(m + 1)]

        for i in range(1, m + 1):
            dp[i][0] = i

        for j in range(1, n + 1):
            dp[0][j] = j

        for i in range(1, m + 1):
            for j in range(1, n + 1):
                if word1[i - 1] == word2[j - 1]:
                    dp[i][j] = dp[i - 1][j - 1]
                else:
                    dp[i][j] = min(dp[i - 1][j - 1], dp[i - 1]
                                   [j], dp[i][j - 1]) + 1

        return dp[m][n]

    # # recursive
    # def minDistance(self, word1: str, word2: str) -> int:
    #     m = len(word1)
    #     n = len(word2)

    #     if m == 0:
    #         return n
    #     if n == 0:
    #         return m
    #     dp = [[0] * (n + 1) for _ in range(m + 1)]

    #     for i in range(1, m + 1):
    #         dp[i][0] = i

    #     for j in range(1, n + 1):
    #         dp[0][j] = j

    #     for i in range(1, m + 1):
    #         for j in range(1, n + 1):
    #             if word1[i - 1] == word2[j - 1]:
    #                 dp[i][j] = dp[i - 1][j - 1]
    #             else:
    #                 dp[i][j] = min(dp[i - 1][j - 1], dp[i - 1]
    #                                [j], dp[i][j - 1]) + 1

    #     return dp[m][n]


if __name__ == '__main__':
    sol = Solution()
    inputs1 = ["horse", "intention",]
    inputs2 = ["ros", "execution",]
    outputs = [3, 5]

    for pos in range(len(outputs)):
        ans = sol.minDistance(inputs1[pos], inputs2[pos])
        if ans != outputs[pos]:
            print(f'{inputs1[pos]}, {inputs2[pos]}, {ans}, {outputs[pos]}')
