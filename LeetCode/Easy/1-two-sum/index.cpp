#include "../../../stdc++.h"

using namespace std;

vector<int> twoSum(vector<int> &nums, int target)
{
  unordered_map<int, int> mpp;

  for (int i = 0; i < nums.size(); i++)
  {
    int secondNum = target - nums[i];
    if (mpp.find(secondNum) != mpp.end())
      return {mpp[secondNum], i};

    mpp[nums[i]] = i;
  }

  return {};
}