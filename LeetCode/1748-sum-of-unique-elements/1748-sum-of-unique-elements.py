from dataclasses import dataclass


class Solution:
    def sumOfUnique(self, nums: list[int]) -> int:
        tdict = {}
        total = 0
        for n in nums:
            m = tdict.get(n, 0)
            if m == 0:
                total += n
                tdict[n] = 1
            elif m == 1:
                total -= n
                tdict[n] += 1
            else:
                tdict[n] += 1

        return total


@dataclass
class Quest():
    inp: list
    sol: int


def checkAns(q: Quest) -> bool:
    sol = Solution()
    out = sol.sumOfUnique(q.inp)
    res = (q.sol == out)
    if not res:
        print(f'{q.inp} -> {q.sol} != {q.inp}')
    return res


if __name__ == '__main__':
    q = Quest(inp=[1, 2, 3, 2], sol=4)
    checkAns(q)

    q = Quest(inp=[1, 1, 1, 1, 1], sol=0)
    checkAns(q)

    q = Quest(inp=[1, 2, 3, 4, 5], sol=15)
    checkAns(q)
