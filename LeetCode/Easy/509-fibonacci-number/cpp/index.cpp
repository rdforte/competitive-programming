// https://leetcode.com/problems/fibonacci-number/
#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int fib(int n)
  {
    if (n == 0)
      return 0;
    if (n == 1)
      return 1;

    int beforePrev = 0;
    int prev = 1;
    for (int i = 2; i <= n; i++)
    {
      int cur = prev + beforePrev;
      beforePrev = prev;
      prev = cur;
    }
    return prev;
  }
};

int main()
{
  cout << Solution().fib(4);
}