#include <iostream>

using namespace std;

char characterBefore(char c)
{
  return (char)((int)c - 1);
}

int main()
{
  cout << characterBefore('c') << "\n";
  cout << characterBefore('b') << "\n";
  cout << characterBefore('a') << "\n";
}