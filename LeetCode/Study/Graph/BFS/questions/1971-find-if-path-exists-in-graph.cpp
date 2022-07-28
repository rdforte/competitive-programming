// https://leetcode.com/explore/learn/card/graph/620/breadth-first-search-in-graph/3894/

// https://leetcode.com/problems/find-if-path-exists-in-graph/
#include <iostream>
#include <vector>
#include <queue>

using namespace std;

class Solution
{
public:
  bool validPath(int n, vector<vector<int>> &edges, int source, int destination)
  {
    vector<vector<int>> adjacencyList(n);
    vector<bool> visited(n);

    for (auto e : edges)
    {
      adjacencyList[e[0]].push_back(e[1]);
      adjacencyList[e[1]].push_back(e[0]);
    }

    queue<int> q;

    q.push(source);

    while (!q.empty())
    {
      int node = q.front();
      q.pop();

      if (node == destination)
        return true;

      for (auto n : adjacencyList[node])
      {
        if (!visited[n])
        {
          visited[n] = true;
          q.push(n);
        }
      }
    }
    return false;
  }
};

int main()
{
  cout << boolalpha;
  vector<vector<int>> edges{
      {0, 1},
      {1, 2},
      {2, 0},
  };

  cout << Solution().validPath(3, edges, 0, 2);
}