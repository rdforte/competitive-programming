// https://cses.fi/problemset/task/1071
#include <bits/stdc++.h>

using namespace std;

#define ll long long

int main()
{
  int t;
  cin >> t;

  while (t--)
  {
    ll y, x;
    cin >> y >> x;

    if (y > x)
    {
      if ((y & 1) == 1)
        cout << ((y - 1) * (y - 1) + x) << "\n";
      else
        cout << ((y * y) - x + 1) << "\n";
    }
    else
    {
      if ((x & 1) == 1)
        cout << x * x - y + 1 << "\n";
      else
        cout << (x - 1) * (x - 1) + y << "\n";
    }
  }
}