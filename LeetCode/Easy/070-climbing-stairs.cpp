#include <iostream>
#include <vector>

class Solution
{
public:
    int climbStairs(int n)
    {
        std::vector<int> s(n + 1, 0); // start at 0 and align index with stair
        s[0] = 1;

        for (int i = 0; i < s.size() - 1; i++)
        { // we dont need to do the last stair
            s[i + 1] += s[i];
            if (i + 2 <= n)
            {
                s[i + 2] += s[i];
            }
        }
        return s[n];
    }
};

int main()
{
    int n;
    std::cin >> n;
    std::cout << Solution().climbStairs(n) << "\n";
}