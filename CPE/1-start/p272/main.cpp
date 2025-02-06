#include <iostream>
#include <string>

using namespace std;

string SYMBOL[2] = {"``", "''"};
int RL_TOGGLE = 0;

string solve(string);

int main()
{
    string input;
    while (getline(cin, input))
        cout << solve(input) << endl;

    return 0;
}

string solve(string input)
{
    string output = "";

    for (int i = 0; i < input.size(); i++)
    {
        if (input[i] == '"')
        {
            output += SYMBOL[RL_TOGGLE];
            RL_TOGGLE ^= 1;
        }
        else
        {
            output += input[i];
        }
    }

    return output;
}
