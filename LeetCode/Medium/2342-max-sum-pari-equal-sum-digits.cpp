// https://leetcode.com/contest/weekly-contest-302/problems/max-sum-of-a-pair-with-equal-sum-of-digits/
#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>
#include <unordered_map>

using namespace std;

int maximumSum(vector<int> &nums)
{
  int res = -1, d_n[82] = {}; // 9 * 9
  for (int n : nums)
  {
    int d = 0;
    for (int nn = n; nn; nn /= 10)
      d += nn % 10;
    if (d_n[d])
      res = max(res, d_n[d] + n);
    d_n[d] = max(d_n[d], n);
  }
  return res;
}

int main()
{
  vector<int> nums{
      18, 43, 36, 13, 7};

  cout << maximumSum(nums);
}