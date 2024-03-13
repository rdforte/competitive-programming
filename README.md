# Competitive Programming

Contains my answers to problems solved on variouse competitive programming sites

### [Code Template](./CODE_TEMPLATES.md)

### Setting up C++ in VSCode

[Using Clang in Visual Studio Code](https://code.visualstudio.com/docs/cpp/config-clang-mac)

### Using bits/stdc++.h

bits/stdc++ is a GNU GCC extension, whereas OSX uses the clang compiler. So we will not have access to this header.
In order to add this header we have to first find out where to put it.

run:

```
echo "" | gcc -xc - -v -E
```

This will give you a large sum of output but what we want is to find where it uses the #include
![finding include](./assets/findingInclude.png)

look for the one with `include` at the end.

Once we have this lets go into that directory and make a new directory called **bits**

```
cd /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/include
```

```
sudo mkdir bits
```

now go into bits and add [stdc++.h](https://github.com/gcc-mirror/gcc/blob/master/libstdc%2B%2B-v3/include/precompiled/stdc%2B%2B.h)

run the following command in the same directory as the stdc++.h file. Note that the path you are copying should match the 'include' path
you retrieved earlier.

```
sudo cp stdc++.h /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/include/bits/stdc++.h
```

note there where a few issues with the \_\_cplusplus preprocessor macro so a had to move a few of the libraries around but once I did
that it seemed to be working fine.

### Note

We use C++ 17 because LeetCode does not support C++ 20

## Tips

Always suss out your constraints. Example graph question constraints:

- `n == rooms.length`
- `2 <= n <= 1000` the array length is between 2 and 1000
- `0 <= rooms[i].length <= 1000` the adjacency list values have an array length between 0 and 1000
- `1 <= sum(rooms[i].length) <= 3000`
- `0 <= rooms[i][j] < n` the value in the adjacency list is always less than n
- `All the values of rooms[i] are unique.` each value in the adjacency list for the node is unique

### Stacks

A lot of parsing questions (valid parentheses, braces, parsing math expressions),
and basically anything where 'future data can impact earlier'.
Another pattern would be monotonically increasing / decreasing stack questions, where you eliminate all entries greater / less
than on each iteration.

### Dynamic Programming

Two common use cases:

- Finding an optimal solution ie: max or min
- Counting the total number of solutions for a problem

### Heaps / Priority Queue

- Good for finding the top **K** elements.
- Dijkstras algorithm for calculating the shortest path in weighted graphs.

### Sliding Window

- The condition for using the sliding window technique is the problem asks to find the `max/min` value for a function that calculates
  the answer repeatedly for a set of `ranges` from the array.
  Example Leetcode 1004: Given a binary array nums and an integer k, return the maximum number of `consecutive 1's` (consecutive 1's creates a range)
  in the array if you can flip at most k 0's.

### BFS (Breadth First Search)

- Great for finding the shortest path in the graph.
- Can be used with priority queue to form Dijkstras algorithm to find shortest path in weighted graph

### Binary Search

Although binary search algorithms are typically used to find one element in a sorted sequence, they have many other uses.
You can apply a binary search to a result, for example. Say you wanted to determine the minimum square footage of office space needed to fit all a company's employees easily.
The square footage is a consecutive number and therefore sorted so we can perform binary search on it
