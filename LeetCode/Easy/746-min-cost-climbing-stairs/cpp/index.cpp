#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
    int minCostClimbingStairs(vector<int> &cost)
    {
        int stepOne = cost[0];
        int stepTwo = cost[1];

        for (int i = 2; i < cost.size(); i++)
        {
            int cur = min(stepOne, stepTwo) + cost[i];
            stepOne = stepTwo;
            stepTwo = cur;
        }

        return min(stepOne, stepTwo);
    }
};

int main()
{
    vector<int> cost{1, 100, 1, 1, 1, 100, 1, 1, 100, 1};
    cout << Solution().minCostClimbingStairs(cost);
}