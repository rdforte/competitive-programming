#include "../../../stdc++.h"

using namespace std;

void reverseString(vector<char> &s)
{
  for (int i = 0, j = s.size() - 1; i < j; i++)
  {
    swap(s[i], s[j]);
    j--;
  }
}

int main()
{
  vector<char> c{'h', 'e', 'l', 'l', 'o'};
  reverseString(c);

  for (auto i : c)
  {
    cout << i;
  }
}