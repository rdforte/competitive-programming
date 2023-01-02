// https://leetcode.com/problems/longest-palindrome/description/
#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int longestPalindrome(string s)
  {
    unordered_map<char, int> m;
    for (int i = 0; i < s.size(); i++)
    {
      m[s[i]]++;
    }

    int count = 0;

    for (auto c : m)
    {
      count += (c.second / 2 * 2);
      if ((count & 1) == 0 and (c.second & 1) == 1)
        count++;
    }

    return count;
  }
};

int main()
{
  cout << Solution().longestPalindrome("ccc");
}