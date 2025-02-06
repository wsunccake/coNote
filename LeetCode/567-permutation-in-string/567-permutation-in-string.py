from typing import List


class Solution:
    # # slow
    # def checkInclusion(self, s1: str, s2: str) -> bool:
    #     l1 = len(s1)
    #     l2 = len(s2)

    #     if l2 < l1:
    #         return False

    #     char_dict_s1 = {}
    #     for e in s1:
    #         char_dict_s1[e] = char_dict_s1.get(e, 0) + 1

    #     for idx in range(l2 - l1 + 1):
    #         char_dict_s2 = {}
    #         for e in s2[idx:idx+l1]:
    #             char_dict_s2[e] = char_dict_s2.get(e, 0) + 1

    #         if char_dict_s2 == char_dict_s1:
    #             return True

    #     return False

    # # fast
    # def checkInclusion(self, s1: str, s2: str) -> bool:
    #     n = len(s1)
    #     m = len(s2)
    #     start = 0
    #     if n > m:
    #         return False

    #     s1_counts = [0] * 26
    #     s2_counts = [0] * 26
    #     for i in range(n):
    #         s1_counts[ord(s1[i]) - ord('a')] += 1
    #         s2_counts[ord(s2[i]) - ord('a')] += 1

    #     if s1_counts == s2_counts:
    #         return True
    #     for j in range(n, m):
    #         s2_counts[ord(s2[j]) - ord('a')] += 1
    #         s2_counts[ord(s2[start]) - ord('a')] -= 1
    #         start += 1

    #         if s1_counts == s2_counts:
    #             return True

    #     return False

    def checkInclusion(self, s1: str, s2: str) -> bool:
        l1 = len(s1)
        l2 = len(s2)

        if l2 < l1:
            return False

        s1_dict = {chr(97+i): 0 for i in range(26)}
        s2_dict = {chr(97+i): 0 for i in range(26)}

        for i in range(l1):
            s1_dict[s1[i]] = s1_dict.get(s1[i], 0) + 1
            s2_dict[s2[i]] = s2_dict.get(s2[i], 0) + 1

        if s1_dict == s2_dict:
            return True

        start = 0
        for i in range(l1, l2):
            s2_dict[s2[i]] = s2_dict.get(s2[i], 0) + 1
            s2_dict[s2[start]] -= 1
            start += 1
            if s1_dict == s2_dict:
                return True

        return False


if __name__ == '__main__':
    sol = Solution()
    inputs1 = ["ab", "ab", "abc"]
    inputs2 = ["eidbaooo", "eidboaoo", "ccccbbbbaaaa"]
    outputs = [True, False, False]

    for idx in range(len(outputs)):
        ans = sol.checkInclusion(inputs1[idx], inputs2[idx])
        if ans != outputs[idx]:
            print(f'{inputs1[idx]}, {inputs2[idx]}, {ans}, {outputs[idx]}')
