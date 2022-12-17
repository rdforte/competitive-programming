#include <bits/stdc++.h>

using namespace std;

// node weight, with (row, col)
typedef pair<int, pair<int, int>> Node;

struct Input
{
  vector<vector<int>> heightMap;
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
      processed.push_back({});
      for (int i = 0; i < s.size(); i++)
      {
        heightMap.back().push_back(s[i]);
        processed.back().push_back(false);
        if (s[i] == 'S')
        {
          S = {row, i};
          heightMap.back().back() = (int)'a';
        }
        if (s[i] == 'E')
        {
          heightMap.back().back() = (int)'z';
          E = {row, i};
        }
      }
      row++;
    }
  }
};

vector<pair<int, int>> findNeighbours(int row, int col, vector<vector<int>> &heightMap)
{
  vector<pair<int, int>> neighbours = {
      {1, 0},
      {-1, 0},
      {0, 1},
      {0, -1},
  };
  vector<pair<int, int>> neighboursWithinHeight;
  for (auto n : neighbours)
  {
    int r = row + n.first;
    int c = col + n.second;
    if (not(r >= 0 and r < heightMap.size() and c >= 0 and c < heightMap[row].size()))
      continue;

    if (heightMap[r][c] - heightMap[row][col] <= 1)
      neighboursWithinHeight.push_back({r, c});
  }

  return neighboursWithinHeight;
};

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q2output.txt", "w", stdout);

  Input input = Input();
  auto heightMap = input.heightMap;
  auto processed = input.processed;
  auto S = input.S;
  auto E = input.E;

  int steps = 0;

  auto cmp = [](Node left, Node right)
  {
    return left.first > right.first;
  };

  priority_queue<Node, vector<Node>, decltype(cmp)> q(cmp);

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

    if (row == E.first and col == E.second)
    {
      steps = node.first;
      break;
    }

    for (auto n : findNeighbours(row, col, heightMap))
    {
      if ((char)heightMap[n.first][n.second] == 'a')
      {
        q.push({node.first, n});
      }
      else
      {
        q.push({node.first + 1, n});
      }
    }
  }

  cout << steps; // 517
}