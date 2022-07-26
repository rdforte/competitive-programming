#include <vector>
#include <utility>
#include <string>
#include <iostream>

using namespace std;

class Solution
{
private:
  string up = "up";
  string down = "down";

public:
  vector<int> findDiagonalOrder(vector<vector<int>> &mat)
  {
    vector<int> diagonalOrder;

    pair<int, int> location(0, 0); // row, col
    string direction("up");

    while (diagonalOrder.size() < (mat.size() * mat[0].size()))
    {
      cout << "starting dir: " << direction << "\n";
      cout << location.first << " : " << location.second << "\n";
      diagonalOrder.push_back(mat[location.first][location.second]);

      if (direction == up)
      {
        cout << direction << "\n";
        int rowAbove = location.first - 1;
        int colRight = location.second + 1;
        if (rowAbove >= 0 && colRight < mat.size())
        {
          location = {rowAbove, colRight};
        }
        else
        {
          direction = down;
          if (location.second + 1 < mat.size())
          {
            location = {location.first, location.second + 1};
          }
          else
          {
            location = {location.first + 1, location.second};
          }
        }
      }
      else
      {
        cout << direction << "\n";
        int rowUnder = location.first + 1;
        int colLeft = location.second - 1;
        if (rowUnder < mat.size() && colLeft >= 0)
        {
          location = {rowUnder, colLeft};
        }
        else
        {
          direction = up;
          if (location.first + 1 < mat.size())
          {
            location = {location.first + 1, location.second};
          }
          else
          {
            location = {location.first, location.second + 1};
          }
        }
      }
    }

    return diagonalOrder;
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