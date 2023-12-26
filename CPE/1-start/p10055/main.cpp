#include <iostream>

using namespace std;

long long solve(long long, long long);

int main()
{
    long long int num1, num2;
    while (cin >> num1 >> num2)
    {
        cout << solve(num1, num2) << endl;
    }
    return 0;
}

long long solve(long long a, long long b)
{
    return (a > b) ? (a - b) : (b - a);
}
