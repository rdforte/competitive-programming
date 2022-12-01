#include <bits/stdc++.h>

using namespace std;

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1Output.txt", "w", stdout);

  string s;
  int elfCal = 0;
  int maxCal = 0;
  while (getline(cin, s))
  {
    if (s.empty())
    {
      maxCal = max(elfCal, maxCal);
      elfCal = 0;
    }
    else
    {
      elfCal += stoi(s);
    }
  }

  cout << maxCal;
}