#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int passThePillow(int n, int time)
  {
    int t = (time % ((n - 1) * 2)) + 1;
    return (t <= n) ? t : n - (t - n);
  }
};

int main()
{
  cout << Solution().passThePillow(8, 9);
}