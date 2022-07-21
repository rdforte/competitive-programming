// https://leetcode.com/explore/learn/card/array-and-string/201/introduction-to-array/1147/
#include <iostream>
#include <vector>

using namespace std;

class Solution
{
public:
  int dominantIndex(vector<int> &nums)
  {
    int index = 0;
    int largest = 0;
    int secondLargest = 0;

    for (int i = 0; i < nums.size(); i++)
    {
      if (nums[i] > largest)
      {
        secondLargest = largest;
        largest = nums[i];
        index = i;
      }
      else if (nums[i] > secondLargest)
      {
        secondLargest = nums[i];
      }
    }

    if (secondLargest == 0)
      return index;

    return (largest / secondLargest) >= 2 ? index : -1;
  }
};

int main()
{
  // 1
  // vector<int> nums{3, 6, 1, 0};

  // 3
  // vector<int> nums{0, 0, 0, 1};

  // 0
  vector<int> nums{1, 0};

  cout << Solution().dominantIndex(nums);
}