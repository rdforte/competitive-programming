[Introductory Problem Two Knights](https://cses.fi/problemset/task/1072/)

[Combinations and Permutations Math formulas](https://www.mathsisfun.com/combinatorics/combinations-permutations.html)

This question is a combinations question. We want to find all the combinations for two Knights placed on an n * n chessboard
such that the two Knights do not attack one another.

We know that it is combinations because we can not have duplicates. If we could have duplicates then it would just be
permutations.

The formula for working out combinations is
```
n! / r!(n - r)!
```

where __n__ is the number of items and __r__ is how many of those items we are going to be using.
For example:
If I had 10 pieces of unique fruit and I was only able to select two at random. What are the total number of combinations
for which I could select 2 fruit?

Note: (apple, banana) and (banana, apple) is a permutation. We only want the combination so it would only be one of these
and not both ie: just banana, apple

10! / 2!(10 - 2)!

In regards to the Knights question instead of 10 apples we have an n*n grid for which we only want to select 2 squares at
a time from this n*n grid. Therefore the combinations would be
```
g = n*n
g! / 2!(g - 2)!
```
we can simplify this to:
```
(g * (g - 1)) / 2
```
for example if g was 5:
```
5 * 4 * 3 * 2 * 1 / (2 * 1)(3 * 2 * 1)
```
You can see that in the numerator we can cancel out the 3 * 2 * 1 and in the denominator we can also cancel this out.

Now that we have all the combinations we need to figure out all the combinations for which the knights attack each other.