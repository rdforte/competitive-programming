#include <bits/stdc++.h>

using namespace std;

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q2output.txt", "w", stdout);

  vector<vector<int>> grid;
  string rowInput;
  while (getline(cin, rowInput))
  {
    vector<int> tempRow;
    for (int i = 0; i < rowInput.size(); i++)
    {
      tempRow.push_back(stoi(string(1, rowInput[i])));
    }
    grid.push_back(tempRow);
  }

  int mostScenic = 0;
  for (int row = 0; row < grid.size(); row++)
  {
    for (int col = 0; col < grid[row].size(); col++)
    {
      int current = grid[row][col];

      int left = 0;
      int right = 0;
      int up = 0;
      int down = 0;

      for (int i = col - 1; i >= 0; i--)
      {
        left++;
        int tree = grid[row][i];
        if (tree >= current)
          break;
      }

      for (int i = col + 1; i < grid[row].size(); i++)
      {
        right++;
        int tree = grid[row][i];
        if (tree >= current)
          break;
      }

      for (int i = row - 1; i >= 0; i--)
      {
        up++;
        int tree = grid[i][col];
        if (tree >= current)
          break;
      }

      for (int i = row + 1; i < grid.size(); i++)
      {
        down++;
        int tree = grid[i][col];
        if (tree >= current)
          break;
      }

      mostScenic = max(mostScenic, (up * down * left * right));
    }
  }
  cout << mostScenic;
}