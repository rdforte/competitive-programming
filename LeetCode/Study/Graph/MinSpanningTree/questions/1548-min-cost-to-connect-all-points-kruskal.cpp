// https://leetcode.com/problems/min-cost-to-connect-all-points/
#include <iostream>
#include <vector>
#include <cmath>
#include <queue>

using namespace std;

class UnionFind
{
private:
  vector<int> roots;
  int nConnections = 0;

public:
  UnionFind(int n) : roots(n)
  {
    for (int i = 0; i < n; i++)
    {
      roots[i] = i;
    }
  }

  int findRoot(int x)
  {
    while (roots[x] != x)
    {
      x = roots[x];
    }
    return x;
  }

  void join(int a, int b)
  {
    int rootA = findRoot(a);
    int rootB = findRoot(b);

    if (rootA != rootB)
    {
      roots[rootB] = rootA;
      nConnections++;
    }
  }

  bool isConnected(int a, int b)
  {
    return findRoot(a) == findRoot(b);
  }

  int getNumConnections()
  {
    return nConnections;
  }
};

class Edge
{
public:
  int point1;
  int point2;
  int cost;
  Edge(int point1, int point2, int cost)
  {
    this->point1 = point1;
    this->point2 = point2;
    this->cost = cost;
  }
};

class Solution
{
public:
  int minCostConnectPoints(vector<vector<int>> &points)
  {
    // Create our Min Heap of edges
    auto cmp = [](Edge left, Edge right)
    { return left.cost > right.cost; };

    priority_queue<Edge, vector<Edge>, decltype(cmp)>
        heap(cmp);

    // Keep track of edges as we need n-1 edges for a MST.
    int nEdges = 0;
    int cost = 0;

    // Initialize our UnionFind data Structure for connecting the edges
    UnionFind uf(points.size());

    // Associate each point with an index.
    // Calculate weighted Edges
    // Add weighted edge to min heap
    for (int i = 0; i < points.size() - 1; i++)
    {
      for (int j = i + 1; j < points.size(); j++)
      {
        int cost = abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1]);
        Edge e = Edge(i, j, cost);
        heap.push(e);
      }
    }

    // Go through min heap. check if the edge we are adding doesn't already have
    // points which are conected because if the points are already connected then
    // we would create a cycle.
    while (!heap.empty() && nEdges <= points.size() - 1)
    {
      Edge e = heap.top();
      heap.pop();

      if (!uf.isConnected(e.point1, e.point2))
      {
        uf.join(e.point1, e.point2);
        nEdges++;
        cost += e.cost;
      }
    }

    return cost;
  }
};

int main()
{
  vector<vector<int>> points{
      {0, 0},
      {2, 2},
      {3, 10},
      {5, 2},
      {7, 0},
  };

  cout << Solution().minCostConnectPoints(points);
}