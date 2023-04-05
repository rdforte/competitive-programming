#include "../../../stdc++.h"

using namespace std;

class Solution
{
public:
  vector<vector<int>> findMatrix(vector<int> &nums)
  {
    sort(nums.begin(), nums.end());
    vector<vector<int>> v(1, vector<int>(1, nums[0]));

    int prev = nums[0];
    int duplicate = 0;
    for (int i = 1; i < nums.size(); i++)
    {
      if (nums[i] == prev)
      {
        duplicate++;
        if (v.size() <= duplicate)
        {
          v.push_back({});
        }
      }
      else
      {
        duplicate = 0;
        prev = nums[i];
      }
      v[duplicate].push_back(nums[i]);
    }
    return v;
  }
};

int main()
{
  // Solution
  vector<int> v{1, 3, 4, 1, 2, 3, 1};
  Solution().findMatrix(v);
}