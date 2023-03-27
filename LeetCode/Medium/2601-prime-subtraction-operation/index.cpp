#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
  bool primeSubOperation(vector<int> &nums)
  {
    vector<bool> primes = seiveAlgo();

    for (int i = 0; i < nums.size(); i++)
    {
      int num = nums[i];
      for (int p = 2; p < nums[i]; p++)
      {

        if (!primes[p])
          continue;

        if ((i > 0 && nums[i] - p > nums[i - 1]) || i == 0)
        {
          num = nums[i] - p;
        }
      }
      nums[i] = num;
      if (i > 0 && nums[i] <= nums[i - 1])
      {
        return false;
      }
    }

    return true;
  }

  vector<bool> seiveAlgo()
  {
    vector<bool> primes(1001, true);
    primes[0] = false;
    primes[1] = false;
    for (int i = 2; i < primes.size(); i++)
      if (primes[i])
        for (int p = i * 2; p < primes.size(); p += i)
          primes[p] = false;
    return primes;
  }
};

int main()
{
}
