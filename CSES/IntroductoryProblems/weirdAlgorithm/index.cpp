// https://cses.fi/problemset/task/1068
#include <bits/stdc++.h>

using namespace std;

int main()
{
  long long n;
  cin >> n;

  vector<long long> res;

  while (true)
  {
    cout << n << " ";
    if (n == 1)
      break;

    if (n % 2 == 0)
    {
      n /= 2;
    }
    else
    {
      n = (n * 3) + 1;
    }
  }
}