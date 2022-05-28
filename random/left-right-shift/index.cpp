#include <iostream>
#include <iomanip> // can output boolean true false
#include <cmath>   // gives us pow

using namespace std;

int main()
{
  cout << boolalpha;

  bool isTheSame = (6 << 3) == (6 * pow(2, 3));

  cout << isTheSame;
}