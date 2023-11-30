#include <iostream>
#include <vector>
#include <map>
#include <cassert>

using namespace std;

class Solution
{
public:
    vector<int> twoSum(vector<int> &nums, int target)
    {
        int goal;
        map<int, int> mp;
        vector<int> *results = new vector<int>({0, 0});

        int i = 0;
        for (auto n = nums.begin(); n < nums.end(); n++)
        {
            goal = target - *n;

            if (mp.count(goal))
            {
                results->at(0) = mp[goal];
                results->at(1) = i;
                break;
            }
            else
            {
                mp[*n] = i;
            }

            i++;
        }

        // cout << "index: " << mp[goal]
        //      << " , " << goal
        //      << endl;
        // cout << "index: " << i
        //      << " , " << nums[i]
        //      << endl;

        return *results;
    }
};

int main()
{
    Solution *sol = new Solution();

    vector<int> q;
    int target;
    int ans1, ans2;
    vector<int> a;

    q = {2, 7, 11, 15};
    target = 9;
    ans1 = 0, ans2 = 1;

    a = sol->twoSum(q, target);
    assert(a[0] == ans1);
    assert(a[1] == ans2);

    q = {3, 2, 4};
    target = 6;
    ans1 = 1, ans2 = 2;

    a = sol->twoSum(q, target);
    assert(a[0] == ans1);
    assert(a[1] == ans2);

    q = {3, 3};
    target = 6;
    ans1 = 0, ans2 = 1;

    a = sol->twoSum(q, target);
    assert(a[0] == ans1);
    assert(a[1] == ans2);

    return 0;
}
