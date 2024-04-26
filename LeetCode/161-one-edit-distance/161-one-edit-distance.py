class Solution:
    def isOneEditDistance(self, s: 'str', t: 'str') -> 'bool':
        if len(s) < len(t):
            s, t = t, s

        m, n = len(s), len(t)

        if m-n > 1:
            return False

        diff_count = 0
        if m == n:
            for i in range(n):
                if s[i] != t[i]:
                    diff_count += 1
        else:
            i = 0
            j = 0
            while j < n:
                while i < m:
                    if s[i] == t[j]:
                        break
                    else:
                        diff_count += 1
                    i += 1
                j += 1

            diff_count -= m - i

        if diff_count > 1:
            return False

        return True


if __name__ == '__main__':
    sol = Solution()
    inputs1 = ["ab", "cab", "1203", "1234"]
    inputs2 = ["acb", "ad", "1213", "1253"]
    outputs = [True, False, True, False]

    for pos in range(len(outputs)):
        ans = sol.isOneEditDistance(inputs1[pos], inputs2[pos])
        if ans != outputs[pos]:
            print(f'{inputs1[pos]}, {inputs2[pos]}, {ans}, {outputs[pos]}')
