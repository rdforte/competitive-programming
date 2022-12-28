#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  bool isIsomorphic(string s, string t)
  {
    unordered_map<char, char> mapping_s_t;
    unordered_map<char, char> mapping_t_s;
    for (int i = 0; i < s.size(); i++)
    {
      if (mapping_s_t[s[i]] or mapping_t_s[t[i]])
      {
        if (mapping_s_t[s[i]] == t[i] and mapping_t_s[t[i]] == s[i])
        {
          continue;
        }
        else
        {
          return false;
        }
      }
      else
      {
        mapping_s_t[s[i]] = t[i];
        mapping_t_s[t[i]] = s[i];
      }
    }
    return true;
  }
};

int main()
{
  cout << boolalpha << Solution().isIsomorphic("babc", "baba");
}