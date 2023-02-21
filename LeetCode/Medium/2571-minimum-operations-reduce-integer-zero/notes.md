[2571. Minimum Operations to Reduce an Integer to 0](https://leetcode.com/problems/minimum-operations-to-reduce-an-integer-to-0/description/)

At first glance the question asks for the minimum number of operations so I thought this would be a DP (Dynamic Programming)
question and I could work my way forward iteratively until I arrived at __n__.
This was not the case because the optimal solution may be passed n and I would not know how far past __n__ this may be.

In order to get __n__ to 0 I must use (+|-) operations which are 2^x. So this may include adding 2 to the power of some number or subtracting 2 to the power of some number.

The first step I took was drawing this out so I could get a visual representation of the question.
Using the first example.
```
Input: n = 39
Output: 3
Explanation: We can do the following operations:
- Add 20 = 1 to n, so now n = 40.
- Subtract 23 = 8 from n, so now n = 32.
- Subtract 25 = 32 from n, so now n = 0.
It can be shown that 3 is the minimum number of operations we need to make n equal to 0.
```

I know that if I can get my input to any number which is 2^x I knew this would be 1 operation and we want the minimum number of operations.