#include "../../../stdc++.h"

using namespace std;

vector<vector<int>> floodFill(vector<vector<int>> &image, int sr, int sc, int color)
{
    if (image[sr][sc] == color)
        return image;

    queue<pair<int, int>> q;

    q.push({sr, sc});

    int nodeColor = image[sr][sc];

    while (!q.empty())
    {
        auto pixel = q.front();
        q.pop();

        image[pixel.first][pixel.second] = color;

        if (pixel.second - 1 >= 0 && image[pixel.first][pixel.second - 1] == nodeColor)
            q.push({pixel.first, pixel.second - 1});

        if (pixel.first - 1 >= 0 && image[pixel.first - 1][pixel.second] == nodeColor)
            q.push({pixel.first - 1, pixel.second});

        if (pixel.second + 1 < image[image.size() - 1].size() && image[pixel.first][pixel.second + 1] == nodeColor)
            q.push({pixel.first, pixel.second + 1});

        if (pixel.first + 1 < image.size() && image[pixel.first + 1][pixel.second] == nodeColor)
            q.push({pixel.first + 1, pixel.second});
    }

    return image;
}

int main()
{
}