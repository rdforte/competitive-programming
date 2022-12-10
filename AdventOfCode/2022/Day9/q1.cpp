#include <bits/stdc++.h>

using namespace std;

struct Position
{
public:
  int row;
  int col;
  Position(int r, int c)
  {
    row = r;
    col = c;
  }
};

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

  set<string> tailPositions;

  Position headPosition = Position(0, 0);
  Position tailPosition = Position(0, 0);

  tailPositions.insert("0,0");

  int prevRow = headPosition.row;
  int prevCol = headPosition.col;

  char direction;
  int spaces;

  while (cin >> direction >> spaces)
  {
    for (int i = 1; i <= spaces; i++)
    {

      if (direction == 'U')
        headPosition.row++;
      if (direction == 'D')
        headPosition.row--;
      if (direction == 'L')
        headPosition.col--;
      if (direction == 'R')
        headPosition.col++;

      if ((abs(headPosition.row - tailPosition.row) > 1) or
          (abs(headPosition.col - tailPosition.col) > 1))
      {
        tailPosition.col = prevCol;
        tailPosition.row = prevRow;
        tailPositions.insert((to_string(tailPosition.row) + "," + to_string(tailPosition.col)));
      }

      prevRow = headPosition.row;
      prevCol = headPosition.col;
    }
  }

  cout << tailPositions.size() << "\n";
}