#include <iostream>
#include <vector>
#include <algorithm>

class Solution
{
public:
    // My fastest leet code solution
    // My first solution was using a vector so the space time would have been O(n)
    // Though after writing this solution I could see that I was only referencing the previouse item in the array
    // So I removed the vector and added a variable to store the previouse value
    // I was now only keeping track of 2 values so I could say the space time was O(1)
    // The runtime would be O(n) as I only have to loop the vector once
    int maxSubArrayV1(std::vector<int> &nums)
    {
        int max = nums[0];
        // std::vector<int> tm(nums.size());
        int prev = nums[0];
        // tm[0] = nums[0];

        for (int i = 1; i < nums.size(); i++)
        {
            int m = prev + nums[i];

            if (nums[i] > m)
            {
                m = nums[i];
            }

            // tm[i] = m;
            prev = m;

            if (m > max)
            {
                max = m;
            }
        }
        return max;
    }

    int maxSubArrayV2(std::vector<int> &nums)
    {
        int m = nums[0];
        int p = nums[0];

        for (int i = 1; i < nums.size(); i++)
        {
            p += nums[i];
            p = std::max(p, nums[i]);

            m = std::max(m, p);
        }
        return m;
    }

    int maxSubArrayV3(std::vector<int> &nums)
    {

        int m = nums[0];
        int p = nums[0];

        for (int i = 1; i < nums.size(); i++)
        {
            p = std::max(p + nums[i], nums[i]);
            m = std::max(m, p);
        }
        return m;
    }
};

int main()
{
    std::vector<int> v{-2, 1, -3, 4, -1, 2, 1, -5, 4};
    std::cout << Solution().maxSubArrayV3(v);
}