
maze = [[1, 1, 1, 1, 1, 1, 1, 1, ],
        [1, 8, 0, 0, 0, 0, 1, 1, ],
        [1, 0, 1, 1, 1, 0, 0, 1, ],
        [1, 0, 0, 1, 0, 1, 0, 1, ],
        [1, 0, 0, 1, 0, 0, 0, 1, ],
        [1, 0, 0, 0, 0, 0, 9, 1, ],
        [1, 1, 1, 1, 1, 1, 1, 1, ]]


class Location:
    X = 1
    Y = 1

    def __init__(self, x, y) -> None:
        self.X = x
        self.Y = y


def load_maze(maze_file):
    maze = []
    with open(maze_file, 'r', encoding='utf-8') as f:
        for line in f:
            maze.append([int(x) for x in line.split()])
    return maze


def print_maze(maze):
    heigh = len(maze)
    width = len(maze[0])
    for y in range(heigh):
        for x in range(width):
            print(f'{maze[y][x]} ', end="")
        print()


if __name__ == "__main__":
    # maze = load_maze('maze.txt')
    # print_maze(maze)
    position = Location(1, 1)
    # inp = input()
    # print(inp)
    while maze[position.Y][position.X] != 9:
        print_maze(maze)
        inp = input()

        new_Y = position.Y
        new_X = position.X

        # '\x1b[A' up
        # '\x1b[B' down
        # '\x1b[C' right
        # '\x1b[D' left

        if inp == '\x1b[A':
            new_Y = new_Y - 1
            if maze[new_Y][new_X] == 9:
                break
            if maze[new_Y][new_X] != 1:
                maze[position.Y][position.X] = 0
                maze[new_Y][new_X] = 8
                position.Y = new_Y

        if inp == '\x1b[B':
            new_Y = new_Y + 1
            if maze[new_Y][new_X] == 9:
                break
            if maze[new_Y][new_X] != 1:
                maze[position.Y][position.X] = 0
                maze[new_Y][new_X] = 8
                position.Y = new_Y

        if inp == '\x1b[C':
            new_X = new_X + 1
            if maze[new_Y][new_X] == 9:
                break
            if maze[new_Y][new_X] != 1:
                maze[position.Y][position.X] = 0
                maze[new_Y][new_X] = 8
                position.X = new_X

        if inp == '\x1b[D':
            new_X = new_X - 1
            if maze[new_Y][new_X] == 9:
                break
            if maze[new_Y][new_X] != 1:
                maze[position.Y][position.X] = 0
                maze[new_Y][new_X] = 8
                position.X = new_X

    print('congratulation, exit maze')
