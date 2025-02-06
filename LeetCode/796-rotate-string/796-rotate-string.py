from typing import List


class Solution:
    # # trick
    # def rotateString(self, s: str, goal: str) -> bool:
    #     if len(s) > len(goal):
    #         return False

    #     if s in goal+goal:
    #         return True
    #     return False

    # normal
    def rotateString(self, s: str, goal: str) -> bool:
        if len(s) > len(goal):
            return False

        l = len(s)
        for i in range(l):
            if s[i:]+s[:i] == goal:
                return True
        return False

    # # slow
    # def rotateString(self, s: str, goal: str) -> bool:
    #     if len(s) > len(goal):
    #         return False

    #     l = len(s)
    #     target = goal + goal
    #     for i in range(l):
    #         if target[i:i+l] == s:
    #             return True
    #     return False

    # # slow
    # def rotateString(self, s: str, goal: str) -> bool:
    #     if len(s) > len(goal):
    #         return False

    #     l = len(s)
    #     target = goal + goal
    #     for i in range(l):
    #         for j in range(l):
    #             if target[i+j] != s[j]:
    #                 break
    #         if target[i+j] == s[j]:
    #             return True
    #     return False


if __name__ == '__main__':
    sol = Solution()
    inputs1 = ["abcde", "abcde", "aa"]
    inputs2 = ["cdeab", "abced", "a"]
    outputs = [True, False, False]

    for pos in range(len(outputs)):
        ans = sol.rotateString(inputs1[pos], inputs2[pos])
        if ans != outputs[pos]:
            print(f'{inputs1[pos]}, {inputs2[pos]}, {ans} {outputs[pos]}')
