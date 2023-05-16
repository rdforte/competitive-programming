#include "../../../stdc++.h"

using namespace std;

class Solution
{
public:
    bool doesValidArrayExist(vector<int> &derived)
    {
        vector<int> original(derived.size(), 0);
        original[0] = 0;

        for (int i = 0; i < derived.size() - 1; i++)
        {
            if (derived[i] == 1)
            {
                original[i + 1] = original[i] == 0 ? 1 : 0;
            }
            else
            {
                original[i + 1] = original[i] == 0 ? 0 : 1;
            }
        }

        return derived[derived.size() - 1] == original[original.size() - 1] ^ original[0];
    }
};

int main()
{
}