#include "../../../stdc++.h"

using namespace std;

class Solution
{
public:
  string makeSmallestPalindrome(string s)
  {
    for (int l = 0, r = s.size() - 1; l <= r; l++, r--)
    {
      s[l] = s[l] - 'a' < s[r] - 'a' ? s[l] : s[r];
      s[r] = s[r] - 'a' < s[l] - 'a' ? s[r] : s[l];
    }
    return s;
  }
};

int main()
{
  cout << Solution().makeSmallestPalindrome("egcfe") << "\n";
  cout << Solution().makeSmallestPalindrome("abcd") << "\n";
  cout << Solution().makeSmallestPalindrome("seven") << "\n";
}
