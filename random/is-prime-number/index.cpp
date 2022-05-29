#include <iostream>
#include <iomanip>

using namespace std;

int main()
{
  int number = 31;

  cout << boolalpha;

  if (number == 1 || number == 0)
  {
    cout << false;
  }
  else if (number == 2)
  {
    cout << true;
  }
  else
  {
    bool isPrime = true;
    for (int i = 3; i < number; i++)
    {
      if (number % i == 0)
      {
        isPrime = false;
        break;
      }
    }
    cout << isPrime;
  }
}