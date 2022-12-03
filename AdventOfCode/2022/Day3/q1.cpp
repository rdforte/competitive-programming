#include <bits/stdc++.h>

using namespace std;

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

  string s;
  int total = 0;
  while (getline(cin, s))
  {
    string first = s.substr(0, s.length() / 2);
    string second = s.substr(s.length() / 2, s.length());
    set<char> s2(second.begin(), second.end());

    for (auto c : first)
    {
      auto ch = s2.find(c);
      if (ch != s2.end())
      {
        if ((int)*ch >= 97)
          total += ((int)*ch - 96);
        else
          total += ((int)*ch - 38);
        break;
      }
    }
  }
  cout << total;
}