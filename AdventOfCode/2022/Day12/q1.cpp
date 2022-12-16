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

bool canMoveToNextMountain(vector<vector<int>> &nodeWeights, vector<vector<int>> &mountainHeights, Node node, int row, int col)
{
  int r = node.second.first;
  int c = node.second.second;
  int curHeight = mountainHeights[r][c];
  char nextHeight = mountainHeights[row][col];
  return (nextHeight <= curHeight + 1) && ((node.first + 1) < nodeWeights[row][col]);
};

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

  Input input = Input();
  auto heightMap = input.heightMap;
  auto mountainWeight = input.mountainWeight;
  auto processed = input.processed;
  auto S = input.S;
  auto E = input.E;

  priority_queue<Node, vector<Node>, greater<Node>> q;

  q.push({0, S});

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

    cout << (char)curHeight << " : " << mountainWeight[row][col] << " -----------------\n";

    // Up
    int rowUp = row + 1;
    if (rowUp < heightMap.size())
    {
      cout << (char)heightMap[rowUp][col] << " : ";
      if (canMoveToNextMountain(mountainWeight, heightMap, node, rowUp, col))
      {
        mountainWeight[rowUp][col] = node.first + 1;
        q.push({mountainWeight[rowUp][col], {rowUp, col}});
      }
    }
    // Down
    int rowDown = row - 1;
    if (rowDown >= 0)
    {
      cout << (char)heightMap[rowDown][col] << " : ";
      if (canMoveToNextMountain(mountainWeight, heightMap, node, rowDown, col))
      {
        mountainWeight[rowDown][col] = node.first + 1;
        q.push({mountainWeight[rowDown][col], {rowDown, col}});
      }
    }
    // Left
    int colLeft = col - 1;
    if (colLeft >= 0)
    {
      cout << (char)heightMap[row][colLeft] << " : ";
      if (canMoveToNextMountain(mountainWeight, heightMap, node, row, colLeft))
      {
        mountainWeight[row][colLeft] = node.first + 1;
        q.push({mountainWeight[row][colLeft], {row, colLeft}});
      }
    }
    // Right
    int colRight = col + 1;
    if (colRight < heightMap[0].size())
    {
      cout << (char)heightMap[row][colRight];
      if (canMoveToNextMountain(mountainWeight, heightMap, node, row, colRight))
      {
        mountainWeight[row][colRight] = node.first + 1;
        q.push({mountainWeight[row][colRight], {row, colRight}});
      }
    }
    cout << "\n\n";
  }

  cout << mountainWeight[E.first][E.second]; // 517
}