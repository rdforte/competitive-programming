#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int distinctIntegers(int n)
  {
    set<int> s;
    s.insert(n);

    while (n > 1)
    {
      int max = 0;
      for (int i = 1; i <= n; i++)
      {
        if (n % i == 1)
        {
          s.insert(i);
          max = i;
        }
      }
      n = max;
    }
    return s.size();
  }
};

int main()
{
  cout << Solution().distinctIntegers(1);
}
