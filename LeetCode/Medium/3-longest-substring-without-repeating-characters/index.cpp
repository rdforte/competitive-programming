#include "../../../stdc++.h"

using namespace std;

int lengthOfLongestSubstring(string str)
{
  int strLen = str.size(), res = 0;
  unordered_map<char, int> mp;

  for (int j = 0, i = 0; j < strLen; j++)
  {
    if (mp.find(str[j]) == mp.end())
    {
      mp[str[j]] = j;
    }
    else
    {
      i = max(mp[str[j]] + 1, i);
      mp[str[j]] = j;
    }

    res = max(res, j - i + 1);
  }

  return res;
}

int main()
{
  cout << lengthOfLongestSubstring("abcabcbb");
}