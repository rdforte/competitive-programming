#include "../../../stdc++.h"

using namespace std;

bool validPath(int n, vector<vector<int>> &edges, int source, int destination) {
  vector<int> adjacencyList[n];
  queue<int> q;

  for (auto e : edges ) { 
    adjacencyList[e[0]].push_back(e[1]);
    adjacencyList[e[1]].push_back(e[0]);
  }

  q.push(source);

  while(!q.empty()) {
    int node = q.front();
    q.pop();

    if (node == destination) {
      return true;
    }

    for (auto n : adjacencyList[node]) {
      q.push(n);
    }

    adjacencyList[node] = {};
  }

  return false;
};

int main() {
  cout << boolalpha;
  int n = 6;
  vector<vector<int>> edges{
    {0,1},
    {0,2},
    {3,5},
    {5,4},
    {4,3},
  };
  int source =  0;
  int destination = 5;

  cout << validPath(n, edges, source, destination);

  return 0;
}