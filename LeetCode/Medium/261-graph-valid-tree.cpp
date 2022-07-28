// https://leetcode.com/problems/graph-valid-tree/
#include <iostream>
#include <vector>
#include <unordered_map>
#include <set>
#include <iostream>

using namespace std;

// Using DFS =================================================================================================
class SolutionOne
{
public:
  bool validTree(int n, vector<vector<int>> &edges)
  {
    unordered_map<int, vector<int>> graph = buildAdjacencyList(n, edges);
    vector<int> stack(1, 0);
    set<int> passed{};

    while (stack.size() > 0)
    {
      int node = stack.back();
      stack.pop_back();

      if (passed.find(node) != passed.end())
      {
        return false;
      }

      passed.insert(node);

      for (auto n : graph[node])
      {
        // because it is a 2 way graph and we are using DFS remove the edge going backwards.
        graph[n].erase(remove(graph[n].begin(), graph[n].end(), node), graph[n].end());
        stack.push_back(n);
      }
    }

    return passed.size() == n;
  }

private:
  unordered_map<int, vector<int>> buildAdjacencyList(int n, vector<vector<int>> &edges)
  {
    unordered_map<int, vector<int>> graph{};
    for (int i = 0; i < n; i++)
    {
      graph[i] = {};
    }

    for (auto e : edges)
    {
      graph[e[0]].push_back(e[1]);
      graph[e[1]].push_back(e[0]);
    }

    return graph;
  }
};

// Advanced Graph Theory + Union Find ========================================================================
// In order for a graph to be a valid tree it must have n-1 edges. Any less and it can't be connected,
// any more and it has to contain cycles.
class UnionFind
{
private:
  vector<int> root;

  int find(int n)
  {
    while (root[n] != n)
    {
      n = root[n];
    }
    return n;
  }

public:
  UnionFind(int n) : root(n)
  {
    for (int i = 0; i < n; i++)
    {
      root[i] = i;
    }
  }

  bool unionise(int a, int b)
  {
    int rootA = find(a);
    int rootB = find(b);

    if (rootA != rootB)
    {
      root[rootB] = rootA;
      return true;
    }

    // if both roots are the same and we connect two nodes then we have a cycle.
    return false;
  }
};

class SolutionTwo
{
public:
  bool validTree(int n, vector<vector<int>> &edges)
  {
    // in order to be a tree the number of edges must be one less than the number of nodes. If it is more then
    // we have a cycle and if it is less then not all nodes are connected.
    if (edges.size() != n - 1)
      return false;

    UnionFind uf(n);

    for (auto e : edges)
    {
      if (!uf.unionise(e[0], e[1]))
        return false;
    }

    return true;
  }
};

int main()
{

  vector<vector<int>> v{
      {0, 1},
      {0, 2},
      {0, 3},
      {1, 4},
  };

  cout << boolalpha;

  cout << SolutionTwo().validTree(5, v);
}