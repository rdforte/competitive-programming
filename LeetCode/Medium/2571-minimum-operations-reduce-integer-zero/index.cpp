#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int minOperations(int n)
  {
    int upper = 1;
    int lower;

    while (upper < n)
    {
      lower = upper;
      upper = upper << 1;
    }

    return findOpsDfs(lower, n, upper);
  }

  int findOpsDfs(int lower, int middle, int upper)
  {
    if (lower == middle || middle == upper)
      return 1;

    int count = 0;

    int lm = abs(middle - lower);
    int mu = abs(upper - middle);

    int mid = lm < mu ? lm : mu;

    int l;
    int u = 1;
    while (u < mid)
    {
      l = u;
      u = u << 1;
    }

    count += findOpsDfs(l, mid, u) + 1;

    return count;
  }
};

int main()
{
  cout << Solution().minOperations(54);
}