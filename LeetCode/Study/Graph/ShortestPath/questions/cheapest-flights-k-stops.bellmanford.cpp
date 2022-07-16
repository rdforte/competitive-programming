// https://leetcode.com/problems/cheapest-flights-within-k-stops/
#include <iostream>
#include <vector>
#include <set>
#include <queue>
#include <algorithm>
#include <limits>
#include <queue>
#include <utility>

using namespace std;

// Implemented with Bellman Ford.
// time complexity is O(E * K) where E is the edges and K is the number of stops.

class Solution
{
public:
  int findCheapestPrice(int n, vector<vector<int>> &flights, int src, int dst, int k)
  {
    // list of prices
    vector<int> prices(n, INT_MAX);
    // set the src price to 0. This is our source.
    prices[src] = 0;

    // iterate through the loop K + 1 times. + 1 because we want the edges to go through one node. if it was just
    // K then this will stop at the stop but we want to go through the stop and arrive at the destination.
    for (int i = 0; i <= k; i++)
    {
      cout << i << " ----------------------------------\n";
      // Create a copy of our prices.
      vector<int> tempPrices(prices);

      // Loop through the flights
      for (auto flight : flights)
      {
        int from = flight[0];
        int to = flight[1];
        int price = flight[2];

        cout << "* " << from << " : " << to << " : " << price << "\n";

        // if the price is INT_MAX then we can not reach this node.
        if (prices[from] == INT_MAX)
          continue;

        // on first iteration of the loop
        // 0 + 100 < INT_MAX
        if ((prices[from] + price) < tempPrices[to])
        {
          cout << "prices[from] " << prices[from] << "\n";
          tempPrices[to] = prices[from] + price;
        }
      }
      prices = tempPrices;
    }

    return prices[dst] == INT_MAX ? -1 : prices[dst];
  }
};

// SPFA shortest path faster algorithm.

class SolutionTwo
{
public:
  int findCheapestPrice(int n, vector<vector<int>> &flights, int src, int dst, int k)
  {
  }
};

int main()
{
  int n = 5;
  vector<vector<int>> flights{
      {4, 1, 1},
      {1, 2, 3},
      {0, 3, 2},
      {0, 4, 10},
      {3, 1, 1},
      {1, 4, 3},
  };

  int src = 2;
  int dst = 1;
  int k = 1;

  cout << Solution().findCheapestPrice(n, flights, src, dst, k);
}
