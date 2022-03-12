#include <iostream>
#include <vector>
#include <ios>

class UnionFind
{
public:
  UnionFind(int sz) : root(sz) // initialise the array with the size
  {
    for (int i = 0; i < sz; i++)
    {
      root[i] = i; // set each array index to be its own root
    }
  }

  int find(int x)
  {
    return root[x]; // This is O(1) to find the root
  }

  // This will join the graphs together. This is O(n) time
  void unionSet(int x, int y)
  {
    int rootX = find(x); // find the root node of x
    int rootY = find(y); // find the root node of y
    if (rootX != rootY)
    {
      for (int i = 0; i < root.size(); i++)
      {
        if (root[i] == rootY) // find all vertices with a root node Y
        {
          root[i] = rootX; // set the new root node to be X
        }
      }
    }
  }

  // check to see if the vertices are connected. O(1) time.
  bool connected(int x, int y)
  {
    return find(x) == find(y);
  }

private:
  std::vector<int> root; // The indices refer to the vertices and its values are the root nodes.
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