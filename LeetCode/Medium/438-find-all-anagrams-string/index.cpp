// https://leetcode.com/problems/find-all-anagrams-in-a-string/?envType=study-plan&id=level-1
#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  vector<int> findAnagrams(string s, string p)
  {
    vector<int> sol;
    if (s.size() < p.size())
      return {};

    unordered_map<char, int> pCharacters;
    unordered_map<char, int> sCharacters;

    for (int i = 0; i < p.size(); i++)
    {
      pCharacters[p[i]]++;
      sCharacters[s[i]]++;
    }

    for (int i = 0, j = p.size() - 1; j < s.size();)
    {
      if (pCharacters == sCharacters)
      {
        sol.push_back(i);
      }

      j++;
      sCharacters[s[j]]++;
      sCharacters[s[i]]--;
      if (sCharacters[s[i]] == 0)
        sCharacters.erase(s[i]);
      i++;
    }

    return sol;
  }
};

int main()
{
  auto sol = Solution().findAnagrams("abacbabc", "abc");
  for (auto s : sol)
  {
    cout << s << ", ";
  }
}