# Code Template

Here are code templates for common patterns. Most examples are given in C++ but the ones that arent are labeled with the language.

## Two pointers: one input, opposite ends

```cpp
int fn(vector<int>& arr) {
    int left = 0;
    int right = int(arr.size()) - 1;
    int ans = 0;

    while (left < right) {
        // do some logic here with left and right
        if (CONDITION) {
            left++;
        } else {
            right--;
        }
    }

    return ans;
}
```

**Questions**

- [11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/description/)

## Two pointers: two inputs, exhaust both

```cpp
int fn(vector<int>& arr1, vector<int>& arr2) {
    int i = 0, j = 0, ans = 0;

    while (i < arr1.size() && j < arr2.size()) {
        // do some logic here
        if (CONDITION) {
            i++;
        } else {
            j++;
        }
    }

    while (i < arr1.size()) {
        // do logic
        i++;
    }

    while (j < arr2.size()) {
        // do logic
        j++;
    }

    return ans;
}
```

## Sliding Window

```cpp
int fn(vector<int>& arr) {
    int left = 0, ans = 0, curr = 0;

    for (int right = 0; right < arr.size(); right++) {
        // do logic here to add arr[right] to curr

        while (WINDOW_CONDITION_BROKEN) {
            // remove arr[left] from curr
            left++;
        }

        // update ans
    }

    return ans;
}
```

## Build a prefix Sum

prefix sum is basically sum all previous elements in the array with a given index.

```cpp
vector<int> fn(vector<int>& arr) {
    vector<int> prefix(arr.size());
    prefix[0] = arr[0];

    for (int i = 1; i < arr.size(); i++) {
        prefix[i] = prefix[i - 1] + arr[i];
    }

    return prefix;
}
```

## Efficient string building

```cpp
string fn(vector<char>& arr) {
    return string(arr.begin(), arr.end())
}
```

Golang:

```go
func fn(chars []byte) string {
	return string(chars)
}
```

`In JavaScript, benchmarking shows that concatenation with += is faster than using .join().`

## Linked List: fast and slow pointer

```cpp
int fn(ListNode* head) {
    ListNode* slow = head;
    ListNode* fast = head;
    int ans = 0;

    while (fast != nullptr && fast->next != nullptr) {
        // do logic
        slow = slow->next;
        fast = fast->next->next;
    }

    return ans;
}
```

## Reversing a Linked List

```cpp
ListNode* fn(ListNode* head) {
    ListNode* curr = head;
    ListNode* prev = nullptr;
    while (curr != nullptr) {
        ListNode* nextNode = curr->next;
        curr->next = prev;
        prev = curr;
        curr = nextNode;
    }

    return prev;
}
```

## Find number of subarrays that fit an exact criteria

```cpp
int fn(vector<int>& arr, int k) {
    unordered_map<int, int> counts;
    counts[0] = 1;
    int ans = 0, curr = 0;

    for (int num: arr) {
        // do logic to change curr
        ans += counts[curr - k];
        counts[curr]++;
    }

    return ans;
}
```

## Monotonic increasing stack

The same logic can be applied to maintain a monotonic queue.

```cpp

```

## Binary tree: DFS (recursive)

```cpp

```

## binary tree: dfs (iterative)

```cpp

```

## Binary tree: BFS

```cpp

```

## Graph: DFS (recursive)

For the graph templates, assume the nodes are numbered from `0` to `n - 1` and the graph is given as an adjacency list.
Depending on the problem, you may need to convert the input into an equivalent adjacency list before using the templates.

```cpp

```

## Graph: DFS (iterative)

```cpp

```

## Graph: BFS

```cpp

```

## Find top k elements with heap

```cpp

```

## Binary Search

```cpp

```

## Binary search: duplicate elements, left-most insertion point

```cpp

```

## Binary search: duplicate elements, right-most insertion point

```cpp

```

## Binary search: for greedy problems

If looking for a minimum:

```cpp

```

If looking for a maximum:

```cpp

```

## Backtracking

```cpp

```

## Dynamic programming: top-down memoization

```cpp

```

## Build a trie

```cpp

```

## Dijkstra's algorithm

```cpp

```
