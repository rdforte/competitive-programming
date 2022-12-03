#include <bits/stdc++.h>

using namespace std;

struct RuckSack
{
  string s;
  unordered_set<char> st;

  RuckSack(string str)
  {
    s = str;
    for (auto ch : s)
    {
      st.insert(ch);
    }
  }
};

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q2output.txt", "w", stdout);

  int total = 0;

  string r1;
  string r2;
  string r3;
  while (getline(cin, r1))
  {
    getline(cin, r2);
    getline(cin, r3);

    auto rs1 = RuckSack(r1);
    auto rs2 = RuckSack(r2);
    auto rs3 = RuckSack(r3);

    vector<RuckSack> rsList{rs1, rs2, rs3};

    sort(rsList.begin(), rsList.end(), [](RuckSack left, RuckSack right)
         { return left.s.size() > right.s.size(); });

    for (auto ch : rsList[0].s)
    {
      if (rsList[1].st.find(ch) != rsList[1].st.end() && rsList[2].st.find(ch) != rsList[2].st.end())
      {
        if ((int)ch >= 97)
          total += ((int)ch - 96);
        else
          total += ((int)ch - 38);
        break;
      }
    }
  }
  cout << total;
}