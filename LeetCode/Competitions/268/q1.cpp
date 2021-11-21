#include <iostream>
#include <vector>
#include <algorithm>

class Solution
{
public:
    int maxDistance(std::vector<int> &colors)
    {
        int i = 0;
        int max = 0;

        for (int j = colors.size() - 1; j >= 0; j--)
        {
            if (colors[j] != colors[i])
            {
                max = std::max(max, j - i);
            }
        }
        int k = colors.size() - 1;
        for (int j = 0; j < colors.size(); j++)
        {
            if (colors[j] != colors[k])
            {
                max = std::max(max, k - j);
            }
        }
        return max;
    }
};

int main()
{

    std::vector<int> t = {4, 4, 4, 11, 4, 4, 11, 4, 4, 4, 4, 4};

    std::cout << Solution().maxDistance(t);
}