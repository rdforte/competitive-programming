#include <iostream>
#include <functional>

std::function<int(int)> makeLambda(int a)
{
  return [a](int b)
  { return a + b; };
}

int main()
{

  auto add5 = makeLambda(5);

  auto add10 = makeLambda(10);

  // makeLambda(5) on line 13 captures 5 and makeLambda(10) captures 10.
  // Then when we invoke the lambda it has already got these captured variables
  // which are assigned to a. Then we pass the number which is assigned to param
  // b.
  std::cout << (add5(10) == add10(5)); // 15 == 15
}