// https://leetcode.com/problems/evaluate-division/
#include <iostream>
#include <vector>
#include <unordered_map>
#include <string>
#include <set>

using namespace std;

// Node represents the Node in the graph with its current weight and a link to its parent.
class Node
{
public:
  string parent;
  double weight;

  Node(string n, double w)
  {
    parent = n;
    weight = w;
  }
};

class UnionFind
{
private:
  unordered_map<string, Node *> roots;

public:
  UnionFind(set<string> nodes)
  {
    for (auto n : nodes)
    {
      Node *node = new Node(n, 1);
      roots[n] = node;
    }
  }

  string findRoot(string n)
  {
    while (roots[n]->parent != n)
    {
      n = roots[n]->parent;
    }
    return n;
  }

  void unionNodes(string a, string b, double result)
  {
    string rootA = findRoot(a);
    string rootB = findRoot(b);

    if (rootA != rootB)
    {

      double bResWeight = roots[a]->weight / result;
      double ratio = bResWeight / roots[b]->weight;
      roots[b]->weight = bResWeight;

      for (auto r : roots)
      {
        if (r.second->parent == rootB && r.first != b)
        {
          r.second->weight *= ratio;
        }
      }
      cout << a << " ---- " << b << "-------------------------\n";
      for (auto r : roots)
      {
        cout << r.first << " : " << r.second->parent << " == " << r.second->weight << "\n";
      }
      cout << "--------------------------------\n";

      roots[rootB]->parent = rootA;
    }
  }

  bool
  isConnected(string a, string b)
  {
    if (roots.find(a) != roots.end() &&
        roots.find(b) != roots.end())
    {
      return findRoot(a) == findRoot(b);
    }

    return false;
  }

  unordered_map<string, Node *> getRoots()
  {
    return roots;
  }
};

class Solution
{
public:
  vector<double> calcEquation(vector<vector<string>> &equations, vector<double> &values, vector<vector<string>> &queries)
  {
    set<string> s{};
    for (auto eq : equations)
    {
      s.insert(eq[0]);
      s.insert(eq[1]);
    }

    UnionFind uf(s);

    for (int i = 0; i < equations.size(); i++)
    {
      uf.unionNodes(equations[i][0], equations[i][1], values[i]);
    }

    vector<double> results{};
    for (auto q : queries)
    {
      bool isConnected = uf.isConnected(q[0], q[1]);
      if (isConnected)
      {
        double val = uf.getRoots()[q[0]]->weight / uf.getRoots()[q[1]]->weight;
        results.push_back(val);
      }
      else
      {
        results.push_back(-1.0);
      }
    }

    return results;
  }
};

int main()
{

  // expected: [1.33333,1.00000,-1.00000]
  // vector<vector<string>> equations{
  //     {"a", "e"},
  //     {"b", "e"},
  // };

  // vector<double> values{
  //     4.0, 3.0};

  // vector<vector<string>> queries{
  //     {"a", "b"},
  //     {"e", "e"},
  //     {"x", "x"},
  // };

  // expected: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
  // vector<vector<string>> equations{
  //     {"a", "b"},
  //     {"b", "c"},
  // };

  // vector<double> values{
  //     2.0, 3.0};

  // vector<vector<string>> queries{
  //     {"a", "c"},
  //     {"b", "a"},
  //     {"a", "e"},
  //     {"a", "a"},
  //     {"x", "x"},
  // };

  // expected: [1.13333,16.80000,1.50000,1.00000,0.05952,2.26667,0.44118,-1.00000,-1.00000]
  vector<vector<string>> equations{
      {"x1", "x2"},
      {"x2", "x3"},
      {"x1", "x4"},
      {"x2", "x5"},
  };

  vector<double> values{
      3.0, 0.5, 3.4, 5.6};

  vector<vector<string>> queries{
      {"x2", "x4"},
      {"x1", "x5"},
      {"x1", "x3"},
      {"x5", "x5"},
      {"x5", "x1"},
      {"x3", "x4"},
      {"x4", "x3"},
      {"x6", "x6"},
      {"x0", "x0"},
  };

  vector<double> sol = Solution().calcEquation(equations, values, queries);

  cout << "-------------------\n";
  cout << "SOLUTION: ";
  for (auto s : sol)
  {
    cout << s << " ";
  }
  cout << "\n";
  cout << "-------------------\n";
}