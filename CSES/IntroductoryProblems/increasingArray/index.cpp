// https://cses.fi/problemset/task/1094
#include <bits/stdc++.h>

using namespace std;

typedef long long LL;

int main()
{
  LL moves = 0LL, n, prev, input;
  cin >> n >> prev;

  while (--n > 0)
  {
    cin >> input;
    if (input < prev)
    {
      moves += (prev - input);
    }
    else
    {
      prev = input;
    }
  }
  cout << moves;
}