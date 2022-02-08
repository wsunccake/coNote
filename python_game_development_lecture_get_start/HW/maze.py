MAZE_W = 8
MAZE_H = 6

maze = [[1, 1, 1, 1, 1, 1, 1, 1, ],
        [1, 0, 0, 0, 0, 0, 1, 1, ],
        [1, 0, 1, 1, 1, 0, 0, 1, ],
        [1, 0, 0, 1, 0, 1, 0, 1, ],
        [1, 0, 0, 0, 0, 0, 0, 1, ],
        [1, 1, 1, 1, 1, 1, 1, 1, ]]


def print_maze(maze):
    for y in range(MAZE_H):
        for x in range(MAZE_W):
            print(f'{maze[y][x]} ', end="")
        print()


if __name__ == "__main__":
    print_maze(maze)
