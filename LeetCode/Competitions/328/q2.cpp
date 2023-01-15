#include <bits/stdc++.h>

using namespace std;

class SolutionBruteForce
{
public:
  vector<vector<int>> rangeAddQueries(int n, vector<vector<int>> &queries)
  {
    vector<vector<int>> grid(n, vector(n, 0));
    for (auto q : queries)
    {
      for (int r = q[0]; r <= q[2] && r < n; r++)
      {
        for (int c = q[1]; c <= q[3] && c < n; c++)
        {
          grid[r][c]++;
        }
      }
    }
    return grid;
  }
};

// Scan-line algorithm.
// Space = O(n*n) we need to store the answer grid.
// Time = O(q*n + n*n) we loop through each query and n number of rows + each cell in the grid.
class Solution
{
public:
  vector<vector<int>> rangeAddQueries(int n, vector<vector<int>> &queries)
  {
    vector<vector<int>> grid(n, vector(n, 0));
    for (auto q : queries)
    {
      for (int r = q[0]; r < n && r <= q[2]; r++)
      {
        grid[r][q[1]]++;
        if (q[3] + 1 < n)
          grid[r][q[3] + 1]--;
      }
    }

    for (int r = 0; r < n; r++)
    {
      for (int c = 1; c < n; c++)
      {
        grid[r][c] += grid[r][c - 1];
      }
    }
    return grid;
  }
};

int main()
{
  vector<vector<int>> queries{
      {1, 1, 2, 2},
      {0, 0, 1, 1}};
  auto sol = Solution().rangeAddQueries(3, queries);
}
