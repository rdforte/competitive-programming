#include <string>
#include <iostream>

using namespace std;

int main()
{
  string s("hello world");

  // holds a reference to the first char of string
  auto firstLetter = s.begin();
  // Change the value of first letter.
  // s.begin() returns an iterator which is a reference.
  *firstLetter = toupper(*firstLetter);

  cout << s;
}