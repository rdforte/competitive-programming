// https://leetcode.com/problems/optimize-water-distribution-in-a-village/
#include <iostream>
#include <vector>
#include <queue>   // gives us access to priority_queue
#include <utility> // gives us access to pair
#include <set>

using namespace std;

// PRIMS ALGORITHM
// N = houses, M = number of pipes.
// Time Complexity is O((N + M)*log(N * M))
// Space Complexity is O(N + M)
class Solution
{
public:
  int minCostToSupplyWater(int n, vector<int> &wells, vector<vector<int>> &pipes)
  {
    // Create a min heap to maintain the order of edges to be visited.
    // first = cost
    // second = next house
    // I want the first element to always be greater than the second so that the top is always the smallest.
    auto cmp = [](auto a, auto b)
    { return a.first > b.first; };
    priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)> edgesHeap(cmp);

    // Represent graph as adjacency list.
    vector<vector<pair<int, int>>> adjacencyList(n + 1);
    for (int i = 0; i <= n; i++)
    {
      adjacencyList[i] = {};
    }

    // add a vertual vertex indexed with 0, and then add an edge to each of the houses weighted by the cost.
    for (int i = 0; i < wells.size(); i++)
    {
      pair<int, int> virtualEdge(wells[i], i + 1);
      adjacencyList[0].push_back(virtualEdge);
      // Add the edge to the min heap.
      edgesHeap.push(virtualEdge);
    }

    // add the bidirectional edges to the adjacencyList.
    for (int i = 0; i < pipes.size(); i++)
    {
      int house1 = pipes[i][0];
      int house2 = pipes[i][1];

      int cost = pipes[i][2];

      adjacencyList[house1].push_back(pair(cost, house2));
      adjacencyList[house2].push_back(pair(cost, house1));
    }

    // kick off the exploration from the virtual vertex 0.
    set<int> houseSet{};
    houseSet.insert(0);

    int totalCost = 0;

    while (houseSet.size() <= n)
    {
      pair<int, int> edge = edgesHeap.top();
      edgesHeap.pop();

      int cost = edge.first;
      int nextHouse = edge.second;

      if (houseSet.find(nextHouse) != houseSet.end())
      {
        continue;
      }

      houseSet.insert(nextHouse);
      totalCost += cost;

      // expand the candidates of edges to choose from in the next round
      for (auto neighbourEdge : adjacencyList[nextHouse])
      {
        if (houseSet.find(neighbourEdge.second) == houseSet.end())
        {
          edgesHeap.push(neighbourEdge);
        }
      }
    }

    return totalCost;
  }
};

int main()
{
  int n = 3;
  vector<int> wells{1, 2, 2};
  vector<vector<int>> pipes{
      {1, 2, 1},
      {2, 3, 1}};

  cout << Solution().minCostToSupplyWater(n, wells, pipes);
}