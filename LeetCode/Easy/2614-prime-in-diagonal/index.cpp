#include "../../../stdc++.h"

using namespace std;

vector<bool> getPrimes()
{
  int max = 4e6 + 1;
  vector<bool> primes(max, true);
  primes[0] = false;
  primes[1] = false;

  for (int i = 2; i < max; i++)
  {
    if (!primes[i])
      continue;
    for (int j = i * 2; j < max; j += i)
      primes[j] = false;
  }

  return primes;
}

vector<bool> primes = getPrimes();

class Solution
{
public:
  int diagonalPrime(vector<vector<int>> &nums)
  {
    int res = 0;

    for (int i = 0; i < nums.size(); i++)
    {
      if (primes[nums[i][i]])
        res = max(res, nums[i][i]);
      if (primes[nums[i][nums.size() - i - 1]])
        res = max(res, nums[i][nums.size() - i - 1]);
    }

    return res;
  }
};

class Solution2
{
public:
  int diagonalPrime(vector<vector<int>> &nums)
  {
    int res = 0;

    for (int i = 0; i < nums.size(); i++)
    {
      if (checkIsPrime(nums[i][i]))
        res = max(res, nums[i][i]);
      if (checkIsPrime(nums[i][nums.size() - i - 1]))
        res = max(res, nums[i][nums.size() - i - 1]);
    }

    return res;
  }

private:
  bool checkIsPrime(int num)
  {
    if (num <= 1)
      return false;
    for (int i = 2; i * i <= num; i++)
    {
      if (num % i == 0)
        return false;
    }
    return true;
  }
};

int main()
{
}