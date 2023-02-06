https : // leetcode.com/problems/take-gifts-from-the-richest-pile/description/
#include <bits/stdc++.h>

        using namespace std;

class Solution
{
public:
  long long pickGifts(vector<int> &gifts, int k)
  {
    for (int i = 0; i < k; i++)
    {
      int m = 0;
      for (int j = 0; j < gifts.size(); j++)
      {
        if (gifts[j] > gifts[m])
        {
          m = j;
        }
      }
      if (gifts[m] > 1)
      {
        gifts[m] = sqrt(gifts[m]);
      }
    }
    long long total = 0;
    for (auto g : gifts)
    {
      total += g;
    }

    return total;
  }
};