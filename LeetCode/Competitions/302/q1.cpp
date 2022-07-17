// https://leetcode.com/contest/weekly-contest-302/problems/maximum-number-of-pairs-in-array/
#include <vector>
#include <algorithm>

using namespace std;

class Solution
{
public:
  vector<int> numberOfPairs(vector<int> &nums)
  {
    int pairs = 0;

    sort(nums.begin(), nums.end());

    for (int i = nums.size() - 1; i > 0; i--)
    {
      if (i == 0)
        break;

      if (nums[i] == nums[i - 1])
      {
        nums.pop_back();
        nums.pop_back();
        i--;
        pairs++;
      }
    }

    int size = nums.size();
    return vector<int>{pairs, size};
  }
};