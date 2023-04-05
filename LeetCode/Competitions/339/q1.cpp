#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int findTheLongestBalancedSubstring(string s)
  {
    int longest = 0;
    int zero = 0;
    int one = 0;
    for (int i = 0; i < s.size(); i++)
    {
      if (s[i] == '1')
        one++;

      if (s[i] == '0' || i == s.size() - 1)
      {
        longest = max(min(zero, one) * 2, longest);
        one = 0;
        if (i > 0 && s[i - 1] == '1')
        {
          zero = 1;
        }
        else
          zero++;
      };
    }
    return longest;
  }
};

int main()
{
  cout << Solution().findTheLongestBalancedSubstring("01011");
}
