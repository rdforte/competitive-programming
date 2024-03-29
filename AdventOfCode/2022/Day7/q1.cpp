#include <bits/stdc++.h>

using namespace std;

int dfs(string dir, vector<int> &v)
{
  string s;
  int totalSizeOfDir = 0;
  while (getline(cin, s))
  {
    if (s == "$ cd ..")
    {
      return totalSizeOfDir;
    }
    else if (regex_match(s, regex("\\$ cd [^\\.]+")))
    {
      int dirSize = dfs(s, v);
      totalSizeOfDir += dirSize;
      if (dirSize <= 100000)
      {
        v.push_back(dirSize);
      }
    }
    else if (s == "$ ls")
    {
      continue;
    }
    else
    {
      if (regex_match(s, regex("dir")))
      {
        continue;
      }
      else
      {
        smatch m;
        if (regex_search(s, m, regex("\\d+")))
        {
          totalSizeOfDir += stoi(m.str(0));
        };
      }
    }
  }

  return totalSizeOfDir;
}

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);
  vector<int> v;
  dfs("", v);

  int total = 0;
  for (auto size : v)
  {
    total += size;
  }
  cout << total;
}