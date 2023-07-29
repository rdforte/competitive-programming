#include <bits/stdc++.h>

using namespace std;

int maxProfit(vector<int> &prices)
{
  int profit = 0;
  int buy = prices[0];

  for (auto p : prices)
  {
    profit = max(profit, p - buy);
    buy = min(buy, p);
  }

  return profit;
}

int main()
{
  vector<int> v{7, 1, 5, 3, 6, 4};
  cout << maxProfit(v);
}