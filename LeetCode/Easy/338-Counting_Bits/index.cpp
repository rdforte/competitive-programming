#include <iostream>
#include <vector>

using namespace std;

class Solution
{
public:
  vector<int> countBits(int n)
  {
    vector<int> numBits{};

    for (int i = 0; i <= n; i++)
    {
      numBits.push_back(calculateNumberOfBits(i));
    }

    return numBits;
  }

  int calculateNumberOfBits(int n)
  {
    u_int mask = 1;
    int bits = 0;

    for (int i = 0; i < 32; i++)
    {
      if ((n & mask) != 0)
        bits++;

      mask <<= 1;
    }
    return bits;
  }
};

class SolutionDynamicProgramming
{
public:
  vector<int> countBits(int n)
  {
    if (n == 0)
      return {0};
    if (n == 1)
      return {0, 1};

    vector<int> numBits{0, 1};
    u_int bitNumber = 2;

    for (int i = 2; i <= n; i++)
    {
      if (bitNumber == i)
      {
        numBits.push_back(1);
        bitNumber <<= 1;
      }
      else
      {
        int prevCalculated = numBits[i - (bitNumber >> 1)];
        numBits.push_back(prevCalculated + 1);
      }
    }

    return numBits;
  }
};

int main()
{
  vector<int> sol = SolutionDynamicProgramming().countBits(8);

  for (auto s : sol)
  {
    cout << s << ", ";
  }
}