[299. Bulls and Cows](https://leetcode.com/problems/bulls-and-cows/description/?envType=study-plan&id=level-1)

The way in which you plays bulls and cows is as follows:
```
On a sheet of paper, the players each write a 4-digit secret number. The digits must be all different. 
Then, in turn, the players try to guess their opponent's number who gives the number of matches. 
The digits of the number guessed also must all be different. 
If the matching digits are in their right positions, they are "bulls", if in different positions, they are "cows". Example:

Secret number: 4271
Opponent's try: 1234
Answer: 1 bull and 2 cows. (The bull is "2", the cows are "4" and "1".)
The first player to reveal the other's secret number wins the game.
```

### Notes
This time when doing the question I started off with the following:
1. I read the question 2 times and understood exactly what it was asking of me before proceeding.
2. I looked at the constraints to see what it was I was working with.
```
1 <= secret.length, guess.length <= 1000 
secret.length == guess.length // cool i don't have to worry about them being different lengths
secret and guess consist of digits only. // I'm only working with digits 0-9
```

Finding the Bulls was pretty straight forward because it was just a matter of making sure the
numbers at each index were aligned. 

The Cows were a bit trickier because the numbers could be in any position, so I needed a way to 
keep track of this. I could have gone a map but a array is much faster and because i was working with only single digits i could just use an array with a length of 10.

because each digit was a character and when converted to an int it will be its ASCII representation. I could just do the character version of the digit minus the smallest character I needed to account for which was '0'.
The same goes with say letters. If I had all lowercase letters from a-z i could just create an
array with length 26 and when referencing the character at a particular index i could do the character - 'a'.