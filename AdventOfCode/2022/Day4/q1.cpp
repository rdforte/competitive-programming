#include <bits/stdc++.h>

using namespace std;

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

  int contains = 0;
  string s;
  while (getline(cin, s))
  {
    int r1Min = 0, r1Max = 0, r2Min = 0, r2Max = 0;

    string tempNum = "";
    for (int i = 0; i < s.size(); i++)
    {
      if (isdigit(s[i]))
      {
        tempNum += s[i];
        if (i + 1 < s.size() && isdigit(s[i + 1]))
        {
          continue;
        }

        if (r1Min == 0)
        {
          r1Min = stoi(tempNum);
        }
        else if (r1Max == 0)
        {
          r1Max = stoi(tempNum);
        }
        else if (r2Min == 0)
        {
          r2Min = stoi(tempNum);
        }
        else
        {
          r2Max = stoi(tempNum);
        }
      }
      else
      {
        tempNum = "";
      }
    }

    if ((r1Min >= r2Min && r1Max <= r2Max) || (r2Min >= r1Min && r2Max <= r1Max))
    {
      contains++;
    }
  }

  cout << contains;
}