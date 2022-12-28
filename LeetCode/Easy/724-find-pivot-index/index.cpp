#include <bits/stdc++.h>

using namespace std;

// Brute Force
class SolutionBruteForce
{
public:
  int pivotIndex(vector<int> &nums)
  {
    int pivot;

    for (int i = 0; i < nums.size(); i++)
    {
      int sumLeft = 0;
      int sumRight = 0;
      // calculate left
      for (int left = i - 1; left >= 0; left--)
      {
        sumLeft += nums[left];
      }

      for (int right = i + 1; right < nums.size(); right++)
      {
        sumRight += nums[right];
      }

      if (sumLeft == sumRight)
      {
        return i;
      }
    }

    return -1;
  }
};

class SolutionNotMemoryOptimized
{
public:
  int pivotIndex(vector<int> &nums)
  {
    vector<int> sums(nums.size(), 0);
    for (int i = 0; i < nums.size(); i++)
    {
      int prevSum = i > 0 ? sums[i - 1] : 0;
      sums[i] = (prevSum + nums[i]);
    }

    for (int i = 0; i < sums.size(); i++)
    {
      int sumLeft = i > 0 ? sums[i - 1] : 0;
      int sumRight = sums.back() - sums[i];
      if (sumLeft == sumRight)
      {
        return i;
      }
    }

    return -1;
  }
};

class Solution
{
public:
  int pivotIndex(vector<int> &nums)
  {
    int totalSum = 0;
    for (int i = 0; i < nums.size(); i++)
    {
      totalSum += nums[i];
    }

    int leftSum = 0;
    for (int i = 0; i < nums.size(); i++)
    {
      leftSum += nums[i];
      int sumLeft = leftSum - nums[i];
      int sumRight = totalSum - leftSum;
      if (sumLeft == sumRight)
      {
        return i;
      }
    }

    return -1;
  }
};

int main()
{
  vector<int> nums = {1, 7, 3, 6, 5, 6};
  // vector<int> sums = {1, 8, 11, 17, 22, 28};
  cout << Solution().pivotIndex(nums);
}