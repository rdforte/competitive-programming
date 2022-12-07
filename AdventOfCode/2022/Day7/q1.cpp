#include <bits/stdc++.h>

using namespace std;

void dfs(string dir)
{
  string s;
  while (getline(cin, s))
  {
    if (s == "$ cd ..")
    {
      cout << "(" + dir + ")";
      cout << s << "\n";
      return;
    }
    else if (regex_match(s, regex("\\$ cd [^\\.]+")))
    {
      cout << "(" + dir + ")";
      cout << "cd dir: " << s << "\n";
      dfs(s);
    }
    else if (s == "$ ls")
    {
      cout << "(" + dir + ")";
      cout << "list dir: " << s << "\n";
      continue;
    }
    else
    {
      cout << "(" + dir + ")";
      cout << "file or dir: " << s << "\n";
    }
  }
}

int main()
{
  freopen("test.txt", "r", stdin);
  dfs("");
}