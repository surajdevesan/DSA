# Shortest Path Algos
These algos find the shortest path between two points in a graph. There are many different algos in place. Some common ones are

- [Dijkstra's Algorithm](#dijkstras-algorithm)

## Dijkstra's Algorithm
There are many variants of Dijkstra's algo, we here talk about the one which produces shortest path to all vertices. It generates SPT(shortest path tree) and finds the shortest path from the sources to all vertices

### Algorithm
 - Use 2 sets dist set which tracks the distances of the vertices from the source and a sptSet(shortest path tree set) that keep track of vertices which are included in the shortest path tree.
 - Initialize both the set dist set with infinity or max int value and sptSet with false as initially none of the nodes are in the shortest path tree set.
 - Distance of the source is initialized as 0
 - Find a min distance point from the source
 - Add the point to the sptSet
 - Update dist set if distance from the new sptSet point is less than the current distance.
 - Repeat the iteration for all the points


## References
- https://www.geeksforgeeks.org/greedy-algorithms-set-6-dijkstras-shortest-path-algorithm/




