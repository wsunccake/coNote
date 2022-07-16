from dataclasses import dataclass


class Solution:
    def reverseString(self, s: list[str]) -> None:
        """
        Do not return anything, modify s in-place instead.
        """
        l = len(s)
        h = int((l+1)/2)
        # i = 0
        # j = l - 1 - i
        for i in range(h):
            j = l - 1 - i
            # print(f"i: {s[i]}, j: {s[j]}")
            s[i], s[j] = s[j], s[i]


@dataclass
class Quest():
    inp: list
    sol: list


def checkAns(q: Quest) -> bool:
    sol = Solution()
    sol.reverseString(q.inp)
    res = (q.sol == q.inp)
    if not res:
        print(f'{q.sol} != {q.inp}')
    return res


if __name__ == '__main__':
    q = Quest(inp=["h", "e", "l", "l", "o"],
              sol=["o", "l", "l", "e", "h"])
    checkAns(q)

    q = Quest(inp=["H", "a", "n", "n", "a", "h"],
              sol=["h", "a", "n", "n", "a", "H"])
    checkAns(q)
