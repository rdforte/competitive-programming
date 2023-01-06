#include <bits/stdc++.h>

using namespace std;

bool isBadVersion(int version)
{
  return version >= 3;
}

class Solution
{
public:
  int firstBadVersion(int n)
  {
    return recursiveCheckVersion(1, n);
  }

private:
  int recursiveCheckVersion(int left, int right)
  {
    int middle = ((long)left + (long)right) / 2;
    if (isBadVersion(middle))
      return !isBadVersion(middle - 1) ? middle : recursiveCheckVersion(left, middle - 1);
    else
      return isBadVersion(middle + 1) ? middle + 1 : recursiveCheckVersion(middle + 1, right);
  }
};

int main()
{
  cout << Solution().firstBadVersion(5);
}