from dataclasses import dataclass


class Solution:
    def fib(self, n: int) -> int:
        # f[0] = 0
        # f[1] = 1
        # f[2] = f[0] + f[1]
        # f[3] = f[1] + f[2]
        f = []
        f.append(0)
        f.append(1)

        for i in range(2, n+1):
            f.append(f[i-2]+f[i-1])

        return f[n]


@dataclass
class Quest():
    inp: int
    sol: int


def checkAns(q: Quest) -> bool:
    sol = Solution()
    out = sol.fib(q.inp)
    res = (q.sol == out)
    if not res:
        print(f'{q.inp} -> {q.sol} != {out}')
    return res


if __name__ == '__main__':
    q = Quest(inp=2, sol=1)
    checkAns(q)

    q = Quest(inp=3, sol=2)
    checkAns(q)

    q = Quest(inp=4, sol=3)
    checkAns(q)
