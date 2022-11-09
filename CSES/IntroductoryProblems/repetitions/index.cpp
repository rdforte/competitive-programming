// https://cses.fi/problemset/task/1069
#include <bits/stdc++.h>

using namespace std;

int main()
{
  string dna;
  cin >> dna;

  char prev;
  int count = 0;
  int longest = 0;

  for (auto c : dna)
  {
    if (c != prev)
    {
      count = 1;
      prev = c;
    }
    else
    {
      count++;
    }
    longest = max(count, longest);
  }
  cout << longest;
}