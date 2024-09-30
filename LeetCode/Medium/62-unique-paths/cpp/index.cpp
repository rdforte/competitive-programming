#include <bits/stdc++.h>

using namespace std;

// Space Complexity = O(m * n * moves) moves is constant so we can say O(m*n)
// Time Complexity = O(m * n) as we loop through each grid once
class Solution
{
public:
  int uniquePaths(int m, int n)
  {
    vector<vector<int>> grid(m, vector<int>(n, 0));
    grid[0][0] = 1;

    vector<pair<int, int>> moves{{1, 0}, {0, 1}};

    for (int i = 0; i < grid.size(); i++)
    {
      for (int j = 0; j < grid[i].size(); j++)
      {
        for (auto m : moves)
        {
          int row = i + m.first;
          int col = j + m.second;
          if ((row >= 0 && row < grid.size()) && (col >= 0 && col < grid[i].size()))
          {
            grid[row][col] += grid[i][j];
          }
        }
      }
    }

    return grid.back().back();
  }
};

int main()
{
  cout << Solution().uniquePaths(3, 2);
}