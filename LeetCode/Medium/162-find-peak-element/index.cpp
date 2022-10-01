#include "../../../stdc++.h"

using namespace std;

int findPeakElement(vector<int> &nums)
{
  return getPeak(nums, 0, nums.size() - 1);
}

int getPeak(vector<int> &nums, int lowerBound, int upperBound)
{
  if (lowerBound == upperBound)
    return lowerBound;

  if ((upperBound - lowerBound) == 1)
  {
    if (nums[upperBound] > nums[lowerBound])
      return upperBound;

    return lowerBound;
  }

  int middle = lowerBound + ((upperBound - lowerBound) / 2);

  if (nums[middle] > nums[middle - 1] && nums[middle] > nums[middle + 1])
    return middle;

  if (nums[middle + 1] > nums[middle])
    return getPeak(nums, middle + 1, upperBound);

  return getPeak(nums, lowerBound, upperBound - 1);
}
