// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/
#include <iostream>
#include <vector>
#include <set>

using namespace std;

class UnionFind
{
private:
  vector<int> root;
  vector<int> parents;

  int findRoot(int n)
  {
    while (root[n] != n)
    {
      n = root[n];
    }
    return n;
  }

public:
  UnionFind(int n) : root(n), parents(n)
  {
    for (int i = 0; i < n; i++)
    {
      root[i] = i;
      parents[i] = i;
    }
  }

  void unionise(int a, int b)
  {
    int rootA = findRoot(a);
    int rootB = findRoot(b);

    if (rootA != rootB)
    {
      root[rootB] = rootA;
      // rootB is no longer a parent so remove it
      parents.erase(remove(parents.begin(), parents.end(), rootB), parents.end());
    }
  }

  int connectedComponents()
  {
    return parents.size();
  }
};

class Solution
{
public:
  int countComponents(int n, vector<vector<int>> &edges)
  {
    if (n == 0)
      return 0;

    UnionFind uf(n);

    for (auto e : edges)
    {
      uf.unionise(e[0], e[1]);
    }

    return uf.connectedComponents();
  }
};

int main()
{
  vector<vector<int>> v{
      {0, 1},
      {1, 2},
      {3, 4},
  };

  cout << Solution().countComponents(5, v);
}