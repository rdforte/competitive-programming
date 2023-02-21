[2571. Minimum Operations to Reduce an Integer to 0](https://leetcode.com/problems/minimum-operations-to-reduce-an-integer-to-0/description/)

At first glance the question asks for the minimum number of operations so I thought this would be a DP (Dynamic Programming) question and I could work my way forward iteratively until I arrived at __n__.
This was not the case because the optimal solution may be passed __n__ and I would not know how far past __n__ this may be.

In order to get __n__ to 0 I must use (+|-) operations which are 2^x. So this may include adding 2 to the power of some number or subtracting 2 to the power of some number __repetitively__ until I arrived at 0.

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

I know that if I can get my input to any number which is 2^x I knew this would be 1 operation and we want the minimum number of operations so getting __n__ to this number in the min number of steps would help to get me the answer I was after.

I drew this out.

```
n = 39

32(2^5)-33-34-35-...[39]...60-61-62-63-64(2^6)
```

from this I can see that 39 is closes to 32(lower bound) and furthest from 64(upper bound). So how many steps would it take to get n to 32.

```
39-32 = 7

but how many steps are involved to get 7 to 0?
```

If I repeat my first step
```
n = 7

4(2^2)-5-6-[7]-8(2^3)
```

hmm ok I can see this is just the same problem repeating itself for a smaller subset.
Ok this looks to be a recursive problem.

7 is closest to 8.

```
8 - 7 = 1

1 is 2^0
```

Looking at this from a high level it was how could i recursively take __n__ and get it
to the nearest 2^x.

#### NOTE: 
I used the bitwise left and right shift operators here. They work really well for squaring and square root.
example:
```
1 in binary is 1
2 in binary is 10
4 in binary is 100
8 in binary is 1000

what << does is move the bit at the front one place to the left and >> moves it 1 place to the right.
```