#include <iostream>

using namespace std;

int calculateCycleLength(int);
int solve(int, int);
void printInputAndResult(int, int, int);

int main()
{
    int number1, number2;

    while (cin >> number1 >> number2)
        printInputAndResult(number1, number2, solve(number1, number2));

    return 0;
}

int calculateCycleLength(int n)
{
    int cycleLength = 1;

    while (n != 1)
    {
        if (n % 2 == 0)
            n /= 2;
        else
            n = n * 3 + 1;

        cycleLength++;
    }
    return cycleLength;
}

int solve(int num1, int num2)
{
    int maxNum = num1 > num2 ? num1 : num2;
    int minNum = num1 < num2 ? num1 : num2;

    int maxCycleLength = 0;
    int cycleLength;
    for (int i = minNum; i <= maxNum; i++)
    {
        cycleLength = calculateCycleLength(i);
        if (cycleLength > maxCycleLength)
        {
            maxCycleLength = cycleLength;
        }
    }

    return maxCycleLength;
}

void printInputAndResult(int num1, int num2, int result)
{
    cout << num1 << " " << num2 << " " << result << endl;
}