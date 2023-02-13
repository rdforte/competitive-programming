#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  long long findTheArrayConcVal(vector<int> &nums)
  {
    long long res = 0;
    for (int i = 0, j = nums.size() - 1; i <= j; i++, j--)
    {
      if (i == j)
        res += nums[j];
      else
        res += stoll(to_string(nums[i]) + to_string(nums[j]));
    }
    return res;
  }
};

int main()
{
  vector<int> nums{5, 14, 13, 8, 12};
  cout << Solution().findTheArrayConcVal(nums);
}