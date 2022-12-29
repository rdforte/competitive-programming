#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  bool isSubsequence(string s, string t)
  {
    string word = "";
    for (int sp = 0, tp = 0; tp < t.size();)
    {
      if (sp < s.size() and s[sp] == t[tp])
      {
        word += t[tp];
        sp++;
      }
      tp++;
    }

    return word == s;
  }
};

int main()
{
  cout << boolalpha << Solution().isSubsequence("abd", "ahbgdc");
}