// https://leetcode.com/problems/number-of-1-bits/
#include <iostream>
#include <vector>

class Solution
{
public:
    // time complexity O(1) we loop through 32 times regardless of n so we can say it is just O(1)
    // space compexity O(1)
    int hammingWeight(uint32_t n)
    {
        int b = 0;
        uint m = 1; // when working with bits we have to use unsigned ints as to not have negatives

        for (int i = 0; i < 32; i++) // loop through 32 times because its 32 bits
        {
            if ((n & m) != 0) // bitwise &. if its 0 & 1 ie false and true its 0 false otherwise 1 & 1 is true and true which is 1.
                b++;

            m <<= 1; // shift the mask by 1 ie: 1 => 10 => 100 => 1000
        }
        return b;
    }
};

int main()
{
    int sol = Solution().hammingWeight(00000000000000000000000000001011);
}
