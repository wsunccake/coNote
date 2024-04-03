#include <iostream>
#include <map>
#include <string>

using namespace std;

bool isUnique(string str)
{
    map<int, bool> seen;
    for (size_t i = 0; i < str.length(); i++)
    {
        // cout << str[i] << endl;
        if (seen.find(str[i]) == seen.end())
        {
            seen.insert(pair<int, bool>(str[i], true));
        }
        else
        {
            return false;
        }
    }

    return true;
}

int main()
{
    string inputs[5] = {"abcde", "hello", "apple", "kite", "padle"};
    bool outputs[5] = {true, false, false, true, true};
    for (int i = 0; i < 5; i++)
    {
        if (isUnique(inputs[i]) != outputs[i])
        {
            cout << inputs[i]
                 << "," << (isUnique(inputs[i]) ? "true" : "false")
                 << "," << (outputs[i] ? "true" : "false")
                 << endl;
        }
    }

    return 0;
}