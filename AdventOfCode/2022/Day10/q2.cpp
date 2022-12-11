#include <bits/stdc++.h>

using namespace std;

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q2output.txt", "w", stdout);

  string picReset = "........................................";

  queue<int> q;
  string s;
  int iterationNum = 0;
  int pointer = 0;
  int X = 1;
  int cycle = 40;
  string pic = picReset;
  while (getline(cin, s) || iterationNum < 240)
  {
    iterationNum++;

    if (pointer >= (X - 1) and pointer <= (X + 1))
    {
      pic[pointer] = '#';
    }
    pointer++;

    smatch nextAddition;
    if (regex_match(s, regex("noop")))
    {
      q.push(0);
    }
    else if (regex_search(s, nextAddition, regex("-?\\d+")))
    {
      q.push(0);
      q.push(stoi(nextAddition.str()));
    }

    if (iterationNum == cycle)
    {
      cout << pic << "\n";
      pic = picReset;
      pointer = 0;
      cycle += 40;
    }

    X += q.front();
    q.pop();
  }
}