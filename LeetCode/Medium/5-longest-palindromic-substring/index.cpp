#include <bits/stdc++.h>

using namespace std;

int palFromMiddle(int left, int right, string s)
{
  int len = 0;
  while (left >= 0 && right < s.size() && (s[left] == s[right]))
  {
    len = right - left + 1;
    left--;
    right++;
  }
  return len;
}

string longestPalindrome(string s)
{
  if (s.size() == 1)
    return s;

  int start = 0, end = 0;

  for (int i = 0; i < s.size(); i++)
  {
    int pOdd = palFromMiddle(i, i, s);
    int pEven = palFromMiddle(i, i + 1, s);
    int maxLen = max(pOdd, pEven);
    if (maxLen > (end - start))
    {
      start = i - ((maxLen - 1) / 2);
      end = i + ((maxLen) / 2);
    }
  }

  return s.substr(start, (end - start + 1));
};

int main()
{
  string s;
  cin >> s;

  cout << longestPalindrome(s) << "\n";
}