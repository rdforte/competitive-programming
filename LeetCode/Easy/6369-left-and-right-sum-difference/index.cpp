#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  vector<int> leftRigthDifference(vector<int> &nums)
  {
    vector<int> res(nums.size());

    for (int i = 0; i < nums.size(); i++)
    {
      int left = i - 1, right = i + 1, sumLeft = 0, sumRight = 0;
      while (left >= 0 || right < nums.size())
      {
        if (left >= 0)
          sumLeft += nums[left];

        if (right < nums.size())
          sumRight += nums[right];

        left--;
        right++;
      }
      res[i] = abs(sumLeft - sumRight);
    }

    return res;
  }
};

int main()
{
  // Solution
  vector<int> n = {10, 4, 8, 3};
  auto sol = Solution().leftRigthDifference(n);
  for (auto s : sol)
  {
    cout << s << ", ";
  }
}