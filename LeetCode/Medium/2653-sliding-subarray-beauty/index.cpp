#include "../../../stdc++.h"

using namespace std;

class Solution
{
public:
  vector<int> getSubarrayBeauty(vector<int> &nums, int k, int x)
  {
    vector<int> res(nums.size() - k + 1);
    vector<int> v(50, 0);

    for (int l = 0, r = 0; l < res.size(); r++)
    {
      if (nums[r] < 0)
        v[nums[r] + 50]++;
      if (r + 1 - l == k)
      {
        // process sub array
        int pos = x;
        for (int n = 0; n < v.size(); n++)
        {
          if (v[n] == 0)
            continue;
          pos -= v[n];
          if (pos <= 0)
          {
            res[l] = n - 50;
            break;
          }
        }

        if (nums[l] < 0)
          v[nums[l] + 50]--;

        l++;
      }
    }
    return res;
  }
};

void test1()
{
  vector<int> nums = {1, -1, -3, -2, 3};
  int k = 3, x = 2;

  auto sol = Solution().getSubarrayBeauty(nums, k, x);
  for (auto s : sol)
  {
    cout << s << ",";
  }
}

void test2()
{
  vector<int> nums = {-1, -2, -3, -4, -5};
  int k = 2, x = 2;

  auto sol = Solution().getSubarrayBeauty(nums, k, x);
  for (auto s : sol)
  {
    cout << s << ",";
  }
}

void test3()
{
  vector<int> nums = {-3, 1, 2, -3, 0, -3};
  int k = 2, x = 1;

  auto sol = Solution().getSubarrayBeauty(nums, k, x);
  for (auto s : sol)
  {
    cout << s << ",";
  }
}

int main()
{
  test1();
  cout << "\n";
  test2();
  cout << "\n";
  test3();
}
