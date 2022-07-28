// https://leetcode.com/problems/binary-search/
#include <iostream>
#include <vector>
#include <string>

class Solution
{
public:
  int search(std::vector<int> &nums, int target)
  {
    return recursiveSearch(nums, target, 0, nums.size() - 1);
  }

private:
  // Using recursion I was able to move through this like a b-tree halving the search every time.
  // Runtime O(log n)
  // Spacetime O(1) since all I have to keep track of is the index's for the upper/lower bounds.
  int recursiveSearch(std::vector<int> &nums, int target, int lower, int upper)
  {

    // if lower is same as upper then it is just 1 number remaining
    if (lower == upper)
    {
      if (nums[lower] == target)
        return lower;
      else
        return -1;
    }

    int middle = ((upper - lower) / 2) + lower;

    if (target <= nums[middle] & target >= nums[lower])
    {
      return recursiveSearch(nums, target, lower, middle);
    }
    else if (target >= nums[middle + 1] & target <= nums[upper])
    {
      return recursiveSearch(nums, target, middle + 1, upper);
    }

    return -1;
  }
};

int main()
{
  std::vector<int> nums{-1, 0, 3, 5, 9, 12};

  int sol = Solution().search(nums, 9);

  std::cout << "\nsolution: " << sol;
}