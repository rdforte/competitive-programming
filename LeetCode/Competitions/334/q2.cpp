#include <bits/stdc++.h>

using namespace std;

typedef long long ll;

class Solution
{
public:
  vector<int> divisibilityArray(string word, int m)
  {
    vector<int> res(word.size());

    ll num = 0;
    for (int i = 0; i < word.size(); i++)
    {
      num = (num * 10 + (word[i] - '0')) % m;
      if (num % m == 0)
        res[i] = 1;
      else
        res[i] = 0;
    }

    return res;
  }
};

int main()
{
  // Solution
  auto sol = Solution().divisibilityArray("998244353", 3);
  for (auto s : sol)
  {
    cout << s << ",";
  }
}
