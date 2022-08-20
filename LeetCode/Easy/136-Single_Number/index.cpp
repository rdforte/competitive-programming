#include <vector>
#include <iostream>

using namespace std;

class Solution
{
public:
  int singleNumber(vector<int> &nums)
  {
    int singleNum = nums[0];

    for (int i = 1; i < nums.size(); i++)
      singleNum ^= nums[i];

    return singleNum;
  }
};

int main()
{
  vector<int> nums = {4, 1, 2, 1, 2};

  cout << Solution().singleNumber(nums);
}