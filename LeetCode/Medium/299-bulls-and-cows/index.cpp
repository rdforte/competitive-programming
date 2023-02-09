#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  string getHint(string secret, string guess)
  {
    int bulls = 0;
    int cows = 0;
    int mx = 1e3 + 5;
    vector<int> s(mx, 0);
    vector<int> g(mx, 0);

    for (int i = 0; i < secret.size(); i++)
    {
      if (secret[i] == guess[i])
      {
        bulls++;
      }
      else
      {
        s[stoi(string(1, secret[i]))]++;
        g[stoi(string(1, guess[i]))]++;
      }
    }

    for (int i = 0; i < mx; i++)
    {
      cows += min(s[i], g[i]);
    }

    return to_string(bulls) + "A" + to_string(cows) + "B";
  }
};

int main()
{
  cout << Solution().getHint("1123", "0111");
}