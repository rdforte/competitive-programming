// https://leetcode.com/explore/learn/card/graph/619/depth-first-search-in-graph/3893/

#include <iostream>
#include <vector>
#include <stack>

using namespace std;

/**
 * Time Complexity O(V + E) where V is the Vertices and E is the edges
 * To create the adjacency list, we must iterate over each of the E edges.
 * In the while loop, at most, we will visit each Vertex once.
 * In the while loop we will have a total sum of at most E iterations as we
 * get each of the nodes neighbours.
 *
 * E + V + E
 * 2E + V
 * E + V
 *
 * Space Complexity O(V + E)
 * The adjacency list will contain each Vertice + each Edge
 * The stack will contain each Vertex
 * The Seen will be each Vertex
 *
 * V + E + E + V
 * 2V + 2E
 * V + E
 */

class Solution
{
public:
  bool validPath(int n, vector<vector<int>> &edges, int start, int end)
  {
    vector<vector<int>> adjacency_list(n);
    for (vector<int> edge : edges)
    {
      adjacency_list[edge[0]].push_back(edge[1]);
      adjacency_list[edge[1]].push_back(edge[0]);
    }

    stack<int> st;
    st.push(start);
    vector<bool> seen(n); // keep track of seen nodes

    while (!st.empty())
    {
      // Get the current node.
      int node = st.top();
      st.pop();

      // Check if we have reached the target node.
      if (node == end)
      {
        return true;
      }

      // Check if we've already visited this node.
      if (seen[node])
      {
        continue;
      }
      seen[node] = true;

      // Add all neighbors to the stack.
      for (int neighbor : adjacency_list[node])
      {
        st.push(neighbor);
      }
    }

    return false;
  }
};

int main()
{
  cout << boolalpha;
  int n = 10;
  vector<vector<int>> edges{
      {0, 7},
      {0, 8},
      {6, 1},
      {2, 0},
      {0, 4},
      {5, 8},
      {4, 7},
      {1, 3},
      {3, 5},
      {6, 5},
  };
  int source = 7;
  int destination = 5;

  vector<int> v{1};
  v.erase(next(v.begin(), 0));

  cout << Solution().validPath(n, edges, source, destination);
}