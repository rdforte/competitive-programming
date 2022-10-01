#include "../../../stdc++.h"

using namespace std;

int lengthOfLongestSubstring(string str)
{
  int longest = 0;
  unordered_set<char> s;

  int a = 0;
  for (int b = 0; b < str.size(); b++)
  {
    if (s.find(str[b]) == s.end())
    {
      s.insert(str[b]);
    }
    else
    {
      while (str[a] != str[b])
      {
        s.erase(s.find(str[a]));
        a++;
      }
      a++;
    }

    if (b - a + 1 > longest)
    {
      longest = b - a + 1;
    }
  }

  return longest;
}

int main()
{
  cout << lengthOfLongestSubstring("abcabcbb");
}