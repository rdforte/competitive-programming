#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int vowelStrings(vector<string> &words, int left, int right)
  {
    int res = 0;
    for (int i = left; i <= right; i++)
    {
      if (isVowel(words[i][0]) && isVowel(words[i][words[i].size() - 1]))
      {
        res++;
      }
    }
    return res;
  }

private:
  bool isVowel(char c)
  {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u';
  }
};

int main()
{
  // Solution
}
