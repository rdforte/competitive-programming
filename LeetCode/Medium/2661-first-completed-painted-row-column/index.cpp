#include "../../../stdc++.h"

using namespace std;

class Solution
{
public:
    int firstCompleteIndex(vector<int> &arr, vector<vector<int>> &mat)
    {
        int numRows = mat.size(), numCols = mat[0].size();
        vector<int> rows(numRows);
        vector<int> cols(numCols);

        // build coordinates.
        vector<pair<int, int>> coords((mat.size() * mat[0].size()) + 1);
        for (int r = 0; r < numRows; r++)
        {
            for (int c = 0; c < numCols; c++)
            {
                coords[mat[r][c]] = {r, c};
            }
        }

        // go through arr, identify coords and increment count for row and col.
        for (int i = 0; i < arr.size(); i++)
        {
            auto [r, c] = coords[arr[i]];
            if (++rows[r] == numCols || ++cols[c] == numRows)
                return i;
        }

        return -1;
    }
};

int main()
{
}