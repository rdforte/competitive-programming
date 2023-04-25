#include "../../../stdc++.h"

using namespace std;

class Solution
{
public:
    vector<int> rowAndMaximumOnes(vector<vector<int>> &mat)
    {
        int p = 0;
        int o = 0;

        for (int i = 0; i < mat.size(); i++)
        {
            int c = 0;
            for (int j = 0; j < mat[0].size(); j++)
                if (mat[i][j] == 1)
                    c++;

            if (c > o)
            {
                p = i;
                o = c;
            }
        }
        return {p, o};
    };
};

int main()
{
}