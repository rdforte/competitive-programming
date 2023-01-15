#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int differenceOfSum(vector<int> &nums)
  {
    int absoluteSum = 0;
    int digitSum = 0;
    for (int i = 0; i < nums.size(); i++)
    {
      absoluteSum += nums[i];
      while (nums[i] > 0)
      {
        int num = nums[i] % 10;
        digitSum += num;
        nums[i] /= 10;
      }
    }

    return absoluteSum - digitSum;
  }
};

int main()
{
  // Solution
  vector<int> nums{123};
  cout << Solution().differenceOfSum(nums);
}
