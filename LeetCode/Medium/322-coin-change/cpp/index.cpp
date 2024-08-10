#include "../../../stdc++.h"

using namespace std;

int minNumberOfCoins(vector<int> &coins, int amount)
{
  vector<int> c(amount + 1, INT_MAX);
  c[0] = 0;

  for (int i = 1; i < int(c.size()); i++)
  {
    for (int j = 0; j < int(coins.size()); j++)
    {
      int coinsUsedPreviously = c[i - coins[j]];
      if (coins[j] <= i && coinsUsedPreviously < INT_MAX)
      {
        int numCoins = coinsUsedPreviously + 1;
        c[i] = min(c[i], numCoins);
      }
    }
  }

  return c[amount] == INT_MAX ? -1 : c[amount];
}

// To execute C++, please define "int main()"
int main()
{
  vector<int> coins{1, 2, 5};
  int amount = 11;

  cout << minNumberOfCoins(coins, amount);

  return 0;
}
