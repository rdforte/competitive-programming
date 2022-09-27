#include "../../stdc++.h"

using namespace std;

vector<int> climbingLeaderboard(vector<int> ranked, vector<int> player)
{
  set<int> s;
  vector<int> playerRank;

  for (auto r : ranked)
  {
    s.insert(r);
  }

  for (auto p : player)
  {
    s.insert(p);
  }

  int pIndex = 0;
  int sIndex = 0;
  for (auto rank : s)
  {
    if (player[pIndex] == rank)
    {
      playerRank.push_back(sIndex + 1);
      pIndex++;
    }
    else
    {
      sIndex++;
    }
  }

  return playerRank;
};

int main()
{
  vector<int> ranked{100, 100, 50, 40, 40, 20, 10};
  vector<int> player{5, 25, 50, 120};

  vector<int> pRanks = climbingLeaderboard(ranked, player);

  for (auto p : pRanks)
  {
    cout << p << ", ";
  }
}
