# Prims Algorithm

Prims algorithm is very much like Kruskal's algorithm as we can use it to figure
out the **Minimum Spanning Tree**.

The way in which we can go about doing this is by:
1. keeping track of the vertices we have visited.
2. only connecting vertices from the visited list to the unvisited list.
3. Start from an arbitary chosen point and mark it as visited.
4. Add the vertices edges to the min heap
5. Take the next edge off the min heap and check to see if you have visited
that vertex if you have continue to pop the edges off the heap until you find
a vertice you havn't visited. Once you find one you havn't visited, mark the vertice
as visited then add its edges to the heap. 
6. Repeat.

## Complexity Analysis
V represents the number of vertices, and E represents the number of edges.

Time: O(E⋅logV)
Space: O(V)