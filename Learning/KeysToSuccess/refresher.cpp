// gives us cin and cout
#include <iostream>
// gives us max, min and swap
#include <algorithm>
// gives us pair
#include <utility>

using namespace std;

int main()
{
  int a = 7;
  int b = 4;

  cout << "MIN & MAX"
       << "\n";

  cout << max(a, b) << "\n";
  cout << min(a, b) << "\n";

  // ===========================================================================

  cout << "SWAP"
       << "\n";

  swap(a, b);
  cout << a << " , " << b << "\n";

  // ===========================================================================

  cout << "PAIR"
       << "\n";

  // pair is very similar to creating a struct which has 2 properties.

  pair<int, int> coord = {a, b};
  cout << coord.first << " : " << coord.second << "\n";
}