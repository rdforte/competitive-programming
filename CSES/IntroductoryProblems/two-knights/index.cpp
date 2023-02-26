#include <bits/stdc++.h>

using namespace std;

typedef long long ll;

int main()
{
  ll n;
  cin >> n;

  for (ll i = 1; i <= n; i++)
  {
    ll gridCombinations = (pow(i, 2) * (pow(i, 2) - 1)) / 2;
    ll knightAttackCombos = 4 * ((i - 1) * (i - 2));
    cout << gridCombinations - knightAttackCombos << "\n";
  }
}