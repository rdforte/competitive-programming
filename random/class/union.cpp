#include <iostream>

union Value
{
  int i;
  double d;
};

int main()
{
  Value v = {123};          // now v holds an int
  std::cout << v.i << '\n'; // write 123
  v.d = 987.654;            // now v holds a double
  std::cout << v.d << '\n'; // write 987.654
}