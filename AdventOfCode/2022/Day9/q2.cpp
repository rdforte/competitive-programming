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

  void move(char direction)
  {
    if (direction == 'U')
      row++;
    if (direction == 'D')
      row--;
    if (direction == 'L')
      col--;
    if (direction == 'R')
      col++;
  }

  void follow(Position p)
  {
    if ((abs(p.row - row) > 1) or
        (abs(p.col - col) > 1))
    {
      int dirRow = p.row - row;
      int dirCol = p.col - col;
      row += abs(dirRow) == 2 ? dirRow / 2 : dirRow;
      col += abs(dirCol) == 2 ? dirCol / 2 : dirCol;
    }
  }
};

void addTailPosition(int row, int col, set<string> &s)
{
  s.insert((to_string(row) + "," + to_string(col)));
}

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q2output.txt", "w", stdout);

  set<string> tailPositions;

  Position head = Position(0, 0);
  Position tail = Position(0, 0);
  vector<Position> positions(10, Position(0, 0));

  tailPositions.insert("0,0");

  char direction;
  int spaces;

  while (cin >> direction >> spaces)
  {
    for (int i = 1; i <= spaces; i++)
    {
      positions[0].move(direction);
      for (int i = 1; i < positions.size(); i++)
      {
        positions[i].follow(positions[i - 1]);
      }
      addTailPosition(positions.back().row, positions.back().col, tailPositions);
    }
  }

  cout << tailPositions.size() << "\n";
}