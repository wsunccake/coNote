from dataclasses import dataclass


# class Solution:
#     def convertToTitle(self, columnNumber: int) -> str:
#         number_to_char_dict = {0: 'Z'}
#         for i in range(65, 91):
#             number_to_char_dict[i-64] = chr(i)
#         quotient, remainder = divmod(columnNumber, 26)
#         res = f'{number_to_char_dict[remainder]}'
#         while (quotient > 26):
#             if (remainder == 0):
#                 quotient -= 1
#             quotient, remainder = divmod(quotient, 26)
#             res = f'{number_to_char_dict[remainder]}{res}'
#         if (remainder == 0):
#             quotient -= 1
#         if (quotient != 0):
#             res = f'{number_to_char_dict[quotient]}{res}'
#         return res

class Solution:
    def convertToTitle(self, columnNumber: int) -> str:
        number_to_char_dict = {0: 'Z'}
        for i in range(65, 91):
            number_to_char_dict[i-64] = chr(i)
        s = ""
        while columnNumber > 0:
            t = columnNumber % 26
            # if t == 0:
            #     s += 'Z'
            # else:
            # s += chr(64+t)
            s += number_to_char_dict[t]
            columnNumber = (columnNumber-1)//26
        return s[::-1]


@dataclass
class Quest():
    inp: int
    sol: str


def checkAns(q: Quest) -> bool:
    sol = Solution()
    out = sol.convertToTitle(q.inp)

    res = (out == q.sol)
    if not res:
        print(f'{q.inp} -> {q.sol} != {out}')

    return res


if __name__ == "__main__":
    q = Quest(1, 'A')
    checkAns(q)

    q = Quest(26, 'Z')
    checkAns(q)

    q = Quest(27, 'AA')
    checkAns(q)

    q = Quest(28, 'AB')
    checkAns(q)

    q = Quest(52, 'AZ')
    checkAns(q)

    q = Quest(78, 'BZ')
    checkAns(q)

    q = Quest(701, 'ZY')
    checkAns(q)

    q = Quest(9999, 'NTO')
    checkAns(q)

    q = Quest(702, 'ZZ')
    checkAns(q)
