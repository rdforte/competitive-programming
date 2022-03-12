#include <iostream>
#include <vector>

class Solution
{
public:
  // Recursive approach
  // Spacetime O(1) we just have to keep track of the middle index
  // Runtime O(log.n) we halve the array every single time.
  // This could be done with a while loop
  int peakIndexInMountainArray(std::vector<int> &arr)
  {
    return peakRecursion(arr, 0, arr.size() - 1);
  }
  int peakRecursion(std::vector<int> &arr, int lower, int upper)
  {
    if (upper == lower)
      return upper;

    int middle = lower + ((upper - lower) / 2);

    if (arr[middle] < arr[middle + 1])
    {
      return peakRecursion(arr, middle + 1, upper);
    }
    else
    {
      return peakRecursion(arr, lower, middle);
    }
  }
};

int main()
{
  std::vector<int> mount{0, 1, 2, 3, 4, 5};

  int res = Solution().peakIndexInMountainArray(mount);
  std::cout << res;
}