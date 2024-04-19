from typing import List


class Solution:
    # # fast
    # def isAnagram(self, s: 'str', t: 'str') -> 'bool':
    #     return ''.join(sorted(s)) == ''.join(sorted(t))

    # # slow
    # def isAnagram(self, s: 'str', t: 'str') -> 'bool':
    #     return sorted(s) == sorted(t)

    def isAnagram(self, s: 'str', t: 'str') -> 'bool':
        if len(s) != len(t):
            return False

        char_dict = {}
        for char in s:
            # if e in char_dict:
            #     char_dict[char] += 1
            # else:
            #     char_dict[char] = 1
            char_dict[char] = char_dict.get(char, 0) + 1

        for char in t:
            if char in char_dict:
                char_dict[char] -= 1
            else:
                return False

        # for k, v in char_dict.items():
        #     if v != 0:
        #         return False

        for value in char_dict.values():
            if value != 0:
                return False

        return True


if __name__ == '__main__':
    sol = Solution()
    inputs1 = ["anagram", "rat",]
    inputs2 = ["nagaram", "car",]
    outputs = [True, False]

    for idx in range(len(outputs)):
        ans = sol.isAnagram(inputs1[idx], inputs2[idx])
        if ans != outputs[idx]:
            print(f'{inputs1[idx]}, {inputs2[idx]}, {ans}, {outputs[idx]}')
