#include <iostream>
#include <vector>
#include <ios>

// Prefered over quick find

class UnionFind
{
public:
  UnionFind(int sz) : root(sz)
  {
    for (int i = 0; i < sz; i++)
    {
      root[i] = i;
    }
  }

  // Traverse through the vertices till we get the root of x
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
    int rootX = find(x); // Get the root of x
    int rootY = find(y); // Get the root of y
    if (rootX != rootY)
    {
      root[rootY] = rootX; // Set the root of Y to now be root node x
    }
  }

  bool connected(int x, int y)
  {
    return find(x) == find(y);
  }

private:
  std::vector<int> root;
};

// Test Case
int main()
{
  // for displaying booleans as literal strings, instead of 0 and 1
  std::cout << std::boolalpha;
  UnionFind uf(10);
  // 1-2-5-6-7 | 3-8-9 | 4
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