from typing import List


class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        n = len(matrix)
        raw_dict = {}

        for i in range(n):
            for j in range(n):
                raw_dict[f'{i}_{j}'] = matrix[i][j]

        for i in range(n):
            for j in range(n):
                matrix[i][j] = raw_dict[f'{n-j-1}_{i}']


if __name__ == '__main__':
    sol = Solution()
    inputs = [
        [[1, 2, 3], [4, 5, 6], [7, 8, 9]],
        [[5, 1, 9, 11], [2, 4, 8, 10], [13, 3, 6, 7], [15, 14, 12, 16]]
    ]
    outputs = [
        [[7, 4, 1], [8, 5, 2], [9, 6, 3]],
        [[15, 13, 2, 5], [14, 3, 4, 1], [12, 6, 8, 9], [16, 7, 10, 11]]
    ]

    for pos in range(len(outputs)):
        sol.rotate(inputs[pos])
        if inputs[pos] != outputs[pos]:
            print(f'{inputs[pos]}, {outputs[pos]}')
