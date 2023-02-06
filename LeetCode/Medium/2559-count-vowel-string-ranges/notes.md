[2559. Count vowel Strings in Ranges](https://leetcode.com/problems/count-vowel-strings-in-ranges/description/)

- space complexity = O(W + Q) where W is words and Q is queries
- time complexity = O(W + Q) because we loop through all words then all queries.

### Notes
My first approach is to loop through all queries then loop through all words for each
querie. This will give us a TLE (Time Limit Exceeded) and a runtime complexity of O(QW) where
Q is queries and W is words. So there must be a more efficient approach.

My next lot of thinking led me to think of ranges and if I knew at which index for each word I had a count for the number of VW's (words starting and ending with a vowel) I had seen then I could use the queries to calculate how many VW's I had seen by working out the difference.

Example
```
words = ["aba","bcb","ece","aa","e"], queries = [[0,2],[1,4],[1,1]]

at "aba" index 0 I have seen 1 VW.
at "bcb" index 1 I have still only seen 1 VW.
at "ece" index 2 I have now seen 2 VW's.
at "aa" index 3 I have seen 3 VW's.
at "e" index 4 I have seen 4 VW's.

Once I know how many VW's I have at each index I can then use the queries to calculate the difference.

for query [1, 4] I can then do how many VW's I have seen at index 4 - how many I have seen at Index 1. 
But because we want 1 and 4 inclusive I must do the index before 1. 
Therefore that would give me VW[4] - VW[1 - 1].
Which is 4 - 1 = 3
```