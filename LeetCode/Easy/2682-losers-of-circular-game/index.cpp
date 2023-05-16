#include "../../../stdc++.h"

using namespace std;

class Solution
{
public:
    vector<int> circularGameLosers(int n, int k)
    {
        vector<bool> friends(n, false);

        for (int i = 0, j = 0;; i += (++j * k))
        {
            if (i >= n)
            {
                i = (i % n);
            }

            if (friends[i])
            {
                break;
            }

            friends[i] = true;
        }

        vector<int> leftOvers;
        for (int i = 0; i < n; i++)
        {
            if (!friends[i])
            {
                leftOvers.push_back(i + 1);
            }
        }

        return leftOvers;
    }
};

int main()
{
}