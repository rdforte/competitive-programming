// https://leetcode.com/problems/course-schedule-ii/
#include <iostream>
#include <vector>
#include <utility>
#include <queue>

using namespace std;

class Solution
{
public:
  vector<int> findOrder(int numCourses, vector<vector<int>> &prerequisites)
  {
    vector<int> adj[numCourses]; // first - neighbour, second - inDegree.
    vector<int> inDegree(numCourses, 0);
    vector<int> schedule;

    queue<int> q;

    // build adjacency list and update the inDegree of the destination.
    for (auto p : prerequisites)
    {
      adj[p[1]].push_back(p[0]);
      inDegree[p[0]]++;
    }

    // find source node with in-degree 0
    for (int i = 0; i < inDegree.size(); i++)
    {
      if (inDegree[i] == 0)
      {
        q.push(i);
      }
    }

    while (!q.empty())
    {
      int node = q.front();
      q.pop();
      schedule.push_back(node);

      // for each neighbour of the node decrease its inDegree by 1.
      // if its inDegree is 0 then lets add it to the queue.
      for (auto neighbour : adj[node])
      {
        inDegree[neighbour]--;
        if (inDegree[neighbour] == 0)
          q.push(neighbour);
      }
    }

    if (schedule.size() != numCourses)
      return {};

    return schedule;
  }
};

int main()
{
  // 0, 1, 2, 3 or 0, 2, 1, 3
  // int numCourses = 4;
  // vector<vector<int>> prerequisites = {
  //     {1, 0},
  //     {2, 0},
  //     {3, 1},
  //     {3, 2},
  // };

  // 0, 1
  // int numCourses = 2;
  // vector<vector<int>> prerequisites = {};

  // {}
  int numCourses = 3;
  vector<vector<int>> prerequisites = {
      {1, 0},
      {1, 2},
      {0, 1},
  };

  vector<int> schedule = Solution().findOrder(numCourses, prerequisites);

  for (auto s : schedule)
  {
    cout << s << " ";
  }
}