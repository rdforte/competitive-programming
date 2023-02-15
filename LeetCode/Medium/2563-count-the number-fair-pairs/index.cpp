#include <bits/stdc++.h>

using namespace std;

// O(n^2) time complexity gives TLE
class Solution
{
public:
  long long countFairPairs(vector<int> &nums, int lower, int upper)
  {
    long long p = 0;
    int n = nums.size();
    if (n < 2)
      return p;

    for (int i = 1; i < n; i++)
    {
      for (int j = 0; j < i; j++)
      {
        long long sum = nums[i] + nums[j];
        if (sum >= lower && sum <= upper)
          p++;
      }
    }

    return p;
  }
};

int main()
{
  vector<int> nums{1, 7, 9, 2, 5};
  cout << Solution().countFairPairs(nums, 11, 11);
}