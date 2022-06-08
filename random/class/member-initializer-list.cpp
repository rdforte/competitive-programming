#include <iostream>
#include <cmath>
// #define PI 3.1416
using namespace std;

/**
 *One advantage of member initializer list is that it makes things compact.
  One no longer has to use assignment statements in the constructor.
  Another important point is that any const class members can’t be initialized
  inside a constructor and must be initialized using member initializer lists.
 */

class Sphere
{
private:
  const double density;
  double radius;

public:
  Sphere(double r) : radius(r), density(4.3)
  {
    // The following initialization wouldn't work, because density is a const
    // density =  4.3;
  }
  double volume()
  {
    return 4 * M_PI * radius * radius * radius / 3;
  }
  double mass()
  {
    return density * volume();
  }
};
int main()
{
  // your code goes here
  Sphere s(3.2);
  cout << "Volume: " << s.volume() << ", mass: " << s.mass();
  return 0;
}