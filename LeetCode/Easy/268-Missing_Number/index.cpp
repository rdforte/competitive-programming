#include <iostream>
#include <vector>

// TimeComplexity = O(n) loop through nums once.
// SpaceComplexity = O(1) as we have to keep track of missing number.

using namespace std;

class Solution
{
public:
    int missingNumber(vector<int> &nums)
    {
        int missingNumber = nums.size();

        for (int i = 0; i < nums.size(); i++)
        {
            missingNumber ^= i ^ nums[i];
        }

        return missingNumber;
    }
};

int main()
{
    vector<int> nums{0, 1, 2, 3, 5};

    cout << Solution().missingNumber(nums);
}