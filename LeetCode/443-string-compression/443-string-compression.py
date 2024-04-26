from typing import List


class Solution:
    def compress(self, chars: List[str]) -> int:
        char_list = [chars[0]]
        num_list = [1]
        for i in range(len(chars)-1):
            if chars[i] == chars[i+1]:
                num_list[len(char_list)-1] += 1
            else:
                char_list.append(chars[i+1])
                num_list.append(1)

        tmp_str = ""
        for i in range(len(char_list)):
            tmp_str += char_list[i]
            if num_list[i] != 1:
                tmp_str += f"{num_list[i]}"
        pos = 0
        for c in tmp_str:
            chars[pos] = c
            pos += 1

        return pos

    # # slow
    # def compress(self, chars: List[str]) -> int:
    #     result_str = f"{chars[0]}"
    #     count = 1

    #     for i in range(len(chars)-1):
    #         if chars[i] == chars[i+1]:
    #             count += 1
    #         else:
    #             if count != 1:
    #                 result_str = f"{result_str}{count}"
    #             result_str = f"{result_str}{chars[i+1]}"
    #             count = 1

    #     if count != 1:
    #         result_str = f"{result_str}{count}"

    #     pos = 0
    #     for c in result_str:
    #         chars[pos] = c
    #         pos += 1

    #     return pos


if __name__ == '__main__':
    sol = Solution()
    inputs = [["a", "a", "b", "b", "c", "c", "c"], ["a"], [
        "a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"],
        ["b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "a", "b"]]
    outputs = [(["a", "2", "b", "2", "c", "3"], 6), (["a"], 1),
               (["a", "b", "1", "2"], 4), (["b", "1", "1", "a", "b"], 5)]

    for pos in range(len(outputs)):
        ans = sol.compress(inputs[pos])
        if inputs[pos][:ans] != outputs[pos][0]:
            print(f'{inputs[pos]}, {ans}, {outputs[pos]}')
