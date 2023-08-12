if n = number of nodes
m = number of edges

Because it is bidirectional we iterate through max m edges and use our priority queue to do sorting when an edge is
added
to the queue. so this would be m*log*n.
We also need to iterate through all of our nodes. Therefore, the final time complexity is n + m*log*n.