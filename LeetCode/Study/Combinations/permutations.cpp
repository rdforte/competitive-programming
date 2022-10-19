#include "../../../stdc++.h"

using namespace std;

class Num
{
public:
  bool seen;
  int number;

  Num(int n, bool s)
  {
    seen = s;
    number = n;
  }
};

void findPermutations(vector<Num *> &nums, vector<int> &perms, vector<vector<int>> &res)
{
  if (perms.size() >= nums.size())
  {
    res.push_back(perms);
    return;
  }

  for (auto num : nums)
  {
    if (num->seen == true)
      continue;
    perms.push_back(num->number);
    num->seen = true;
    findPermutations(nums, perms, res);
    num->seen = false;
    perms.pop_back();
  }
}

int main()
{

  vector<Num *> nums{
      new Num(1, false),
      new Num(2, false),
      new Num(3, false)};

  vector<int> perms;
  vector<vector<int>> res;

  findPermutations(nums, perms, res);

  for (auto r : res)
  {
    cout << "-----------\n";
    for (auto n : r)
    {
      cout << n << ", ";
    }
    cout << "\n";
  }
}