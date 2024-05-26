# [3164. Find the Number of Good Pairs II](https://leetcode.com/problems/find-the-number-of-good-pairs-ii/description/)

The brute force approach would be for each nums1[i] go through all the nums of nums2 and see if the current nums1[i] is divisible by the nums2[j]\*k.
This would result in a time complexity of n^2.

Because we are trying to see if nums2[j]\*k is a factor of nums1[i]. If we knew all the factors of nums1 ahead of time then determining if nums2[j] was
a factor would be an O(1) operation.

To get the factors we can just loop up until the sqrt of nums1[i].
The reason we can just go to the sqrt of nums1[i] is because for example:

- 1 x 12 = 12
- 2 x 6 = 12
- 3 x 4 = 12
- 4 x 3 = 12 -- exclude
- 6 x 2 = 12 -- exclude
- 12 x 1 = 12 -- exclued

To get the factors we exclude its simple algebra:

```
4 * x = 12

12/4 = 3
```

Because we are just trying to find the count of good pairs we can just keep a tally of the factors for example:

lets say we have nums1 = [10,20,30]

10 -> 1,2,5,10

20 -> 1,2,5,10,20

30 -> 1,3,5,10,15,30

on the hash table it will be like this:

- [1] = 3 -- we have a total of 3 nums for which 1 is a factor
- [2] = 2
- [3] = 1
- [5] = 3
- [10] = 3
- [15] = 1
- [20] = 1
- [30] = 1
