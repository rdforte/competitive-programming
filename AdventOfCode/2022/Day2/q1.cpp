#include <bits/stdc++.h>

using namespace std;

struct RPCObject
{
private:
  char object;
  char beats;
  char draws;

public:
  RPCObject(char o, char b, char d) : object(o), beats(b), draws(d){};

  int getPointsFromPlay(char play)
  {
    if (play == beats)
      return 6;
    if (play == draws)
      return 3;

    return 0;
  }
};

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1Output.txt", "w", stdout);

  RPCObject Rock('X', 'C', 'A');
  RPCObject Paper('Y', 'A', 'B');
  RPCObject Scissors('Z', 'B', 'C');

  string s;
  int totalPoints = 0;
  while (getline(cin, s))
  {
    auto opponent = s[0];
    auto self = s[2];

    int points = 0;
    if (self == 'X')
    {
      points += Rock.getPointsFromPlay(opponent) + 1;
    }
    else if (self == 'Y')
    {
      points += Paper.getPointsFromPlay(opponent) + 2;
    }
    else if (self == 'Z')
    {
      points += Scissors.getPointsFromPlay(opponent) + 3;
    }

    totalPoints += points;
  }
  cout << totalPoints;
}