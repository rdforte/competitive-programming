#include <vector>
#include <iostream>

using namespace std;

class Solution
{
public:
  int findDuplicate(vector<int> &nums)
  {
    while (nums[0] != nums[nums[0]])
    {
      swap(nums[0], nums[nums[0]]);
    }
    return nums[0];
  }
};

int main()
{
  vector<int> nums = {1, 3, 5, 4, 4};

  cout << Solution().findDuplicate(nums); // 4
}