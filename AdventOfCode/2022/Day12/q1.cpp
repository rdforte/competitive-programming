#include <bits/stdc++.h>

using namespace std;

struct Input
{
  vector<vector<int>> heightMap;
  vector<vector<int>> mountainWeight;
  vector<vector<bool>> processed;
  pair<int, int> S;
  pair<int, int> E;

  Input()
  {
    string s;
    int row = 0;
    while (getline(cin, s))
    {
      heightMap.push_back({});
      mountainWeight.push_back({});
      processed.push_back({});
      for (int i = 0; i < s.size(); i++)
      {
        heightMap.back().push_back(s[i]);
        mountainWeight.back().push_back(INT_MAX);
        processed.back().push_back(false);
        if (s[i] == 'S')
        {
          S = {row, i};
          heightMap.back().back() = (int)'a' - 1;
          mountainWeight.back().back() = 0;
        }
        if (s[i] == 'E')
        {
          heightMap.back().back() = (int)'z' + 1;
          E = {row, i};
        }
      }
      row++;
    }
  }
};

// node weight, with (row, col)
typedef pair<int, pair<int, int>> Node;

int main()
{
  freopen("test.txt", "r", stdin);
  freopen("q1output", "w", stdout);

  Input input = Input();
  auto heightMap = input.heightMap;
  auto mountainWeight = input.mountainWeight;
  auto processed = input.processed;
  auto S = input.S;
  auto E = input.E;

  priority_queue<Node, vector<Node>, greater<Node>> q;

  q.push({0, S});
  // processed[S.first][S.second] = true;

  while (!q.empty())
  {
    Node node = q.top();
    q.pop();
    int row = node.second.first;
    int col = node.second.second;
    if (processed[row][col])
      continue;

    processed[row][col] = true;

    int curHeight = heightMap[row][col];

    cout << curHeight << "\n";

    // Up
    if (row + 1 < heightMap.size())
    {
      char upHeight = heightMap[row + 1][col];
      // check is within range of 1 height
      if ((upHeight - curHeight) == 0 || (upHeight - curHeight) == 1)
      {
        // check if shorter path
        if (node.first + 1 < mountainWeight[row + 1][col])
        {
          mountainWeight[row + 1][col] = node.first + 1;
          q.push({mountainWeight[row + 1][col], {row + 1, col}});
        }
      }
    }
    // Down
    if (row - 1 >= 0)
    {
    }
    // Left
    if (row - 1 >= 0)
    {
    }
    // Right
    if (row + 1 < heightMap[0].size())
    {
    }
  }
}