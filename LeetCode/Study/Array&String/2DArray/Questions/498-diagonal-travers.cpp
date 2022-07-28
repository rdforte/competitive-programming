// https://leetcode.com/problems/diagonal-traverse/
#include <vector>
#include <utility>
#include <string>
#include <iostream>
#include <algorithm>
#include <functional>

using namespace std;

class Solution
{
public:
  vector<int> findDiagonalOrder(vector<vector<int>> &mat)
  {
    if (mat.size() == 0)
      return {};

    int numRows = mat.size();
    int numCols = mat[0].size();

    vector<int> diagonals;

    vector<int> currentDiagonal;

    // iterate over start of row and columns.
    for (int i = 0; i < numRows + numCols - 1; i++)
    {
      currentDiagonal.clear();

      // determine the head of the diagonal.
      int row = i < numCols ? 0 : (i - numCols + 1);
      int col = i < numCols ? i : (numCols - 1);

      while (row < numRows && col >= 0)
      {
        currentDiagonal.push_back(mat[row][col]);

        row++;
        col--;
      }

      if (i % 2 == 0)
        reverse(currentDiagonal.begin(), currentDiagonal.end());

      for (auto d : currentDiagonal)
        diagonals.push_back(d);
    }

    return diagonals;
  }
};

int main()
{
  // 1,2,4,7,5,3,6,8,9
  // vector<vector<int>> mat{
  //     {1, 2, 3},
  //     {4, 5, 6},
  //     {7, 8, 9},
  // };

  // 1,2,3,4
  // vector<vector<int>> mat{
  //     {1, 2},
  //     {3, 4},
  // };

  // 2, 3
  vector<vector<int>> mat{
      {2, 3},
  };

  vector<int> sol = Solution().findDiagonalOrder(mat);

  for (auto s : sol)
  {
    cout << s << ", ";
  }
}