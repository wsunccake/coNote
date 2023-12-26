#include <iostream>
#include <string>

using namespace std;

int main()
{
    string lines[100];
    int line = 0, maxLen = 0, l;

    while (getline(cin, lines[line]) && lines[line] != "")
    {
        l = lines[line].size();
        maxLen = (maxLen > l) ? maxLen : l;
        line++;
    }

    for (int i = 0; i < maxLen; i++)
    {
        for (int j = line - 1; j >= 0; j--)
        {
            if (lines[j].size() > i)
                cout << lines[j][i];
            else
                cout << " ";
        }
        cout << endl;
    }

    return 0;
}
