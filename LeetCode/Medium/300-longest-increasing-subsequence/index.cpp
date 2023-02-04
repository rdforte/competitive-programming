// https://leetcode.com/problems/longest-increasing-subsequence/description/
#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int lengthOfLIS(vector<int> &nums)
  {
    vector<int> l(nums.size(), 1);
    int res = 1;
    for (int i = 1; i < nums.size(); i++)
    {
      for (int j = 0; j < i; j++)
      {
        if (nums[j] < nums[i])
        {
          l[i] = max(l[j] + 1, l[i]);
        }
      }
      res = max(l[i], res);
    }
    return res;
  }
};

int main()
{
  vector<int> nums{10, 9, 2, 5, 3, 7, 101, 18};
  cout << Solution().lengthOfLIS(nums);
}