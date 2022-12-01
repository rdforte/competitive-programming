#include <bits/stdc++.h>

using namespace std;

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q2Output.txt", "w", stdout);

  string s;
  int currentElf = 0, e1 = 0, e2 = 0, e3 = 0;

  while (getline(cin, s))
  {
    if (s.empty())
    {
      if (currentElf > e1)
      {
        e3 = e2;
        e2 = e1;
        e1 = currentElf;
      }
      else if (currentElf > e2)
      {
        e3 = e2;
        e2 = currentElf;
      }
      else if (currentElf > e3)
      {
        e3 = currentElf;
      }
      currentElf = 0;
    }
    else
    {
      currentElf += stoi(s);
    }
  }

  cout << (e1 + e2 + e3);
}