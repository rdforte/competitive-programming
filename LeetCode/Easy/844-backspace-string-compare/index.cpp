#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  bool backspaceCompare(string s, string t)
  {
    string sa;
    string ta;

    for (int i = 0; i < max(s.size(), t.size()); i++)
    {
      if (i < s.size())
        updateString(sa, s[i]);

      if (i < t.size())
        updateString(ta, t[i]);
    }
    return sa == ta;
  }

private:
  void updateString(string &s, char a)
  {
    if (a == '#')
    {
      if (s.size() > 0)
        s.pop_back();
    }
    else
    {
      s.push_back(a);
    }
  }
};

int main()
{
  cout << boolalpha << Solution().backspaceCompare("y#fo##f", "y#f#o##f");
}