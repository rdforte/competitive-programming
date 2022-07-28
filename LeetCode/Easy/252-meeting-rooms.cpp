// https://leetcode.com/problems/meeting-rooms/
#include <iostream>
#include <vector>
#include <set>
#include <algorithm>

class Solution
{
public:
    // Brute Force Approach
    // Time Complexity O(n^2) because we have a loop within a loop that interates through all intervals.
    // Space Cmplexity O(1) No additional space is used.
    bool canAttendMeetings(std::vector<std::vector<int>> &intervals)
    {
        for (int i = 0; i < intervals.size(); i++)
        {
            for (int j = 0; j < intervals.size(); j++)
            {
                if (i == j)
                    continue;
                if (intervals[j][0] >= intervals[i][0] && intervals[j][0] < intervals[i][1])
                {
                    return false;
                }
            }
        }
        return true;
    }

    // Optimal Solution
    // Time Complexity O(n*log*n) It takes n*log*n time to sort the intervals and n time to loop them
    // so with n*log*n being the most time intensive compared to n we can just say it takes O(n*log*n)
    // Space Complexity is O(1) as no additional space is used.
    bool canAttendMeetingsV2(std::vector<std::vector<int>> &intervals)
    {
        if (intervals.size() <= 1)
            return true;

        std::sort(intervals.begin(), intervals.end());

        for (int i = 1; i < intervals.size(); i++)
        {
            if (intervals[i][0] < intervals[i - 1][1])
                return false;
        }

        return true;
    }
};

int main()
{
    std::vector<std::vector<int>> intervals{{0, 30}, {5, 10}, {15, 20}};
    // std::vector<std::vector<int>> intervals{{20, 30}, {30, 31}, {15, 20}};
    // std::vector<std::vector<int>> intervals{{7, 10}, {2, 4}};

    std::cout << Solution().canAttendMeetingsV2(intervals);
}