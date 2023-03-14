#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int subarraySum(vector<int> &nums, int k)
  {
    int sum = 0, count = 0;
    unordered_map<int, int> mp;
    mp[0] = 1;

    for (int i = 0; i < nums.size(); i++)
    {
      sum += nums[i];
      mp[sum]++;
      if (mp.count(sum - k))
        count += mp[sum - k];
    }

    return count;
  }
};

int main()
{
  vector<int> v{3, 4, 7, 2, -3, 1, 4, 2};
  int k = 7;
  cout << Solution().subarraySum(v, k);
}