#include <bits/stdc++.h>

using namespace std;

struct RPCObject
{
private:
  int loses;
  int draws;
  int wins;

public:
  RPCObject(int l, int d, int w) : loses(l), draws(d), wins(w){};

  int getPointsFromPlay(char play)
  {
    if (play == 'X')
      return loses;
    if (play == 'Y')
      return 3 + draws;

    return 6 + wins;
  }
};

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q2Output.txt", "w", stdout);

  RPCObject Rock(3, 1, 2);
  RPCObject Paper(1, 2, 3);
  RPCObject Scissors(2, 3, 1);

  string s;
  int totalPoints = 0;
  while (getline(cin, s))
  {
    auto opponent = s[0];
    auto self = s[2];

    int points = 0;
    if (opponent == 'A')
    {
      points += Rock.getPointsFromPlay(self);
    }
    else if (opponent == 'B')
    {
      points += Paper.getPointsFromPlay(self);
    }
    else if (opponent == 'C')
    {
      points += Scissors.getPointsFromPlay(self);
    }

    totalPoints += points;
  }
  cout << totalPoints;
}