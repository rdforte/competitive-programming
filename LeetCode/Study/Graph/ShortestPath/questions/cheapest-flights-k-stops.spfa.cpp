// https://leetcode.com/problems/cheapest-flights-within-k-stops/
#include <iostream>
#include <vector>
#include <limits> // int max, min
#include <queue>
#include <utility> // pair

using namespace std;

class Solution
{
public:
  int findCheapestPrice(int n, vector<vector<int>> &flights, int src, int dst, int k)
  {
    // Create adjacency list
    vector<pair<int, int>> adjacencyList[n]; // to, price
    for (auto f : flights)
    {
      int from = f[0];
      int to = f[1];
      int price = f[2];

      adjacencyList[from].push_back({to, price});
    }

    // Setup Queue
    queue<pair<int, int>> q;
    // Setup list for keeping track of prices.
    vector<int> prices(n, INT_MAX);
    // Add the starting point to the queue
    q.push({src, 0});
    // update source price to 0
    prices[src] = 0;

    int stops = 0;

    // iterate throught the queue
    while (!q.empty())
    {
      // Break the loop if we have reached our max stops.
      if (stops >= k + 1)
        break;

      int nStopsAtSameLevel = q.size();

      while (nStopsAtSameLevel--)
      {
        auto edge = q.front();
        q.pop();

        int to = edge.first;
        int price = edge.second;

        for (auto n : adjacencyList[to])
        {
          int nTo = n.first;
          int nPrice = n.second;

          if ((price + nPrice) < prices[nTo])
          {
            prices[nTo] = price + nPrice;
            q.push({nTo, nPrice + price}); // add to previouse price
          }
        }
      }
      stops += 1;
    }

    return prices[dst] == INT_MAX ? -1 : prices[dst];
  }
};

int main()
{
  // int n = 4;
  // vector<vector<int>> flights{
  //     {0, 1, 100},
  //     {1, 2, 100},
  //     {2, 0, 100},
  //     {1, 3, 600},
  //     {2, 3, 200},
  // };

  // int src = 0;
  // int dst = 3;
  // int k = 1;

  int n = 4;
  vector<vector<int>> flights{
      {0, 1, 1},
      {0, 2, 5},
      {1, 2, 1},
      {2, 3, 1},
  };

  int src = 0;
  int dst = 3;
  int k = 1;

  cout << Solution().findCheapestPrice(n, flights, src, dst, k);
}
