#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int search(vector<int> &nums, int target)
  {
    for (int i = 0, j = nums.size() - 1; i <= j;)
    {
      int middle = (j + i) / 2; // same as ((j - i) / 2) + i
      if (nums[middle] == target)
        return middle;

      if (target > nums[middle])
        i = middle + 1;
      else
        j = middle - 1;
    }
    return -1;
  }
};