#include <bits/stdc++.h>

using namespace std;

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q2output.txt", "w", stdout);

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
    if (packets.size() == 14)
    {
      cout << i + 1;
      break;
    }
  }
}