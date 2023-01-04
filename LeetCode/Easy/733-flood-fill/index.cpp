#include <bits/stdc++.h>

using namespace std;

typedef vector<vector<int>> vvi;

class Solution
{
public:
  vvi floodFill(vvi &image, int sr, int sc, int color)
  {
    dfs(image, sr, sc, image[sr][sc], color);
    return image;
  }

private:
  void dfs(vvi &image, int sr, int sc, int ogColor, int newColor)
  {
    if (!(sr >= 0 and sr < image.size()) or !(sc >= 0 and sc < image[0].size()))
      return;

    if (image[sr][sc] != ogColor or image[sr][sc] == newColor)
      return;

    image[sr][sc] = newColor;

    dfs(image, sr - 1, sc, ogColor, newColor);
    dfs(image, sr + 1, sc, ogColor, newColor);
    dfs(image, sr, sc - 1, ogColor, newColor);
    dfs(image, sr, sc + 1, ogColor, newColor);
  }
};