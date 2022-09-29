#include "../../../stdc++.h"

using namespace std;

class UnionFind
{
private:
  vector<int> roots;

public:
  UnionFind(int n)
  {
    for (int i = 0; i < n; i++)
    {
      roots.push_back(i);
    }
  }

  int findRoot(int n)
  {
    if (roots[n] == n)
      return n;

    roots[n] = findRoot(roots[n]);

    return roots[n];
  }

  void joinNodes(int a, int b)
  {
    int rootA = findRoot(a);
    int rootB = findRoot(b);

    if (rootA != rootB)
    {
      roots[rootB] = rootA;
    }
  }

  vector<int> &getRoots()
  {
    return roots;
  }
};

int countComponents(int n, vector<vector<int>> &edges)
{
  UnionFind uf = UnionFind(n);

  for (auto e : edges)
  {
    uf.joinNodes(e[0], e[1]);
  }

  set<int> parentNodes;
  for (auto i : uf.getRoots())
  {
    parentNodes.insert(uf.findRoot(i));
  }

  return parentNodes.size();
};

int main()
{
  // int n = 5;
  // vector<vector<int>> e{
  //     {0, 1},
  //     {1, 2},
  //     {3, 4},
  // };

  int n = 4;
  vector<vector<int>> e{
      {2, 3},
      {1, 2},
      {1, 3},
  };

  cout << countComponents(n, e);
}