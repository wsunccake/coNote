
class Solution:
    def is_unique(self, s):
        seen = {}
        for i in s:
            if i in seen:
                return False
            else:
                seen[i] = True

        return True


if __name__ == '__main__':
    inputs = ["abcde", "hello", "apple", "kite", "padle"]
    outputs = [True, False, False, True, True]
    sol = Solution()

    for i in range(len(inputs)):
        if sol.is_unique(inputs[i]) != outputs[i]:
            print(inputs[i], outputs[i], sol.is_unique(inputs[i]))
