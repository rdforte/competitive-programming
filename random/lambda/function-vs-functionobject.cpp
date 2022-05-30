#include <iostream>
#include <iomanip>

int addFunc(int a, int b) { return a + b; };

int main()
{

  // the trailing const means:
  // The compiler will check to make sure that the method does not attemp to
  // modify the object.

  // An operator is something along the lines of  +, -
  // We can overide such operators to give new functionality.
  // In this case we are overriding the structs operator with our own custom
  // operator. This means that the object will now behave like a function as we
  // have overloaded it through the use of the operator.
  struct AddObj
  {
    int operator()(int a, int b) const { return a + b; };
  };

  AddObj addObj;

  std::cout << std::boolalpha;

  std::cout << (addObj(3, 4) == addFunc(3, 4)) << "\n";

  std::cout << "------------------------------------------------------------\n";

  // this behaves the same as AddObj.
  // because AddObj is essentialy an object which can hold state. so can
  // our lambda.
  auto addObj2 = [](int a, int b)
  { return a + b; };

  std::cout << (addObj2(3, 4) == addFunc(3, 4));
}
