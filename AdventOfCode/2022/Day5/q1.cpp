#include <bits/stdc++.h>

using namespace std;

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

  vector<stack<char>> crates;
  for (int i = 0; i < 9; i++)
  {
    stack<char> st;
    crates.push_back(st);
  }
  vector<vector<char>> tempCrates(9);

  string c;
  int lineNum = 0;
  while (getline(cin, c))
  {
    lineNum++;
    for (int i = 0; i < c.size(); i++)
    {
      if ((i - 1) % 4 == 0 && c[i] != ' ')
      {
        tempCrates[(i - 1) / 4].push_back(c[i]);
      }
    }
    if (lineNum == 8)
      break;
  }

  for (int i = 0; i < 9; i++)
  {
    while (!tempCrates[i].empty())
    {
      crates[i].push(tempCrates[i].back());
      tempCrates[i].pop_back();
    }
  }

  getline(cin, c);
  getline(cin, c);

  int amount = -1, from = -1, to = -1;
  while (getline(cin, c))
  {
    string temp = "";
    for (int i = 0; i < c.size(); i++)
    {
      if (!isdigit(c[i]))
        temp = "";

      if (isdigit(c[i]))
      {
        temp += c[i];
        if (i + 1 < c.size() - 1 && isdigit(c[i + 1]))
          continue;

        if (amount == -1)
        {
          amount = stoi(temp);
        }
        else if (from == -1)
        {
          from = (stoi(temp) - 1);
        }
        else
        {
          to = (stoi(temp) - 1);
        }
      }
    }

    while (amount > 0)
    {
      amount--;
      char crate = crates[from].top();
      crates[from].pop();
      crates[to].push(crate);
    }
    amount = -1;
    from = -1;
    to = -1;
  }

  for (auto individualCrate : crates)
  {
    cout << individualCrate.top();
  }
}