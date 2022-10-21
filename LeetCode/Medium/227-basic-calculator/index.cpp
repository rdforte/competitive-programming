#include "../../../stdc++.h"

using namespace std;

int calculator(string s)
{
  int prev = 0;
  int res = 0;
  char operand = '+';

  for (int i = 0; i < s.size(); i++)
  {
    cout << prev << operand << res << "\n";
    if (isdigit(s[i]))
    {
      string cur = {s[i]};

      while (isdigit(s[++i]))
      {
        cur += s[i];
      }

      if (operand == '+')
      {
        res += stoi(cur);
        prev = stoi(cur);
      }

      if (operand == '-')
      {
        res -= stoi(cur);
        prev = -stoi(cur);
      }

      if (operand == '*')
      {
        int diff = res - prev;
        res = (stoi(cur) * prev) + diff;
        prev = (stoi(cur) * prev);
      }

      if (operand == '/')
      {
        int diff = res - prev;
        res = (prev / stoi(cur)) + diff;
        prev = (prev / stoi(cur));
      }
    }

    if (s[i] == ' ')
      continue;

    operand = s[i];
  }

  return res;
}

int main()
{
  cout << calculator("12-3*4");
}