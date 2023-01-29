#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int alternateDigitSum(int n)
  {
    bool isPositive = true;
    int res = 0;
    for (auto c : to_string(n))
    {
      int num = stoi(""s + c);
      isPositive ? (res += num) : (res -= num);
      isPositive = !isPositive;
    }
    return res;
  };
};

int main()
{
  cout << Solution().alternateDigitSum(521);
  // Solution
}
