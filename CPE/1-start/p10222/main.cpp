#include <iostream>
#include <string>

using namespace std;

string ENCODE_TEXT = " 234567890-=ertyuiop[]\\dfghjkl;'cvbnm,./ERTYUIOPDFGHJKLCVBNM";
string DECODE_TEXT = " `1234567890qwertyuiop[asdfghjklzxcvbnm,qwertyuiasdfghjzxcvb";

string solve(string);

int main()
{
    string s, w;

    while (getline(cin, s))
        cout << solve(s) << endl;

    return 0;
}

string solve(string line)
{
    string word = "";
    int size = line.size();

    while (size)
        word = DECODE_TEXT[ENCODE_TEXT.find(line[--size])] + word;

    return word;
}
