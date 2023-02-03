// https://leetcode.com/problems/two-sum/description/?envType=study-plan&id=level-1
#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  vector<int> twoSum(vector<int> &nums, int target)
  {
    unordered_map<int, int> numMap;
    for (int i = 0; i < nums.size(); i++)
    {
      if (numMap.count(target - nums[i]) == 1)
      {
        return {numMap[target - nums[i]], i};
      }
      numMap[nums[i]] = i;
    }

    return {};
  }
};

int main()
{
  vector<int> nums{
      2, 7, 11, 15};

  auto sol = Solution().twoSum(nums, 9);

  for (auto s : sol)
  {
    cout << s << ", ";
  }
}