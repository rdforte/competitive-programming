// https://leetcode.com/problems/find-pivot-index/
#include <vector>
#include <iostream>

using namespace std;

class Solution
{
public:
  int pivotIndex(vector<int> &nums)
  {
    int sum = 0;
    int leftSum = 0;
    for (auto n : nums)
      sum += n;

    for (int i = 0; i < nums.size(); i++)
    {
      if (sum - leftSum - nums[i] == leftSum)
      {
        return i;
      }

      leftSum += nums[i];
    }

    return -1;
  }
};

int main()
{
  // 3
  vector<int> nums{1, 7, 3, 6, 5, 6};

  // -1
  // vector<int> nums{1, 2, 3};

  // 0
  // vector<int> nums{2, 1, -1};

  cout << Solution().pivotIndex(nums);
}