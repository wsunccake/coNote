# class Solution:
#     def bestHand(self, ranks: list[int], suits: list[str]) -> str:
#         suit_dict = {}
#         for s in suits:
#             suit_dict[s] = suit_dict.get(s, 0) + 1
#             # if s in suit_dict:
#             #     suit_dict[s] += 1
#             # else:
#             #     suit_dict[s] = 1
#         if len(suit_dict) == 1:
#             return 'Flush'
#         rank_dict = {}
#         for r in ranks:
#             rank_dict[r] = rank_dict.get(r, 0) + 1
#         m = max(rank_dict.values())
#         if m >= 3:
#             return 'Three of a Kind'
#         if m == 2:
#             return 'Pair'
#         return 'High Card'

class Solution:
    def bestHand(self, ranks: list[int], suits: list[str]) -> str:
        suit_dict = {}
        for s in suits:
            suit_dict[s] = suit_dict.get(s, 0) + 1
        if len(suit_dict) == 1:
            return 'Flush'
        rank_dict = {}
        for r in ranks:
            rank_dict[r] = rank_dict.get(r, 0) + 1
        m = max(rank_dict.values())
        result_dict = {
            4: 'Three of a Kind',
            3: 'Three of a Kind',
            2: 'Pair',
        }
        return result_dict.get(m, 'High Card')


# class Solution:
#     def bestHand(self, ranks: list[int], suits: list[str]) -> str:
#         max_rank_cnt = max(Counter(ranks).values())
#         max_suit_cnt = max(Counter(suits).values())
#         return {
#             max_rank_cnt == 2: 'Pair',
#             max_rank_cnt >= 3: 'Three of a Kind',
#             max_suit_cnt == 5: 'Flush',
#         }.get(True, 'High Card')


if __name__ == '__main__':
    sol = Solution()
    ranks = [13, 2, 3, 1, 9]
    suits = ["a", "a", "a", "a", "a"]
    assert sol.bestHand(ranks, suits) == 'Flush', 'Fail'

    ranks = [4, 4, 2, 4, 4]
    suits = ["d", "a", "a", "b", "c"]
    assert sol.bestHand(ranks, suits) == 'Three of a Kind', 'Fail'

    ranks = [10, 10, 2, 12, 9]
    suits = ["a", "b", "c", "a", "d"]
    assert sol.bestHand(ranks, suits) == 'Pair', 'Fail'

    ranks = [2, 10, 7, 10, 7]
    suits = ["a", "b", "a", "d", "b"]
    assert sol.bestHand(ranks, suits) == 'Pair', 'Fail'
