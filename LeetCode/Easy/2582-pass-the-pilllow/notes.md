## [2582. Pass the Pillow](https://leetcode.com/problems/pass-the-pillow/description/)

Doing this question in O(n) is relatively easy but doing this in O(1) involves being a bit more clever.

Because we pass the pillow every second and we move up and down the line lets try and work
out how long it will take to do one full cycle.

```
n = 4

1 -> 2 -> 3 -> 4 -> 3 -> 2 -> 1

looking at this it takes a total of 6 seconds.
We can derive a formula for this (n - 1) * 2
where n is the number of people in the line.
```

then logically you would think if it takes 6 seconds to do one full cycle and the time im given is say 7 well whats the remainder of seconds after one complete cycle
```
7 / 6 = 1

ok cool but im 1 shy of 2. Oh yeah thats because i start with 1 person so I just add 1 to that and i get:

(7 / 6) + 1 = 2

we can now derive the formula for working out the person as

time % ((n - 1) * 2) + 1
```

this works but we have an edge case.
```
8 % 6 + 1 = 3 ✅
9 % 6 + 1 = 4 ✅
10 % 6 + 1 = 5 ❌ should be 3
11 % 6 + 1 = 6 ❌ shoule be 2
12 % 6 + 1 = 1 ✅

```
# 🤔 hmm ok 

well maybe if the result is larger than the number of people we now need to work out a different approach. 

maybe i can work out how much I have overlapped the number of people by
```
10 % 6 + 1 = 5 ❌

4 - 5 = 1 ok so I have overlapped the line of people by 1.

so now lets subtract the line of people by how much we overlap by.

4 - 1 = 3 ✅
```

# 🎉
__thats what where after__

so any time we go over the number of people using our formula we derived we just have to cover that edge case using the above approach.