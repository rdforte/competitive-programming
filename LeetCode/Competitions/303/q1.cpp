#include <iostream>
#include <set>
#include <string>

using namespace std;

class Solution
{
public:
  char repeatedCharacter(string s)
  {
    set<char> charSet;

    char letter = s[0];

    for (auto c : s)
    {
      if (charSet.find(c) != charSet.end())
      {
        letter = c;
        break;
      }
      else
      {
        charSet.insert(c);
      }
    }
    return letter;
  }
};
