#include "../../../stdc++.h"

using namespace std;

int climbStairs(int n)
{
  vector<int> s(n + 1);
  s[0] = 1;

  for (int i = 0; i < n; i++)
  {
    if (i + 1 <= n)
      s[i + 1] += s[i];

    if (i + 2 <= n)
      s[i + 2] += s[i];
  }

  return s[s.size() - 1];
}

int main()
{
  cout << climbStairs(3);
}