#include <iostream>

class Solution
{
public:
  int countOperations(int num1, int num2)
  {
    return recursiveOperations(num1, num2);
  }

private:
  int recursiveOperations(int num1, int num2)
  {

    if (num1 <= 0 || num2 <= 0)
      return 0;

    if (num1 >= num2)
    {
      return recursiveOperations(num1 - num2, num2) + 1;
    }
    else
    {
      return recursiveOperations(num1, num2 - num1) + 1;
    }
  }
};

int main()
{
  std::cout << Solution().countOperations(10, 10);
}