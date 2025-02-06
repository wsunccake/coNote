#include <iostream>

using namespace std;

int solve(long);

int main()
{
    long number;
    while (cin >> number && number != 0)
        cout << solve(number) << endl;

    return 0;
}

int solve(long number)
{
    while (number >= 10)
        number = number / 10 + number % 10;

    return number;
}
