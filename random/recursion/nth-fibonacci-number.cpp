#include <iostream>
#include <unordered_map>

int fibonacci(int n)
{
  if (n == 0)
    return 0;
  if (n == 1)
    return 1;

  return fibonacci(n - 1) + fibonacci(n - 2);
};

int fibonacciMemoized(int n, std::unordered_map<int, int> &store)
{
  if (n == 0)
    return 0;
  if (n == 1)
    return 1;

  if (store.find(n) != store.end())
  {
    return store[n];
  }
  else
  {
    store[n] = fibonacciMemoized(n - 1, store) + fibonacciMemoized(n - 2, store);
    return store[n];
  }
}

int main()
{
  int fib;
  std::cout << "fibonacci number? ";
  std::cin >> fib;

  std::unordered_map<int, int> s{};
  std::cout << fibonacciMemoized(fib, s);
}