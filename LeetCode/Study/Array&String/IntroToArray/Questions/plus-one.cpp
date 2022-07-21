// https://leetcode.com/problems/plus-one/
#include <vector>
#include <iostream>

using namespace std;

class Solution
{
public:
  vector<int> plusOne(vector<int> &digits)
  {
    for (int i = digits.size() - 1; i >= 0; i--)
    {
      if (digits[i] < 9)
      {
        digits[i]++;
        return digits;
      }
      else
      {
        digits[i] = 0;
      }
    }

    digits.push_back(0);
    digits[0] = 1;

    return digits;
  }
};

int main()
{
  vector<int> digits{9};

  for (auto i : Solution().plusOne(digits))
  {
    cout << i << ", ";
  }
}