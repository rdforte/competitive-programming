#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int maxScore(vector<int> &nums)
  {
    long long sum = 0;
    int res = 0;
    sort(nums.begin(), nums.end(), greater<int>());
    for (int i = 0; i < nums.size(); i++)
    {
      sum += nums[i];
      if (sum > 0)
      {
        res++;
      }
    }

    return res;
  }
};

int main()
{
  // Solution
  vector<int> v{2, -1, 0, 1, -3, 3, -3};
  Solution().maxScore(v);
}