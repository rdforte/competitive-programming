[6368. Find the Divisibility Array of a String](https://leetcode.com/contest/weekly-contest-334/problems/find-the-divisibility-array-of-a-string/)

This question was a little bit tricky in regards to how we approach the modules operation. If we look at the constraints
```
Constraints:

1 <= n <= 105
word.length == n
word consists of digits from 0 to 9
1 <= m <= 109
```

The length of the word when converted to a digit is over 105 digits long so we will get an overflow.

A nice trick for working out the modules of a very large number is to work your way forward through the number. Multiplying each stage by 10 and performing the modulo operation on that stage.

For example to work out the modules of 124 % 2 we can do the following.
```
int res = 0;

res = (res * 10 + 1) % 2

res is now 1;

res = (1 * 10 + 2) % 2

res is now 0

res = (0 * 10 + 4) % 2

res is 0
```
therefore the answer is 0.

We can use this approach to work out the modulo for each stage of the word[0,....,i] while also doing a check to see if at the current stage we can do the modulo m and if its 0 then we can push back a 1 into our array otherwise we push back a 0.

### Algorithm for Modules of Large Numbers.
__We can represent a large number as a string to avoid an overflow__
```cpp
// C++ program to compute mod of a big number represented
// as string
#include <iostream>
using namespace std;
 
// Function to compute num (mod a)
int mod(string num, int a)
{
    // Initialize result
    int res = 0;
 
    // One by one process all digits of 'num'
    for (int i = 0; i < num.length(); i++)
        res = (res * 10 + (int)num[i] - '0') % a;
 
    return res;
}
 
// Driver program
int main()
{
    string num = "12316767678678";
    cout << mod(num, 10);
    return 0;
}
```
Output: 
```
8
```