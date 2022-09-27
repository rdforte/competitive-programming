// https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/
#include "../../../stdc++.h"

using namespace std;

// because we always take the smallest number and subtract one this means at
// each iteration we remove all numbers which are the same.
// The only numbers which we do not do any arithmetic on are the 0's so we can
// subtract the 0 at the end which will give us the total number of operations.

class Solution
{
public:
  int minimumOperations(vector<int> &nums)
  {
    unordered_set<int> s(nums.begin(), nums.end());
    return s.size() - s.count(0);
  }
};

int main()
{
  vector<int> nums{1};

  cout << Solution().minimumOperations(nums);
}