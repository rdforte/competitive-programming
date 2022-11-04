#include "../../../stdc++.h"

using namespace std;

vector<vector<int>> floodFill(vector<vector<int>> &image, int sr, int sc, int color)
{
    if (image[sr][sc] == color)
        return image;

    dfs(image, sr, sc, color, image[sr][sc]);

    return image;
}

void dfs(vector<vector<int>> &image, int sr, int sc, int newColor, int nodeColor)
{
    if (sr >= image.size() || sr < 0 || sc >= image[image.size() - 1].size() || sc < 0)
        return;

    if (image[sr][sc] != nodeColor)
        return;

    image[sr][sc] = newColor;

    dfs(image, sr, sc - 1, newColor, nodeColor);
    dfs(image, sr - 1, sc, newColor, nodeColor);
    dfs(image, sr, sc + 1, newColor, nodeColor);
    dfs(image, sr + 1, sc, newColor, nodeColor);
}

int main()
{
}