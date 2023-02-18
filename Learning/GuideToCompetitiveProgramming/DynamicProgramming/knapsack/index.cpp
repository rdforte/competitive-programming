#include <bits/stdc++.h>

using namespace std;

typedef vector<vector<bool>> vvb;

vector<bool> isPossibleKnapsack(vector<int> weights)
{
  int totalSum = 0;
  for (auto w : weights)
    totalSum += w;

  weights.insert(weights.begin(), 0);

  int col = totalSum + 1;
  int row = weights.size();
  vvb possible(col, vector<bool>(row, false));

  possible[0][0] = true;

  for (int w = 1; w < weights.size(); w++)
  {
    for (int s = 0; s <= totalSum; s++)
    {
      if (s - weights[w] >= 0)
      {
        possible[s][w] = possible[s][w] | possible[s - weights[w]][w - 1];
      }
      possible[s][w] = possible[s][w] | possible[s][w - 1];
    }
  }

  vector<bool> res;
  for (auto p : possible)
  {
    res.push_back(p[row - 1]);
  }
  return res;
};

int main()
{
  vector<int> w{1, 3, 3, 5};
  auto sol = isPossibleKnapsack(w);
  for (auto s : sol)
  {
    cout << boolalpha << s << ", ";
  }
}