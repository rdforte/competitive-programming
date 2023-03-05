#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int passThePillow(int n, int time)
  {
    bool isForward = true;
    int j = 1;
    for (int i = 0; i < time; i++)
    {
      if (isForward)
      {
        j++;
      }
      else
        j--;

      if (j == n || j == 1)
        isForward = !isForward;
    }
    return j;
  }
};

int main()
{
  // Solution
  cout << Solution().passThePillow(4, 5);
}
