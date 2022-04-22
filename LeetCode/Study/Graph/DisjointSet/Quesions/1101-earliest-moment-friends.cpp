// https://leetcode.com/problems/the-earliest-moment-when-everyone-become-friends/

#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class UnionFind
{
private:
  vector<int> root;
  int friends;
  int timestamp = 0;

  int findRoot(int n)
  {
    while (root[n] != n)
    {
      n = root[n];
    }
    return n;
  }

public:
  UnionFind(int n) : root(n), friends(n)
  {
    for (int i = 0; i < n; i++)
    {
      root[i] = i;
    }
  }

  bool unionise(long stamp, int a, int b)
  {
    int rootA = findRoot(a);
    int rootB = findRoot(b);

    if (rootA != rootB)
    {
      root[rootB] = rootA;
      timestamp = stamp;
      friends -= 1;
    }

    return friends == 1;
  }

  long getTimestamp()
  {
    return timestamp;
  }
};

class Solution
{
public:
  int earliestAcq(vector<vector<int>> &logs, int n)
  {
    if (n == 0)
      return 0;

    // sort the timestamps so we make sure that we always get the earliest time at which they were friends.
    sort(logs.begin(), logs.end());

    UnionFind uf(n);

    for (auto l : logs)
    {
      if (uf.unionise(l[0], l[1], l[2]))
      {
        return uf.getTimestamp();
      }
    }

    return -1;
  }
};

int main()
{
  // vector<vector<int>> v{
  //     {20190101, 0, 1},
  //     {20190104, 3, 4},
  //     {20190107, 2, 3},
  //     {20190211, 1, 5},
  //     {20190224, 2, 4},
  //     {20190301, 0, 3},
  //     {20190312, 1, 2},
  //     {20190322, 4, 5}
  //     };
  vector<vector<int>> v{
      {9, 3, 0},
      {0, 2, 1},
      {8, 0, 1},
      {1, 3, 2},
      {2, 2, 0},
      {3, 3, 1},
  };

  sort(v.begin(), v.end());

  cout << Solution().earliestAcq(v, 4);
}