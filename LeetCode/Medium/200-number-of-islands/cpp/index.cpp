#include <bits/stdc++.h>

using namespace std;

// Space = O(2*m*n) = O(m*n) space for processed and space for queue.
// Time = O(m*n + m*n) = O(2(m*n)) = O(m*n). loop through each m and n as well as islands not processed.
class Solution
{
public:
  int numIslands(vector<vector<char>> &grid)
  {
    int islands = 0;
    vector<vector<bool>> processed(grid.size(), vector<bool>(grid[0].size(), false));
    stack<pair<int, int>> s;

    vector<pair<int, int>> directions = {
        {1, 0},
        {-1, 0},
        {0, 1},
        {0, -1}};

    for (int r = 0; r < grid.size(); r++)
    {
      for (int c = 0; c < grid[r].size(); c++)
      {
        if (processed[r][c] or grid[r][c] == '0')
          continue;

        islands++;
        s.push({r, c});

        while (!s.empty())
        {
          auto rc = s.top();
          s.pop();

          for (auto dir : directions)
          {
            int row = rc.first + dir.first;
            int col = rc.second + dir.second;
            if ((row >= 0 and row < grid.size()) and (col >= 0 and col < grid[0].size()) and !processed[row][col] and grid[row][col] == '1')
            {
              s.push({row, col});
            }
          }

          processed[rc.first][rc.second] = true;
        }
      }
    }

    return islands;
  }
};

int main()
{
  vector<vector<char>> grid = {
      {'1', '0', '1'},
      {'1', '0', '1'},
      {'1', '1', '1'},
  };
  cout << Solution().numIslands(grid);
}