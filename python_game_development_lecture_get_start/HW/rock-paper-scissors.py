import random
import sys

GUESSES = ['rock', 'paper', 'scissors']


def get_answer():
    return random.choice(GUESSES)


def get_player_input():
    prompt = """input your select
1. rock
2. paper
3. scissors
q. quit (defult)
"""
    inp = input(prompt)
    return inp


def get_player_answer(inp):
    result = {'1': 'rock', '2': 'paper',
              '3': 'scissors', 'q': 'quit'}.get(inp, 'quit')
    return result


def get_result(ans1, ans2):
    result = {'rock_rock': 'draw', 'paper_paper': 'draw', 'scissors_scissors': 'draw',
              'rock_scissors': 'win', 'scissors_paper': 'win', 'paper_rock': 'win',
              'rock_paper': 'lose', 'paper_scissors': 'lose', 'scissors_rock': 'lose'}
    return result.get(f'{ans1}_{ans2}', 'default')


def launch_game():
    answer = get_answer()
    player_input = get_player_input()
    player_answer = get_player_answer(player_input)
    if player_answer is 'quit':
        sys.exit(1)
    result = get_result(player_answer, answer)
    print(player_answer, answer, result)


if __name__ == '__main__':
    launch_game()
