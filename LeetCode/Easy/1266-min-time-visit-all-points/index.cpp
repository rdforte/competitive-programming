// https://leetcode.com/problems/minimum-time-visiting-all-points/
#include <bits/stdc++.h>

using namespace std;

int minTimeToVisitAllPoints(vector<vector<int>> &points)
{
  if (points.size() == 1)
    return 0;
  int ans = 0;
  for (int i = 1; i < points.size(); i++)
  {
    int x = abs(points[i][0] - points[i - 1][0]);
    int y = abs(points[i][1] - points[i - 1][1]);
    ans += min(x, y) + abs(x - y);
  }
  return ans;
};
