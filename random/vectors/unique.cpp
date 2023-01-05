#include <bits/stdc++.h>

using namespace std;

int main()
{
  // elements need to be in order for unique to work
  vector<int> v = {1, 2, 2, 2, 3, 4, 4, 5};

  // points to the last element that follows the element which was not removed.
  auto it = unique(v.begin(), v.end());

  v.erase(it, v.end()); // need to resize the array

  for (auto i : v)
  {
    cout << i << ", ";
  }
}