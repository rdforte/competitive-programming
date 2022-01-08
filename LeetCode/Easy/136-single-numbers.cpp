#include <iostream>
#include <map>
#include <vector>
#include <stdexcept>

class Solution
{
public:
    // Hash Map
    // time complexity = O(n) we have to complete a max of 2 loops of n operations. The biggest loop is most likely to be the first
    // where we add the items to the hashmap.
    // space complexity O(n) we structure a map to keep track of the items
    int singleNumberV1(std::vector<int> &nums)
    {
        std::map<int, int> m;
        for (int i = 0; i < nums.size(); i++)
            m[nums[i]]++;

        for (auto const [key, value] : m)
        {
            if (value == 1)
                return key;
        }

        throw std::invalid_argument("There was no single int");
    }
    // XOR method
    // XOR of 0 and bit is always the bit
    // (x ^ x) always cancels out
    // therefore we can be left with the number that is different
    // time complexity O(n)
    // space complexity O(1) we only keep track of the current xor
    int singleNumberV2(std::vector<int> &nums)
    {
        int n = 0;
        for (int num : nums)
        {
            n ^= num;
        }
        return n;
    }
};

int main()
{
    std::vector<int> v{1, 2, 2, 1, 4};
    std::cout << Solution().singleNumberV2(v);
}
