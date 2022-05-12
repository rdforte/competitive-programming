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
      // roots[rootB]->parent = rootA;
      roots[b]->parent = a;

      string head = b;
      double bWeight = roots[b]->weight;
      do
      {
        // string parent = roots[head]->parent
        roots[roots[head]->parent]->weight *= (result * bWeight);
        head = roots[head]->parent;
      } while (roots[head]->parent != head);

      // if (roots[a]->parent == a)
      // {
      //   roots[a]->weight *= result;
      //   return;
      // }

      // for (auto r : roots)
      // {
      //   if (r.second->parent == rootA && r.first != b)
      //   {
      //     r.second->weight *= result;
      //   }
      // }
    }
  }

  bool isConnected(string a, string b)
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

    // for (auto r : uf.getRoots())
    // {
    //   cout << r.first << " : " << r.second->parent << " : " << r.second->weight << "\n";
    // }

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

  vector<vector<string>> equations{
      {"a", "b"},
      {"b", "c"},
  };

  vector<double> values{
      2.0, 3.0};

  vector<vector<string>> queries{
      {"a", "c"},
      {"b", "a"},
      {"a", "e"},
      {"a", "a"},
      {"x", "x"},
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