#include <iostream>
#include <map>
#include <vector>

class Solution
{
public:
    int singleNumber(std::vector<int> &nums)
    {
        std::map<int, int> m;
        for (int i = 0; i < nums.size(); i++)
        {
            if (m[nums[i]] == 0)
            {
                m[nums[i]] = 1;
            }
            else
            {
                m[nums[i]] = 1;
            }
        };
        for (auto element : m)
        {
        }
    }
};

int main()
{
    // std::map<int, int> m;
    // std::cout << (m[1] == 0);
    int data[] = {1, 2, 3, 4, 5, 6};

    for (int n : data)
    {
    }
}
