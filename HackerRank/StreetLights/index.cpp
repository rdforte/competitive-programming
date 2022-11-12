#include <bits/stdc++.h>

using namespace std;

int main()
{
  int numTestCases;
  cin >> numTestCases;

  while (numTestCases--)
  {
    int n, k;
    cin >> n >> k;

    vector<pair<int, int>> coverage;

    while (n--)
    {
      int light;
      cin >> light;
      coverage.push_back({light - k, light + k});
    }

    sort(coverage.begin(), coverage.end());

    int count = coverage.size();

    for (int i = 0; i < coverage.size(); i++)
    {
      bool isLastLight = i == coverage.size() - 1;
      int lower = coverage[i].first;
      auto nextLight = isLastLight ? coverage[i].second : coverage[i + 1].first;
      int upper = min(coverage[i].second, nextLight);

      if (!isLastLight && coverage[i].second >= coverage[i + 1].first)
        count--;

      count += (upper - lower);
    }

    cout << count << "\n";
  }
}