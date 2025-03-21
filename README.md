![cplogo](assets/cplogo.png)

![](https://leetcard.jacoblin.cool/rdforte?border=0&radius=20)

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

---

---

---

---

---

---

---

---

# DSA Algorithms and Tips

### Always identify constraints

Always suss out your constraints. Example graph question constraints:

- `n == rooms.length`
- `2 <= n <= 1000` the array length is between 2 and 1000
- `0 <= rooms[i].length <= 1000` the adjacency list values have an array length between 0 and 1000
- `1 <= sum(rooms[i].length) <= 3000`
- `0 <= rooms[i][j] < n` the value in the adjacency list is always less than n
- `All the values of rooms[i] are unique.` each value in the adjacency list for the node is unique

---

### Stacks

A lot of parsing questions (valid parentheses, braces, parsing math expressions),
and basically anything where 'future data can impact earlier'.
Another pattern would be monotonically increasing / decreasing stack questions, where you eliminate all entriesgreater / less
than on each iteration.

---

### Dynamic Programming

Two common use cases:

- Finding an optimal solution ie: max or min.
- Counting the total number of solutions for a problem.

**NOTE:** That if the question asks for the max/min of something or along those lines and involves consecutive/continguous items in
a list it may be a sliding window question. In which case refer to sliding window below.

examples may include:

- What is the minimum cost of doing...
- What is the maximum profit from...
- How many ways are there to do...
- What is the longest possible...
- Is it possible to reach a certain point...

**2D vs 1D**
Identifying a 2D dynamic programming problem can be tricky but with DP especially top down it is all about memoization.
If I have a question where im trying to see if I can make 'X' give items 'Y' and I can reuse items 'Y' as many times as I like then why bother
keeping track of 'Y' because 'Y' is not important in this case because I can just reuse it as much as I like because in the end all I care about
is if I can make 'X' from 'Y'. This is a good example of 1D because no need to keep track of Y's state.

If on the other hand I can not reuse 'Y' and I can only use these items once well now I need a way to keep track of what I have/havn't used. In this
case I need to keep track of Y's state. This is a good case for a 2D dynamic programming question.

So in the end ask yourself "do i need to keep track of theses items state?", if not then 1D otherwise 2D.
The state you need to track determines the dimension of the array.

- [322. Coin Change 1D question](LeetCode/Medium/322-coin-change/go/main.go)

Dynamic Programming can be performed in bottom-up (iterative) approach or top-down (recursive with memoization). If not sure how to perform
memoization for top-down think about memoizing the function call which in our case would be the arguments provided to the function at that time
of calling it. A good example of how I did this here [1770. Maximum Score from Performing Multiplication Operations](LeetCode/Hard/1770-maximum-score-from-performing-multiplication-operations/go/main.go)

**A good example of top down converted to bottom up for a 3D DP question**: [3418. Maximum Amount of Money Robot Can Earn](LeetCode/Medium/3418-maximum-amount-of-money-robot-can-earn/go/main.go)

**Sometimes** (i say sometimes because that may not always be the case) it can be easier to to do the top-down approach first and then translate this over to bottom-up for improved performance. A good example of this
is:

- [1143 Longest Common Subsequence](LeetCode/Medium/1143-longest-common-subsequence/go/main.go)
- [188 Best Time to Buy and Sell Stock IV](LeetCode/Hard/188-best-time-to-buy-and-sell-stock-iv/go/main.go)
- [714 Best Time to Buy and Sell Stock with Transaction Fee](LeetCode/Medium/714-best-time-to-buy-and-sell-stock-with-transaction-fee/)

cases where it was easier to perform bottom-up vs top-down:

- [718. Maximum Length of Repeated Subarray](LeetCode/Medium/718-maximum-length-of-repeated-subarray/go/main.go)
<p align="center">
  <img width="50%" height="50%" src="LeetCode/Medium/718-maximum-length-of-repeated-subarray/assets/solution.png" />
</p>
**State Reduction** is a way in which we can reduce the number of states which will in turn reduce our space complexity and possibly our time complexity.
If you look at [House Robber](LeetCode/Medium/198-house-robber/go/main.go) especially the Top Down recursive approach we could have include another parameter which was
a boolean to indicate if we could/couldn't rob the house but instead we used the index's which reduced the state.

**Note: state reductions for space complexity usually only apply to bottom-up implementations, while improving time complexity by reducing the number of state variables applies to both implementations.**

Another common scenario for improving space complexity is when the recurence relation is static (no iteration) along one dimension (flat array). An example of this is
[nth fibonacci number](LeetCode/Easy/509-fibonacci-number/go/main.go). Because we only care about the previouse 2 fibonacci numbers there is no need to keep an array
holding all previouse fibonacci numbers.

---

### Dynamic Programming (Kadane's Algorithm)

Kadane's Algorithm is an algorith that can find the maximum sum subarray when given an array of numbers. Its time complexity is O(n) and space O(1).

```
// Given an input array of numbers "nums",
1. best = nums[0]
2. current = nums[0]
3. for num in nums[1:]: // start at index 1
    3.1. current = Max(current + num, num)
    3.2. best = Max(best, current)

4. return best
```

- [53. Maximum Subarray](LeetCode/Medium/53-maximum-subarray/go/main.go)
- [121. Best Time to Buy and Sell Stock](LeetCode/Easy/121-best-time-buy-sell-stock/go/main.go)
- [918. Maximum Sum Circular Subarray](LeetCode/Medium/918-maximum-sum-circular-subarray/go/main.go)

![kadanes algorithm](assets/kadanes_algorithm.gif)

---

### Heaps / Priority Queue

- Good for finding the top **K** elements.
- Dijkstras algorithm for calculating the shortest path in weighted graphs.
- Good for keeping track of the max/min without having to sort the full list [2043 Stock Price Fluctuation](LeetCode/Medium/2034-stock-price-fluctuation/go/main.go)

---

### Sliding Window / 2 pointers

- The condition for using the sliding window technique is the problem asks to find the `max/min` value for a function that calculates
  the answer repeatedly for a set of `ranges` from the array.
  Example Leetcode 1004: Given a binary array nums and an integer k, return the maximum number of `consecutive 1's` (consecutive 1's creates a range)
  in the array if you can flip at most k 0's.
- Another good candidate for sliding window algorith is **SUB-ARRAY**. Because it is a contiguous and the order matters we can use sliding windown of DP.
- Sliding Window is a good candidate for **CONTIGUOUS** elements.
- Sliding Window also goes hand in hand with a hash map in the case where we want to keep track of the elements in the winddow.
  ex: [904 Fruits into baskets](LeetCode/Medium/904-fruits-into-baskets/go/main.go).
- **TIP**: I've found that when the brute force solution looks like a loop within a loop (**NESTED LOOPS**) which results in a n^2 solution, this may be a good candidate for a
  sliding window.
- The term subarray/contiguous elements of length k is equivalent to window of length k.

(1604. Alert Using Same Key-Card Three or More Times in a One Hour Period)[LeetCode/Medium/1604-alert-using-same-key-card-three-or-more-times-in-a-one-hour-perfiod/]
is a great sliding window question because it does not appear to be a sliding window question at first glance. This question gave hours and minutes. The question became
a lot easier once we converted hours and minutes to just minutes so we had a range between 0 - 1439 (23 \* 60 + 59). We could then see if any time was outside the 60 min
range.

---

### BFS (Breadth First Search)

- Great for finding the shortest path in the graph.
- Involves the use of a queue.
- Can be used with priority queue to form Dijkstras algorithm to find shortest path in weighted graph

---

### DFS (Depth First Search)

- Great for finding the cycle in a graph
  [207. Course Schedule](LeetCode/Medium/207-course-schedule/go/main.go)
- Involves the use of a stack or recursion (stack frame).

---

### Kahn's Algorithm for Topological Sorting

When selecting courses for the next semester in college, you might have noticed that some advanced courses have prerequisites that require you to take some
introductory courses first. For example, to take Course C, you need to complete Course B first, and to take Course B, you need to complete Course A first.
There are many courses that you must complete for an academic degree.
You do not want to find out in the last semester that you have not completed some prerequisite courses for an advanced course.
So, how can we arrange the order of the courses adequately while considering these prerequisite relationships between them?

[Course A] --> [Course B] --> [Course C]

_In order to take course C we need to take course B and to take course B we need to first take course A._

“Topological sorting” helps solve the problem.
**It provides a linear sorting based on the required ordering between vertices in directed acyclic graphs**.
To be specific, given vertices u and v, to reach vertex v, we must have reached vertex u first.
In “topological sorting”, u has to appear before v in the ordering.
The most popular algorithm for “topological sorting” is Kahn’s algorithm.

**NOTE** Kahn's algorith is great for identifying the cycle in a graph

- Great for finding the cycle in a graph
  [207. Course Schedule](LeetCode/Medium/207-course-schedule/go/main.go)

---

### Trie - Text Search

A Trie is an effective way to search for words. A good example of how a trie can be very effective in enhancing the performance of our algorithm look at
(212 Word Search II)[LeetCode/Hard/212-word-search-2/go/main.go].

An example of a TrieNode where we want to find all words is:

```
type TrieNode struct {
	children map[byte]*TrieNode
	words    []string
}
```

Note the root nodes character is "" (empty).

---

### Binary Search

Although binary search algorithms are typically used to find one element in a sorted sequence, they have many other uses.
You can apply a binary search to a result, for example. Say you wanted to determine the minimum square footage of office space needed to fit all a company's employees easily.
The square footage is a consecutive number and therefore sorted so we can perform binary search on it

---

### Floyd's Algorithm

Used for finding the start of a linked list cycle.

- [287 find-the-duplicate-number](https://leetcode.com/problems/find-the-duplicate-number/description/)
- [Floyd's cycle detection algorithm (Tortoise and hare)](https://www.youtube.com/watch?v=PvrxZaH_eZ4)

---

### GCD (Greatest Common Divisor) Euclidean algorithm

Used to find the greatest common divisor between two numbers

```go
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
```

---

### LCM (Least Common Multiple)

The Least Common Multiple of two numbers is the smallest multiple of those 2 numbers for example take LCM(6, 10)

Multiples of 6:
6, 12, 18, 24, `30`, 36, 42

Multiples of 10:
10, 20, `30`, 40, 50

Therfore the LCM(6, 10) = 30.

```
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
```

### Tarjan's Algorithm

Used to find the articulation point in a graph. An articulation point is a node that when removed will divide the graph into 2.

- [LeetCode Hard - 1568 minimum number of days to disconnect island](LeetCode/Hard/1568-minimum-number-of-days-to-disconnect-island/go/main.go).
- [YouTube: Tarjan's Strongly Connected Component (SCC) Algorithm (UPDATED) | Graph Theory](https://www.youtube.com/watch?v=wUgWX0nc4NY)

---

### Cyclic Arrays.

Sometimes you might have an array that cycles back arround.

For example this will always find the indexes 2 infront of the current index when the array is a loop.

```
(i + 1) % len(array) , (i + 2) % len(array)

```

```
i, arrLen := 1, 3
(i + 1) % arrLen = 2
(i + 2) % arrLen = 0 // looped back around
```

- [256 Paint House](LeetCode/Medium/256/go/main.go)
- [265 Paint House II](LeetCode/Hard/265-paint-house-2/go/main.go) _arrLen varies per input_

---

### Gauss' Formula

Find the sum all numbers up to n in O(1) time.
For example if I said find the sum of the first 100 numbers you could loop through all the numbers from [1..n]
to get the sum which would take O(n) time or you could use the below formula:

![Guass Formula](./assets/gaussFormula.png)

```
n := 100
sum := n * (n+1) / 2
```

### Bit Operations

Find Complement of number
[476. Number Complement](LeetCode/Easy/476-number-complement/go/main.go)

#### AND (&)

Will produce a bit in the same position where x and y have the same bit.

```
  1 0 1 1 0 (22)
& 1 1 0 1 0 (26)
----------------
  1 0 0 1 0 (18)
```

Because in Binary even numbers always have their right most bit as 0 we can quickly determine if a number is even by:

```
x & 1 == 0
```

and converseley if its odd by (because odd numbers always have their right most bit as 1):

```
x & 1 == 1
```

#### OR (|)

Will produce a bit in the same position where one of the bits is 1.

```
  1 0 1 1 0 (22)
& 1 1 0 1 0 (26)
----------------
  1 1 1 1 0 (30)
```

#### XOR (^)

Will produce a bit in the same position where both bits are different.

```
  1 0 1 1 0 (22)
& 1 1 0 1 0 (26)
----------------
  0 1 1 0 0 (12)
```

An example of this can be found here [LeetCode 268 Missing Number](LeetCode/Easy/268-Missing_Number/go/main.go)

#### NOT (~)

The NOT operation ~x produces a number where all bits of x have been inverted.
Used to invert the bits. The formula ~x == -x - 1 holds true.
In Go there is no ~ operator but we can use XOR.

```
^29 = -30

This holds true to ~x == -x - 1
```

#### Bit Shift (<<) (>>)

Left bit shift << appends K zero bits to the number and Right bit shift >> removes K zero bits from the number.
NOTE:

```
x << k is the same as x * 2^k

x >> k is the same as x / 2^k
```

![bitwise properties](./assets/bitwiseProperties.png)

### 1e9 + 7

Some questions becasue the number can get really large ask you to mod by 1,000,000,007.
This is equivalent to 1e9+7.
1e9 is the same as saying 1 multiplied by 10^9.
Therefore we can do the following:

```
answer % (1e9+7)
```

### Bases

123.45 in decimal base 10

```
123.45 = (1 * 10^2) + (2 * 10^1) + (3 * 10^0) + (4 * 10^-1) + (5 * 10^-2)
```

720.5 in octal base 8. This will also convert the number to base 10.

```
720.5 = (7 * 8^2) + (2 * 8^1) + (0 * 8^0) + (5 * 8^-1)
```

Conversion from base-10 (decimal) to base-2 (binary) we divide by 2:
Convert 50 to base-2.

```
    50 / 2 = 25; 50 % 2 = 0
    25 / 2 = 12; 25 % 2 = 1
    12 / 2 = 6; 12 % 2 = 0
    6 / 2 = 3; 6 % 2 = 0
    3 / 2 = 1; 3 % 2 = 1
    1 / 2 = 0; 1 % 2 = 1

decimal (base-2) = 010011
```

Convert 0.6875 to base-2 binary.
Keep multiplying the fractional part until the fractional part becomes 0.

```
0.6875 * 2 = 1.375 = 1
0.375 * 2 = 0.75 = 0
0.75 * 2 = 1.5 = 1
0.5 * 2 = 1.0 = 1

decima (base-2) = 1011
```

When converting between different bases the best approach is to convert to base-10 first and then convert to the target base.
The exception to this is when converting binary to octal or hexadecimal.

```
101110010 = 101 | 110 | 010 = 562 in octal/base-8 (we group into 3)


101110010 = 1 | 0111 | 0010 = 172 in hexadecimal/base-16 (we group into 4)
```

### Lexicographically smallest/largest string

A string is lexicographically smaller when:

- the first character is smaller despite the length.
- smaller length when starting characters are the same.
  example:
  `a` is lexicographically smaller than `aa`.
  `abcd` is lexicographically smaller than `b`

A string is lexicographically larger when:

- the first character is larger despite the length.
- larger length if starting characters are the same.
  example:
  `ba` is lexicographically larger than `b`.
  `b` is lexicographically larger than `abcd`

a great way to compare strings lexicographically in Go is via

```
strings.Compare(str1, str2)
```

### Prefix Sum

**Problems that require sequences of elements to meet criteria often utilize prefix sums.**

The prefix sum, cumulative sum, inclusive scan, or simply scan of a sequence of numbers x0, x1, x2, ... is a
second sequence of numbers y0, y1, y2, ..., the sums of prefixes (running totals) of the input sequence:

| 0   | 1   | 2   | 3   | 4   | 5   | 6   |
| --- | --- | --- | --- | --- | --- | --- |
| 0   | 1   | 3   | 6   | 10  | 15  | 21  |

This is really handy now because if we get a list of ranges for which we have to calculate the sum for ie: the range [3,5] as an example.
Then we know that at index 5 its ranges [0,2] + [3,5] so the sum of the range [3,5] would be the sum at index 5 (sum 0-5) minus the sum at index 2 (sum 0-2).
Therefore 15-3 = 12.
We can then be given a large number of ranges and because we computed the sum in O(n) time then for each range m its a simple O(1) lookup.

#### Calculating total number of contiguous sub arrays

| 1   | 2   | 3   | 4   |
| --- | --- | --- | --- |
| 1   | 3   | 6   | 10  |

By using Prefix Sum we can calculate the total number of contiguous sub arrays.

For example if I had the array:
| 1 | 2 |
| - | - |
| 1 | 3 | _Prefix Sum_

The contiguous sub arrays I can have are:

- 1
- 1,2
- 2

= 3 (same as prefix sum)

Another example if I had the array:
| 1 | 2 | 3 |
| - | - | - |
| 1 | 3 | 6 | _Prefix Sum_

The contiguous sub arrays I can have are:

- 1
- 1,2
- 1,2,3
- 2
- 2,3
- 3

= 6 (same as prefix sum)

This can also be calculated simply with the below formula in O(1) time:

```math
n*(n+1) / 2
```

### Fenwick Trees / Binary Index Tree

Good for Range based calculations:

Prefix Sum is good for range based calculations but if we decide to update the values constantly then that means
we would need to recalculate the Prefix Sum each time the values changed. For this specific situation a Fenwick tree or
Binary Index Tree becomes very useful as we can search in O(log.n) time and update the values in O(log.n) time.

[Sum based Fenwick Tree Implementation](DataStructures/fenwick-tree/go/sum/main.go)

[Tutorial](https://www.youtube.com/watch?v=CWDQJGaN1gY)

Example Use cases

- Range Sum
- Range Min/Max
- Range XOR

Space/Time

- space O(n)
- search O(log.n)
- update O(log.n)
- create/initialise O(n.log.n)

---

---

## Achievements

![100daysbadge](assets/100daysbadge.gif)
