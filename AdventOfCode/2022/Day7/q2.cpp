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
      v.push_back(dirSize);
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
  freopen("q2output.txt", "w", stdout);
  vector<int> v;
  int totalSpace = dfs("", v);

  if (totalSpace <= 40000000)
  {
    cout << 0 << "\n";
  }
  else
  {
    int dirSize = INT_MAX;
    for (auto size : v)
    {
      if (totalSpace - size <= 40000000)
      {
        dirSize = min(dirSize, size);
      }
    }
    cout << dirSize << "\n";
  }
}