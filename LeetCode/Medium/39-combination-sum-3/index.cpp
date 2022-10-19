#include "../../../stdc++.h"

using namespace std;

class Num
{
public:
    int number;
    bool seen;

    Num(int n, bool s)
    {
        number = n;
        seen = s;
    }
};

void dfsCombination(vector<Num *> &nums, vector<vector<int>> &combinations, vector<int> &combo, int k, int sum, int n, int start)
{
    if (combo.size() >= k || sum > n)
    {
        if (sum == n)
            combinations.push_back(combo);
        return;
    }

    for (int i = start; i < 9; i++)
    {
        if (nums[i]->seen == true)
            continue;
        combo.push_back(nums[i]->number);
        nums[i]->seen = true;
        dfsCombination(nums, combinations, combo, k, (sum + nums[i]->number), n, (i + 1));
        nums[i]->seen = false;
        combo.pop_back();
    }
}

vector<vector<int>> combinationSum3(int k, int n)
{
    vector<Num *> nums;
    for (int i = 1; i <= 9; i++)
    {
        nums.push_back(new Num(i, false));
    }

    vector<vector<int>> combinations;
    vector<int> combo;

    dfsCombination(nums, combinations, combo, k, 0, n, 0);

    return combinations;
};

int main()
{
    int k = 3, n = 7;
    vector<vector<int>> res = combinationSum3(k, n);

    for (auto r : res)
    {
        cout << "----------\n";
        for (auto i : r)
        {
            cout << i << ", ";
        }
        cout << "\n";
    }
}