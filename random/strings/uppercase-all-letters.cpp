#include <string>
#include <iostream>

using namespace std;

int main()
{
  string s("hello world");

  for (auto it = s.begin(); it != s.end(); it++)
  {
    if (!isspace(*it))
      *it = toupper(*it);
  }

  cout << s;
}