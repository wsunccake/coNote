#include <iostream>
#include <string>

using namespace std;

string solve(string, string);

int main()
{
    string s1, s2;
    while (cin >> s1 >> s2)
        cout << solve(s1, s2) << endl;

    return 0;
}

string solve(string s1, string s2)
{
    int ascii1[26] = {0};
    int ascii2[26] = {0};

    string commonChars = "";
    int i = 0;

    for (i = 0; i < s1.size(); i++)
        ascii1[s1[i] - 'a'] += 1;

    for (i = 0; i < s2.size(); i++)
        ascii2[s2[i] - 'a'] += 1;

    for (i = 0; i < 26; i++)
    {
        if (ascii1[i] > 0 && ascii2[i] > 0)
            commonChars += static_cast<char>(i + 'a');
    }

    return commonChars;
}
