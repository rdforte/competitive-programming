#include <iostream>
#include <vector>

class Solution
{
public:
    // time complexity O(n) we loop through the whole array
    // space complexity O(1) we keep track of 2 variables
    int maxProfit(std::vector<int> &prices)
    {
        int l = INT_MAX; // the lowest price we start at is the max and work backwards
        int p = 0;       // my current profit is 0

        for (int i = 1; i < prices.size(); i++)
        {
            l = prices[i - 1] < l ? prices[i - 1] : l;
            if ((prices[i] - l) > p)
            {
                p = prices[i] - l;
            }
        }
        return p;
    }
};

int main()
{
    std::vector<int> v{7, 6, 4, 3, 1};
    std::cout << Solution().maxProfit(v);
}