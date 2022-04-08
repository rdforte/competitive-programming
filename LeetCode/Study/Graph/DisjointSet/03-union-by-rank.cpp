#include <iostream>
#include <vector>
#include <ios>

class UnionFind
{
private:
  std::vector<int> root;
  std::vector<int> rank; // The rank represents the height of the graph

public:
  UnionFind(int sz) : root(sz), rank(sz)
  {
    for (int i = 0; i < sz; i++)
    {
      root[i] = i;
      rank[i] = 1;
    }
  }

  int find(int x)
  {
    while (x != root[x])
    {
      x = root[x];
    }
    return x;
  }

  void unionSet(int x, int y)
  {
    int rootX = find(x);
    int rootY = find(y);
    if (rootX != rootY)
    {
      if (rank[rootX] > rank[rootY])
      {
        root[rootY] = rootX; // RootX has the higher rank so we set this to Y
      }
      else if (rank[rootX] < rank[rootY])
      {
        root[rootX] = rootY; // RootY has the higher rank so we set this to X
      }
      else
      {
        root[rootY] = rootX; // They are both the same so we just set the root of Y to X
        rank[rootX] += 1;    // and now because X is the root we have to increment its rank by 1
      }
    }
  }

  bool connected(int x, int y)
  {
    return find(x) == find(y);
  }
};

// Test Case
int main()
{
  // for displaying booleans as literal strings, instead of 0 and 1
  std::cout << std::boolalpha;
  UnionFind uf(10);
  // 1-2-5-6-7 3-8-9 4
  uf.unionSet(1, 2);
  uf.unionSet(2, 5);
  uf.unionSet(5, 6);
  uf.unionSet(6, 7);
  uf.unionSet(3, 8);
  uf.unionSet(8, 9);
  std::cout << uf.connected(1, 5) << std::endl; // true
  std::cout << uf.connected(5, 7) << std::endl; // true
  std::cout << uf.connected(4, 9) << std::endl; // false
  // 1-2-5-6-7 3-8-9-4
  uf.unionSet(9, 4);
  std::cout << uf.connected(4, 9) << std::endl; // true

  return 0;
}