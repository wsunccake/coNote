#include <iostream>

using namespace std;

void solve(int, int);

int main()
{
    long a, b;

    while (cin >> a >> b && a != 0 && b != 0)
        solve(a, b);

    return 0;
}

void solve(int num1, int num2)
{
    int carry = 0;
    int totalCarry = 0;

    int digit1, digit2;

    while (num1 != 0 || num2 != 0)
    {
        digit1 = num1 % 10;
        digit2 = num2 % 10;

        if (digit1 + digit2 + carry >= 10)
        {
            carry = 1;
            totalCarry++;
        }
        else
        {
            carry = 0;
        }

        num1 /= 10;
        num2 /= 10;
    }

    if (totalCarry == 0)
        cout << "No carry operation." << endl;
    else if (totalCarry == 1)
        cout << "1 carry operation." << endl;
    else
        cout << totalCarry << " carry operations." << endl;
}