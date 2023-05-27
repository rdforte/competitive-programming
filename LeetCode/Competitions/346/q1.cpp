#include "../../../stdc++.h"

using namespace std;

class Solution
{
public:
  int minLength(string s)
  {
    string oldS = s;
    string newS = "";

    while (true)
    {
      for (int i = 0; i < oldS.size();)
      {
        if (i < oldS.size() - 1 && ((oldS[i] == 'A' && oldS[i + 1] == 'B') || (oldS[i] == 'C' && oldS[i + 1] == 'D')))
        {
          i += 2;
          continue;
        }
        newS += oldS[i];
        i++;
      }

      if (newS == oldS)
        break;

      oldS = newS;
      newS = "";
    }

    return oldS.size();
  }
};

int main()
{
  cout << Solution().minLength("ACBBD");
}
