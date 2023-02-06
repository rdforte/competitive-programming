// https://leetcode.com/problems/count-vowel-strings-in-ranges/description/
#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  vector<int> vowelStrings(vector<string> &words, vector<vector<int>> &queries)
  {
    vector<int> r(words.size(), 0);
    for (int i = 0; i < words.size(); i++)
    {
      int v = hasVowelStartEnd(words[i]) ? 1 : 0;
      if (i == 0)
      {
        r[i] = v;
        continue;
      }

      r[i] = r[i - 1];
      r[i] += v;
    }

    vector<int> res(queries.size());
    for (int i = 0; i < queries.size(); i++)
    {
      int s = queries[i][0] == 0 ? 0 : r[queries[i][0] - 1];
      res[i] = r[queries[i][1]] - s;
    }

    return res;
  }

private:
  bool hasVowelStartEnd(string s)
  {
    char l = s.size() - 1;
    return (s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u') &&
           (s[l] == 'a' || s[l] == 'e' || s[l] == 'i' || s[l] == 'o' || s[l] == 'u');
  }
};