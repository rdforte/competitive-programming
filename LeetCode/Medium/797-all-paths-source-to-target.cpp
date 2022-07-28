// https://leetcode.com/problems/all-paths-from-source-to-target/

#include <iostream>
#include <vector>
#include <stack>

using namespace std;

class Solution
{
public:
  vector<vector<int>> allPathsSourceTarget(vector<vector<int>> &graph)
  {
    int len = graph.size() - 1;
    stack<vector<int>> s;
    vector<vector<int>> paths{};

    s.push(vector<int>{0});

    while (s.size() > 0)
    {
      vector<int> node = s.top();
      s.pop();

      if (node[node.size() - 1] == len)
      {
        paths.push_back(node);
        continue;
      }

      if (graph[node[node.size() - 1]].size() == 0)
      {
        continue;
      }

      for (auto n : graph[node[node.size() - 1]])
      {
        vector<int> p = node;
        p.push_back(n);
        s.push(p);
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