// https://leetcode.com/problems/counting-bits/
#include <iostream>
#include <vector>

/**
 * In order to get the number of 1's in the binary version of the number we have to use the bitwise & operator as a mask and shift the number.
 * Example:
 * The number 6
 *
 * 110 <- number 6
 * 001 <-mask
 * = 0 because 0 & 1 = 0
 *
 * shift 110 by 1
 *
 * 011
 * 001
 * = 1 so now we add 1 to the count
 *
 * shift by 1 again
 *
 * 001
 * 001
 * = 1 add 1 to the count
 *
 * if we shift 001 again we get the number 0 so we just check for 0 and if it is 0 then we stop and return the count.
 *
 * total = 2
 */

// time complexity is outer loop of n + inner loopf of log.n which = n.log.n
// space complexity is n because we have to keep track of n items in ans vector.
class Solution
{
public:
    std::vector<int> countBits(int n)
    {
        std::vector<int> ans(n + 1);

        // to run this loop is n
        for (int i = 0; i <= n; i++)
        {
            int c = 0;
            int num = i;
            // to Run this loop is O(logn)
            while (num)
            {
                c += num & 1;
                num = num >> 1;
            }
            ans[i] = c;
        }
        return ans;
    }
};

int main()
{
    int n = 7;
    int count = 0;

    std::vector<int> ans = Solution().countBits(n);

    for (int i = 0; i < ans.size(); i++)
    {
        std::cout << ans[i] << " ";
    }
}