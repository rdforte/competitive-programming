#include <bits/stdc++.h>

using namespace std;

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

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

  int visible = 0;
  for (int row = 0; row < grid.size(); row++)
  {
    for (int col = 0; col < grid[row].size(); col++)
    {
      int current = grid[row][col];

      int leftMax = INT_MIN;
      int rightMax = INT_MIN;
      int upMax = INT_MIN;
      int downMax = INT_MIN;

      for (int i = col - 1; i >= 0; i--)
      {
        leftMax = max(grid[row][i], leftMax);
      }

      for (int i = col + 1; i < grid[row].size(); i++)
      {
        rightMax = max(grid[row][i], rightMax);
      }

      for (int i = row - 1; i >= 0; i--)
      {
        upMax = max(grid[i][col], upMax);
      }

      for (int i = row + 1; i < grid.size(); i++)
      {
        downMax = max(grid[i][col], downMax);
      }

      if (current > leftMax || current > rightMax || current > upMax || current > downMax)
      {
        visible++;
      }
    }
  }
  cout << visible;
}