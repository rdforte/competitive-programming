# Breadth First Search
Breadth first search (BFS) can traverse all vertices in a graph and all paths
between two vertices.

**DFS** can find the shortest path between two vertices providing that
the graph is balanced. Though it must traverse all paths.

**BFS** on the other case can find the shortest path without traversing all
paths. This is because when the path is found between the source and target vertex
it is guaranteed to be the shortest.

**In Graph theory, the primary use cases of the BFS algorithm are:**
1. Traversing all vertices in a graph.
2. Finding the shortest path between two vertices where all edges have equal
and positive weights.

DFS = stack (first in last out) [std::stack](https://en.cppreference.com/w/cpp/container/stack)
BFS = queue (first in first out) [std::queue](https://en.cppreference.com/w/cpp/container/queue)