// https://leetcode.com/problems/maximum-number-of-groups-entering-a-competition/
#include <iostream>
#include <queue>
#include <stack>
#include <vector>
#include <set>
#include <algorithm>
#include <utility>

using namespace std;

// To maximize the number of groups we need to form groups 1, 2, 3, 4 ...
// The sum of the grades will be satisfied if the grades are sorted.
// Therefore we only care about the size of the grades array and not the grades themselves.

// The number of students in a group is

class Solution
{
public:
  int maximumGroups(vector<int> &grades)
  {
  }
};

int main()
{
  vector<int> groups{
      10, 6, 12, 7, 3, 5};

  cout << Solution().maximumGroups(groups);
}