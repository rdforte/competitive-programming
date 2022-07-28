// https://leetcode.com/problems/contains-duplicate/
#include <iostream>
#include <vector>
#include <map>

using namespace std;

class Solution
{
public:
    // brute force
    // time complexity is O(n^2)
    // space complexity is O(n)
    bool containsDuplicateBruteForce(vector<int> &nums)
    {
        bool isDuplicate = false;

        for (int i = 0; i < nums.size(); i++)
        {
            for (int j = 0; j < nums.size(); j++)
            {
                if (j != i && nums[j] == nums[i])
                    isDuplicate = true;
            }
        }
        return isDuplicate;
    };
    // time complexity is O(n) If we were to loop through the array once and lookup in map is constant time.
    // space complexity is O(n) if we were to loop through all elements once that would be n allocations to the map.
    bool containsDuplicate(vector<int> &nums)
    {
        std::map<int, bool> map;

        bool isDuplicate = false;

        for (int i = 0; i < nums.size(); i++)
        {
            int mapKey = nums[i];
            if (map.find(mapKey) == map.end())
            {
                map[mapKey] = false;
            }
            else
            {
                map[mapKey] = true;
                isDuplicate = map[mapKey];
                break;
            }
        }

        return isDuplicate;
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

    std::cout << Solution().containsDuplicate(nums);
};