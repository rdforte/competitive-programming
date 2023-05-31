#include "../../../stdc++.h"

using namespace std;

class Solution
{
public:
  int minExtraChar(string s, vector<string> &dictionary)
  {
    unordered_set<string> st;
    for (auto d : dictionary)
    {
      st.insert(d);
    }

    vector<int> dp(s.size() + 1, 0); // make first character empty string

    for (int i = 1; i < dp.size(); i++)
    {
      dp[i] = dp[i - 1] + 1;
      for (int j = i; j > 0; j--)
      {
        int sIndex = j - 1;
        string sub = s.substr(sIndex, i - sIndex);
        if (st.count(sub))
        {
          dp[i] = min(dp[i], dp[sIndex]);
        }
      }
    }

    return dp.back();
  }
};

int main() 
{
  vector<string> dict{"leet", "code", "leetcode"};
  cout << Solution().minExtraChar("leetscode", dict);
}