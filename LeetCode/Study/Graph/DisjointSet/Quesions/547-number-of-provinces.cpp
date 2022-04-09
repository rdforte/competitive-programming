#include <iostream>
#include <vector>

class UnionFind
{
private:
  std::vector<int> root;
  int count;

public:
  UnionFind(int sz) : root(sz), count(sz)
  {
    for (int i = 0; i < sz; i++)
    {
      root[i] = i;
    }
  };

  int findRoot(int x)
  {
    while (x != root[x])
    {
      x = root[x];
    }
    return x;
  }

  // x will become the root of y
  void unionSet(int x, int y)
  {
    int rootX = findRoot(x);
    int rootY = findRoot(y);

    if (rootX != rootY)
    {
      root[rootY] = rootX;
      count--;
    }
  }

  bool isConnected(int x, int y)
  {
    return findRoot(x) == findRoot(y);
  }

  int getCount()
  {
    return count;
  }
};

class Solution
{
public:
  int findCircleNum(std::vector<std::vector<int>> &isConnected)
  {
    int conSize = isConnected.size();

    if (conSize == 0)
      return 0;

    UnionFind uFind(conSize);

    for (int i = 0; i < conSize; i++)
    {
      for (int j = 0; j < conSize; j++)
      {
        if (isConnected[i][j] == 1)
        {
          uFind.unionSet(i, j);
        }
      }
    }

    return uFind.getCount();
  }
};

int main()
{
  std::vector<std::vector<int>> v{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}};

  std::cout << Solution().findCircleNum(v);
}