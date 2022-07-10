#include <vector>
#include <algorithm>
#include <functional>

using namespace std;

class Solution
{
public:
  int fillCups(vector<int> &amount)
  {
    int seconds = 0;

    while (amount[0] > 0 || amount[1] > 0 || amount[2] > 0)
    {
      sort(amount.begin(), amount.end(), greater<int>());
      if (amount[0] > 0 && amount[1] > 0)
      {
        amount[0]--;
        amount[1]--;
        seconds++;
      }
      else
      {
        amount[0]--;
        seconds++;
      }
    }
    return seconds;
  }
};