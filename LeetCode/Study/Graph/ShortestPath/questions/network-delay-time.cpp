// https://leetcode.com/problems/network-delay-time/
#include <iostream>
#include <vector>
#include <limits>
#include <queue>
#include <utility>
#include <set>
#include <algorithm>

using namespace std;

class Node
{
public:
  int time;
  int destination;

  Node(int d, int t)
  {
    destination = d;
    time = t;
  }

  void increaseTimeBy(int t)
  {
    time += t;
  }
};

class Solution
{
public:
  int networkDelayTime(vector<vector<int>> &times, int n, int k)
  {

    // Min Heap
    auto cmp = [](Node left, Node right)
    { return left.time > right.time; };
    priority_queue<Node, vector<Node>, decltype(cmp)> minHeap(cmp);

    // Adjacency List
    vector<vector<Node>> adjacencyList(n + 1);

    // Keep track of visited vertices
    set<int> visited;

    // Keep track of time
    int time = 0;

    // build adjacency list
    for (auto t : times)
    {
      int source = t[0];
      int destination = t[1];
      int time = t[2];

      Node n = Node(destination, time);
      adjacencyList[source].push_back(n);
    }

    // add the initial source to the min heap (starting point)
    Node source = Node(k, 0);
    minHeap.push(source);

    // Iterate through min Heap
    while (!minHeap.empty())
    {
      Node node = minHeap.top();
      minHeap.pop();

      // if already visited destination then don't add all of its neighbours again to the heap.
      if (visited.find(node.destination) != visited.end())
        continue;

      visited.insert(node.destination);

      time = max(time, node.time);

      for (auto neighbour : adjacencyList[node.destination])
      {
        if (visited.find(neighbour.destination) == visited.end())
        {
          // when we add the destination to the queue add the previouse time of the source ndoe to the neighbours weight
          // so that we have the total time it takes to reach that node.
          neighbour.increaseTimeBy(node.time);
          minHeap.push(neighbour);
        }
      }
    }

    // if we
    //   have visited all vertices then the sets length should equal n if we have not visited all vertices then return -1 cout << time << "\n";
    return visited.size() == n ? time : -1;
  }
};

int main()
{
  // vector<vector<int>> times{
  //     {2, 1, 1},
  //     {2, 3, 1},
  //     {3, 4, 1},
  // };
  // int numVertices = 4;
  // int source = 2;

  vector<vector<int>> times{
      {1, 2, 1},
  };
  int numVertices = 2;
  int source = 1;
  cout << Solution().networkDelayTime(times, numVertices, source);
}