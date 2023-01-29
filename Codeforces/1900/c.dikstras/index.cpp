// https://codeforces.com/problemset/problem/20/C
#include <bits/stdc++.h>

using namespace std;

typedef long long ll;
typedef pair<ll, ll> pll;

const int MAX = 1e5 + 5;
vector<vector<pll>> adjList;
vector<ll> weights(MAX, LLONG_MAX);
vector<int> prevPath(MAX);
vector<bool> seen(MAX, false);

void printSol(ll n, vector<int> &prevPath)
{
  if (n == 0)
    return;
  printSol(prevPath[n], prevPath);
  cout << n << ' ';
};

bool dijkstra(int s, int t)
{
  priority_queue<pll, vector<pll>, greater<pll>> pq;

  pq.push({0, s});
  weights[s] = 0;

  while (!pq.empty())
  {
    auto u = pq.top();
    pq.pop();

    if (u.second == t)
      return true;

    seen[u.second] = true;

    for (auto &pr : adjList[u.second])
    {
      int v = pr.first, w = pr.second;

      if (!seen[v] && u.first + w < weights[v])
      {
        weights[v] = u.first + w;
        pq.push({weights[v], v});
        prevPath[v] = u.second;
      }
    }
  }

  return false;
}

int main()
{
  // freopen("testData.txt", "r", stdin);
  int n, e;
  cin >> n >> e;

  adjList.resize(n + 1);

  int a, b, w;
  while (e > 0)
  {
    e--;
    scanf("%d %d %d", &a, &b, &w);
    adjList[a].push_back({b, w});
    adjList[b].push_back({a, w});
  }

  if (dijkstra(1, n))
  {
    printSol(n, prevPath);
  }
  else
    cout << -1;
}