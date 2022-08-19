#include <iostream>

using namespace std;

class Solution
{
public:
  int hammingWeight(uint32_t n)
  {
    int numBits = 0;
    u_int mask = 1;

    for (int i = 0; i < 32; i++)
    {
      if ((n & mask) != 0)
        numBits++;
      mask <<= 1;
    }

    return numBits;
  }
};

int main()
{
  // We are using actual numbers though when we use bitwise operators it
  // treats the numbers as bits.
  // the mask starts at 1 and then increments to 2,4,8,16,32 ...
  // though in binary this looks like, 1, 10, 100, 1000, 10000
  // so we can use this as a mask.

  cout << Solution().hammingWeight(11); // 3
}
