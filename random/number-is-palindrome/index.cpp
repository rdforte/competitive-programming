#include <iostream>
#include <string>
#include <cmath>

using namespace std;

int main()
{
  int number = 123421;
  string n = to_string(number);

  bool isPalindrone = true;

  for (int i = 0; i < ceil(((float)n.size()) / 2); i++)
  {
    if (n[i] != n[n.size() - 1 - i])
    {
      isPalindrone = false;
      break;
    }
  }

  cout << (isPalindrone ? "is palindrome" : "not a palindrome");
}