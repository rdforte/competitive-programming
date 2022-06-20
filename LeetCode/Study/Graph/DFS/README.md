# DFS (Depth First Search)

__Depth-first search is a method to traverse a graph where the deepest not yet explored node is always explored first.__

Given a graph, how can we find all of its vertices, and how can we find all paths between two vertices?

The depth-first search algorithm is ideal in solving these kinds of problems because it can explore all paths from the start vertex to all other vertices. 

In Graph theory, the depth-first search algorithm (abbreviated as DFS) is mainly used to:

1. Traverse all vertices in a “graph”;
2. Traverse all paths between any two vertices in a “graph”.

### Cycle detection
We can use DFS to detect cycles in graphs. For this, we’ll start with DFS from an arbitrary vertex. After finishing that vertex, we’ll continue with DFS from a remaining white vertex, until all vertices are finished (blue).

If we find a back edge during the execution of DFS that leads from the current node to another node that is still in progress (gray), we’ve found a cycle in the graph. This indicates that if a node is still in progress, there must be a path from it to the current node, and the back edge thus completes a cycle. If such a back edge is never discovered, the graph is acyclic.

### Prefered way to display graph information
We can display graph information using an __adjacency list__ where
the keys are nodes/vertices and the value is a list containing the neighbouring nodes for that given key(node).
```
  unordered_map<string, vector<string>> graph{
      {"a", {"c", "b"}},
      {"b", {"d"}},
      {"c", {"e"}}
      {"d", {"f"}},
      {"e", {}},
      {"f", {}},
  };
```
_For the above if the vertices are not strings then we can use a vector instead of an unordered_map and refer to the index as the vertice._ 

### NOTES
DFS is utilised using a stack. This can be done through creating our own stack along with an interative approach or we can utilise the call stack with a recursive approach.

When implementing DFS we can use a stack.
[std::stack](https://en.cppreference.com/w/cpp/container/stack)
