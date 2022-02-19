#include <iostream>
#include <vector>
#include <string>

class Solution
{
public:
  int search(std::vector<int> &nums, int target)
  {
    std::cout << "new search -----------\n";

    for (int i : nums)
    {
      std::cout << i;
    }

    if (nums.size() == 1)
    {
      if (nums[0] == target)
        return target;
      else
        return -1;
    }

    std::vector<int> left(nums.begin(), nums.begin() + nums.size() / 2);
    std::vector<int> right(nums.begin() + nums.size() / 2, nums.end());

    // for (int i : left)
    // {
    //   std::cout << i;
    // }

    // std::cout << "\n";

    // for (int i : right)
    // {
    //   std::cout << i;
    // }
    // std::cout << "\n";

    std::cout
        << "---------------------------------\n";

    if (target >= right[0] & target <= right[right.size() - 1])
    {
      return search(right, target);
    }
    else if (target <= left[left.size() - 1] & target >= left[0])
    {
      return search(left, target);
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