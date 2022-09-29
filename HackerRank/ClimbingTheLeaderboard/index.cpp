#include "../../stdc++.h"

using namespace std;

vector<int> climbingLeaderboard(vector<int> ranked, vector<int> player)
{
  vector<int> playerRanks;

  ranked.push_back(0);

  int currentRank = 1;
  int currentPlayerRankIndex = player.size() - 1;

  for (int i = 0; i < ranked.size(); i++)
  {
    if (i > 0 && ranked[i] < ranked[i - 1])
      currentRank++;

    while (player[currentPlayerRankIndex] >= ranked[i] && currentPlayerRankIndex >= 0)
    {
      playerRanks.push_back(currentRank);
      currentPlayerRankIndex--;
    }
  }

  reverse(playerRanks.begin(), playerRanks.end());

  return playerRanks;
};

int main()
{
  // vector<int> ranked{100, 100, 50, 40, 40, 20, 10};
  // vector<int> player{1, 2, 3, 5, 25, 50, 120};

  // vector<int> ranked{100, 100, 50, 40, 40, 20, 10};
  // vector<int> player{5, 25, 50, 120};

  // vector<int> ranked{100, 90, 90, 80, 75, 60};
  // vector<int> player{50, 65, 77, 90, 102};

  vector<int> ranked{100, 50, 20};
  vector<int> player{10, 20, 50, 60, 70, 100, 110, 120, 130};

  vector<int> pRanks = climbingLeaderboard(ranked, player);

  for (auto p : pRanks)
  {
    cout << p << ", ";
  }
}
