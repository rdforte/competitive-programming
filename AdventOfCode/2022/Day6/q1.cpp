#include <bits/stdc++.h>

using namespace std;

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

  string s;
  getline(cin, s);
  unordered_set<char> packets;

  int j = 0;
  for (int i = 0; i < s.size(); i++)
  {
    if (packets.find(s[i]) != packets.end())
    {
      packets.clear();
      i = j;
      j++;
      continue;
    }
    packets.insert(s[i]);
    if (packets.size() == 4)
    {
      cout << i + 1;
      break;
    }
  }
}