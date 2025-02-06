
class Solution:
    def is_permutation(self, s1, s2):
        if len(s1) != len(s2):
            return False

        count = {}
        for s in s1:
            if s in count:
                count[s] += 1
            else:
                count[s] = 1

        for s in s2:
            if s in count:
                count[s] -= 1
            else:
                count[s] = -1

        for v in count.values():
            if v != 0:
                return False

        return True


if __name__ == '__main__':
    inputs1 = ["testest", "hello", "", ""]
    inputs2 = ["estxest", "oellh", "", "abc"]
    outputs = [False, True, True, False]
    sol = Solution()

    for i in range(len(outputs)):
        if sol.is_permutation(inputs1[i], inputs2[i]) != outputs[i]:
            print(inputs1[i],  inputs2[i], outputs[i],
                  sol.is_permutation(inputs1[i], inputs2[i],))
