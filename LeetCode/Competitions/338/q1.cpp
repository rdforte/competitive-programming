#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int kItemsWithMaximumSum(int numOnes, int numZeros, int numNegOnes, int k)
  {
    int res = min(numOnes, k);

    int remainder = k - res;

    if (remainder > 0)
    {
      remainder -= min(numZeros, remainder);
    }

    if (remainder > 0)
    {
      res -= min(numNegOnes, remainder);
    }

    return res;
  }
};

int main()
{
  // Solution
  cout << Solution().kItemsWithMaximumSum(3, 2, 0, 2) << "\n";
  cout << Solution().kItemsWithMaximumSum(3, 2, 0, 4) << "\n";
  cout << Solution().kItemsWithMaximumSum(3, 0, 1, 4) << "\n";
}
