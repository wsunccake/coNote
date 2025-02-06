#include <iostream>

using namespace std;

int MAX_COUNT = 0;
int COUNT[26] = {0};

void solve(string);

int main()
{
    int n, i, j;
    string s;

    cin >> n;

    while (--n && cin.get() == '\n')
    {
        while (cin >> s)
            solve(s);
    }

    for (i = MAX_COUNT; i > 0; i--)
    {
        for (j = 0; j < 26; j++)
        {
            if (i == COUNT[j])
                cout << static_cast<char>(j + 'A') << " " << COUNT[j] << endl;
        }
    }

    return 0;
}

void solve(string s)
{
    for (int i = 0; i < s.size(); i++)
    {
        char c = s[i];
        if (c >= 'A' && c <= 'Z')
        {
            COUNT[c - 'A'] += 1;
            MAX_COUNT = max(MAX_COUNT, COUNT[c - 'A']);
        }
        else if (c >= 'a' && c <= 'z')
        {
            COUNT[c - 'a'] += 1;
            MAX_COUNT = max(MAX_COUNT, COUNT[c - 'a']);
        }
    }
}
