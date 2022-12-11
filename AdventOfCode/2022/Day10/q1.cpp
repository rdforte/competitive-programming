#include <bits/stdc++.h>

using namespace std;

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

  queue<int> q;
  string s;
  int total = 0;
  int iterationNum = 0;
  int X = 1;
  int cycle = 20;
  while (getline(cin, s) || iterationNum < 220)
  {
    iterationNum++;
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
      total += (X * cycle);
      cycle += 40;
    }

    X += q.front();
    q.pop();
  }

  cout << total;
}