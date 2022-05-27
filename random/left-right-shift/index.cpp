#include <iostream>
#include <iomanip> // can output boolean true false
#include <cmath>   // gives us pow

using namespace std;

int main()
{
  cout << boolalpha;

  bool isTheSame = (2 << 3) == (2 * pow(2, 3));

  cout << isTheSame;
}