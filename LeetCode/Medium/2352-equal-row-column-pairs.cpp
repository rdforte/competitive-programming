// https://leetcode.com/problems/equal-row-and-column-pairs/
#include <iostream>
#include <vector>
#include <set>
#include <string>

using namespace std;

class Solution
{
public:
  int equalPairs(vector<vector<int>> &grid)
  {
    vector<string> rows;
    vector<string> cols(grid.size());

    int count = 0;

    for (auto row : grid)
    {
      string r = "";
      for (int i = 0; i < row.size(); i++)
      {
        r.append((to_string(row[i]).append(",")));
        cols[i].append(to_string(row[i]).append(","));
      }
      rows.push_back(r);
    }

    for (auto c : cols)
    {
      for (auto r : rows)
      {
        if (c == r)
          count++;
      }
    }

    return count;
  }
};

int main()
{
  // vector<vector<int>> grid{
  //     {3, 2, 1},
  //     {1, 7, 6},
  //     {2, 7, 7},
  // };

  vector<vector<int>> grid{
      {11, 1},
      {1, 11},
  };

  // vector<vector<int>> grid{
  //     {3, 1, 2, 2},
  //     {1, 4, 4, 5},
  //     {2, 4, 2, 2},
  //     {2, 4, 2, 2},
  // };

  cout << Solution().equalPairs(grid);
}