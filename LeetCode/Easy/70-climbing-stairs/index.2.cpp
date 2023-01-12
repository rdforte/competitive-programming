#include <bits/stdc++.h>

using namespace std;

class SolutionLeastOptimal
{
public:
  int climbStairs(int n)
  {
    vector<int> steps(n + 1, 0);
    steps[0] = 1;
    for (int i = 0; i < n; i++)
    {
      if (i + 2 <= n)
        steps[i + 2] += steps[i];

      if (i + 1 <= n)
        steps[i + 1] += steps[i];
    }
    return steps[n];
  }
};

class Solution
{
public:
  int climbStairs(int n)
  {
    if (n == 1)
      return 1;
    int firstStep = 1;
    int secondStep = 2;
    for (int i = 3; i <= n; i++)
    {
      int curStep = firstStep + secondStep;
      firstStep = secondStep;
      secondStep = curStep;
    }
    return secondStep;
  }
};

int main()
{
  cout << Solution().climbStairs(4);
}