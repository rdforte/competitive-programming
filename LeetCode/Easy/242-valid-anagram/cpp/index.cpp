#include <string>

class Solution
{
public:
  bool isAnagram(std::string s, std::string t)
  {
    sort(s.begin(), s.end());
    sort(t.begin(), t.end());

    return s == t;
  }
};