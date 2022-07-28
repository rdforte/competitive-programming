// https://leetcode.com/problems/missing-number/
#include <iostream>
#include <vector>
#include <algorithm>
#include <set>

using namespace std;

class Solution
{
public:
    // time complexity O(n.log.n) it takes nlogn time to sort and only o(n) time to check for the missing number nlogn > n therefore nlogn
    // space complexity o(n) the size of the array
    int missingNumberV1(vector<int> &nums)
    {
        sort(nums.begin(), nums.end());
        int n = nums.size();
        int missingNumber;

        if (nums[nums.size() - 1] != n)
            return n;

        if (nums[0] != 0)
            return 0;

        for (int i = 1; i < nums.size(); i++)
        {
            if (nums[i - 1] != nums[i] - 1)
            {
                missingNumber = nums[i] - 1;
                break;
            }
        }
        return missingNumber;
    };
    // time complexity O(n) takes n time to populate the set and n time to retrieve from the set so we can just say n
    // space complexity O(n) to add the values to the set
    int missingNumberV2(vector<int> nums)
    {
        set<int> s;
        for (int i = 0; i < nums.size(); i++)
            s.insert(nums[i]);

        for (int i = 0; i <= nums.size(); i++)
        {
            if (s.find(i) == s.end())
                return i;
        }

        return -1;
    }
    // Prefered solution using XOR is more performant.
    // time complexity O(n)
    // space complexity O(n)
    int missingNumber(vector<int> nums)
    {
        int n = nums.size();

        for (int i = 0; i < nums.size(); i++)
        {
            n ^= i ^ nums[i];
        }
        return n;
    }
};

int main()
{
    std::vector<int> nums;

    int n;
    while (std::cin >> n)
    {
        nums.push_back(n);
    }

    std::cout << Solution().missingNumber(nums);
};