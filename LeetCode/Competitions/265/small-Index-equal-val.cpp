#include <iostream>
#include <vector>
#include <algorithm>

class Solution
{
public:
    int smallestEqual(std::vector<int> &nums)
    {
        int s = -1;

        for (int i = 0; i < nums.size(); i++)
        {
            if (i % 10 == nums[i])
            {
                s = s == -1 ? i : std::min(s, i);
            }
        }
        return s;
    }
};

int main()
{
    std::vector<int> v{6, 5, 4, 4, 9, 1, 5, 0, 8, 8, 5, 8, 0, 9, 8, 3, 6, 5, 2, 7, 7, 6, 6, 8, 9, 6, 5, 6, 5, 6, 8, 6, 9, 5, 1, 0, 5, 5};
    std::cout << Solution().smallestEqual(v);
}
