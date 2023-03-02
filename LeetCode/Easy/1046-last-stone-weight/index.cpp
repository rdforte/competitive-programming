#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  int lastStoneWeight(vector<int> &stones)
  {
    sort(stones.begin(), stones.end());
    while (stones.size() > 1)
    {
      int s1 = stones.back();
      stones.pop_back();
      int s2 = stones.back();
      stones.pop_back();

      if (s1 == s2)
        continue;

      stones.push_back(s1 - s2);

      sort(stones.begin(), stones.end());
    }
    return stones.size() == 0 ? 0 : stones.back();
  }
};

int main()
{
  vector<int> s{2, 7, 4, 1, 8, 1};
  cout << Solution().lastStoneWeight(s);
}