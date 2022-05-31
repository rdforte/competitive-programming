#include <iostream>

using namespace std;

// Recursive factorial function
int factorial(int n)
{
  // Invalid value
  if (n < 0)
  {
    return -1;
  }
  // Base case
  if (n == 1 || n == 0)
  {
    return 1;
  }
  // Recursive Case
  else
  {
    return n * factorial(n - 1);
  }
}

// main function
int main()
{
  int n;
  cout << "which number would you like to know the factorial for? ";
  cin >> n;

  int result;
  // Call factorial function in main and store the returned value in result
  result = factorial(n);
  // Prints value of result
  cout << "Factorial of " << n << " = " << result;
  return 0;
}