class Solution:
    def longestPalindrome(self, s: str) -> int:
        string_dict = {}
        for c in s:
            string_dict[c] = string_dict.get(c, 0) + 1
        r = 0
        q = 0
        for _, v in string_dict.items():
            quotient = int(v / 2)
            remainder = v % 2
            r += quotient
            if remainder == 1:
                q = 1
        return r * 2 + q


if __name__ == '__main__':
    sol = Solution()
    a = [
        ("abccccdd", 7),
        ("a", 1),
        ("ab", 1),
        ("aa", 2),
    ]
    for k in a:
        assert sol.longestPalindrome(k[0]) == k[1], f'{k[0]} ans != {k[1]}'
