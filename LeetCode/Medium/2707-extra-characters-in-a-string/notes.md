## [2707. Extra Characters in a String](https://leetcode.com/problems/extra-characters-in-a-string/description/)

If I read the sentence __Return the minimum number of extra characters left over if you break up s optimally.__

The keyword __optimally__ is a give away that this might be a DP (Dynamic Programming) question.

Using the example: 

```
Input: s = "leetscode", dictionary = ["leet","code","leetcode"]
Output: 1
Explanation: We can break s in two substrings: "leet" from index 0 to 3 and "code" from index 5 to 8. There is only 1 unused character (at index 4), so we return 1.
```

Lets break this down further to try and make it simpler. Usually if we can break something down into smaller sub problems
and apply recursion with memoization its a pretty good advocate for dynamic programming.


```
Input: s = "leet", dictionary = []
```

If I had 0 substrings to match on how many characters would I return?

```
"" -> L -> E -> E -> T
0     1    2    3    4
```


I start with an empty string because with DP we always have a base case. In this case its an empty string. Each time I move through the string
I increment the count for each index to represent the number of characters I have not used.

I would return 4 characters. 

Now what If I had an exact substring that matched my original substring.

```
Input: s = "leet", dictionary = ["leet"]
```

If I started at L I don't have the substring L so I move to E. I have to see if I have the substring L and LE now. I don't have
that substring so I move to LEE and check to see if I have E, EE and LEE I don't have those substrings either so my total characters which I have left over at the moment
is 3. 

I get to the letter T so i check for


- T
- ET
- EET
- LEET - I have a match now what?

Ive used 4 characts here and my starting position is L right through to T. So i can negate all those characters used and say that at index 4 I've used as many characters as the index before letter L which is 0.

I can then repeat this process for longer strings and substrings.


