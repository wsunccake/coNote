#include <iostream>
#include <algorithm>

using namespace std;

int solve(int[], int);

int main()
{
    int arr[500];
    int testcase_number, street_number, i;

    cin >> testcase_number;
    while (--testcase_number)
    {
        cin >> street_number;
        for (i = 0; i < street_number; i++)
            cin >> arr[i];

        cout << solve(arr, street_number) << endl;
    }
}

int solve(int arr[], int size)
{
    int middle_index = size / 2;
    int total = 0;

    int i;
    sort(arr, arr + size);
    for (i = 0; i < size; i++)
    {
        total += abs(arr[i] - arr[middle_index]);
    }

    return total;
}
