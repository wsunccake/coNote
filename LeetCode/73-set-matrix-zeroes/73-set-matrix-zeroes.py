from typing import List


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        m = len(matrix)
        n = len(matrix[0])

        zero_point_list = []
        for i in range(m):
            for j in range(n):
                if matrix[i][j] == 0:
                    zero_point_list.append((i, j))

        for p in zero_point_list:
            x, y = p[0], p[1]
            for j in range(n):
                matrix[x][j] = 0
            for i in range(m):
                matrix[i][y] = 0


if __name__ == '__main__':
    sol = Solution()
    inputs = [
        [[1, 1, 1], [1, 0, 1], [1, 1, 1]],
        [[0, 1, 2, 0], [3, 4, 5, 2], [1, 3, 1, 5]]
    ]
    outputs = [
        [[1, 0, 1], [0, 0, 0], [1, 0, 1]],
        [[0, 0, 0, 0], [0, 4, 5, 0], [0, 3, 1, 0]]
    ]

    for pos in range(len(outputs)):
        sol.setZeroes(inputs[pos])
        if inputs[pos] != outputs[pos]:
            print(f'{inputs[pos]}, {outputs[pos]}')
