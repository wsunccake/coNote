class Solution:
    # def canPermutePalindrome(self, s: str) -> bool:
    #     oddChars = set()

    #     for c in s:
    #         if c in oddChars:
    #             oddChars.remove(c)
    #         else:
    #             oddChars.add(c)

    #     return len(oddChars) <= 1

    def canPermutePalindrome(self, s: str) -> bool:
        char_dict = {}
        for c in s:
            char_dict[c] = char_dict.get(c, 0)+1

        odd_count = 0
        for v in char_dict.values():
            if v % 2 == 1:
                odd_count += 1
                if odd_count > 1:
                    return False

        return True


if __name__ == '__main__':
    sol = Solution()
    inputs = ["code", "aab", "carerac"]
    outputs = [False, True, True]

    for pos in range(len(outputs)):
        ans = sol.canPermutePalindrome(inputs[pos])
        if ans != outputs[pos]:
            print(f'{inputs[pos]}, {ans}, {outputs[pos]}')
