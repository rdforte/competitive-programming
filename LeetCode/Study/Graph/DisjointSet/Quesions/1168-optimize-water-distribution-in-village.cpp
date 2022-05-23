// https://leetcode.com/problems/optimize-water-distribution-in-a-village/
#include <iostream>
#include <vector>
#include <queue>   // gives us access to priority_queue
#include <utility> // gives us access to pair

using namespace std;

class Solution
{

public:
  int minCostToSupplyWater(int n, vector<int> &wells, vector<vector<int>> &pipes)
  {
    // Create a min heap to maintain the order of edges to be visited.
    // first = cost
    // second = next house
    auto cmp = [](auto a, auto b)
    { return a.first > b.second; };
    priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)> edgesHeap(cmp);

    // Represent graph as adjacency list.

    for (int i = 0; i < pipes.size(); i++)
    {
      cout << i << ", ";
    }

    cout << "\n--------------";

    return 1;
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