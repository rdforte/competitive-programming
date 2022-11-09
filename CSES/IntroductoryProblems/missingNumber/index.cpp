// https://cses.fi/problemset/task/1083
#include <bits/stdc++.h>

using namespace std;

int main()
{
  int n;
  cin >> n;

  int missing = n;

  for (int i = 1; i < n; i++)
  {
    int input;
    cin >> input;
    missing ^= input ^ i;
  }

  cout << missing;
}