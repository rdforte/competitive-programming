// https://leetcode.com/problems/all-paths-from-source-to-target/

#include <iostream>
#include <vector>
#include <queue>

using namespace std;

class Solution
{
public:
  vector<vector<int>> allPathsSourceTarget(vector<vector<int>> &graph)
  {
    int len = graph.size() - 1;
    queue<vector<int>> q;
    vector<vector<int>> paths{};

    q.push(vector<int>{0});

    while (!q.empty())
    {
      vector<int> node = q.front();
      q.pop();

      if (node.back() == len)
      {
        paths.push_back(node);
        continue;
      }

      for (auto n : graph[node.back()])
      {
        vector<int> p = node;
        p.push_back(n);
        q.push(p);
      }
    }

    return paths;
  }
};

int main()
{
  // vector<vector<int>> graph{
  //     {1, 2},
  //     {3},
  //     {3},
  //     {},
  // };

  // vector<vector<int>> graph{
  //     {4, 3, 1},
  //     {3, 2, 4},
  //     {3},
  //     {4},
  //     {},
  // };

  vector<vector<int>> graph{
      {4, 3, 1},
      {3, 2, 4},
      {},
      {4},
      {},
  };

  // vector<vector<int>> graph{
  //     {2},
  //     {},
  //     {1},
  // };

  vector<vector<int>> paths = Solution().allPathsSourceTarget(graph);

  for (auto p : paths)
  {
    cout << "NEW PATH --------------- \n";
    for (auto i : p)
    {
      cout << i << ",";
    }
    cout << "--------------------------\n";
  }
}