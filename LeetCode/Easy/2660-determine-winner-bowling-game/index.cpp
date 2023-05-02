#include "../../../stdc++.h";

using namespace std;

class Solution
{
public:
  int isWinner(vector<int> &player1, vector<int> &player2)
  {
    int p1Score = 0, p2Score = 0, p1Multiplier = 0, p2Multiplier = 0;

    for (int i = 0; i < player1.size(); i++)
    {
      p1Score += p1Multiplier > 0 ? player1[i] * 2 : player1[i];
      p2Score += p2Multiplier > 0 ? player2[i] * 2 : player2[i];

      p1Multiplier--;
      p2Multiplier--;

      if (player1[i] == 10)
        p1Multiplier = 2;
      if (player2[i] == 10)
        p2Multiplier = 2;
    }

    return p1Score == p2Score ? 0 : p1Score > p2Score ? 1
                                                      : 2;
  }
};

int main()
{
}