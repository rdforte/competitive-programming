// https://leetcode.com/problems/evaluate-division/
#include <iostream>
#include <vector>
#include <unordered_map>
#include <string>
#include <set>

using namespace std;

struct Node
{
  string parent;
  double weight;

  Node() : parent(""), weight(1) {}
  Node(string p) : parent(p), weight(1) {}
  Node(string p, double w) : parent(p), weight(w) {}
};

// Apprach Union Find With Weights
class Solution
{
private:
  unordered_map<string, Node> graphIdWeight;

public:
  Solution()
  {
    graphIdWeight = {};
  }

  vector<double> calcEquation(vector<vector<string>> &equations, vector<double> &values, vector<vector<string>> &queries)
  {

    // Build Union Groups
    for (int i = 0; i < equations.size(); i++)
    {
      vector<string> equation = equations[i];
      string dividend = equation[0];
      string divisor = equation[1];
      double quotient = values[i];

      // call union
      unionise(dividend, divisor, quotient);
    }

    // run through the evaluation with 'lazy' updates in the find function.
    vector<double> results{};

    for (int i = 0; i < queries.size(); i++)
    {
      string dividend = queries[i][0], divisor = queries[i][1];

      if (graphIdWeight.find(dividend) == graphIdWeight.end() || graphIdWeight.find(divisor) == graphIdWeight.end())
      {
        results.push_back(-1);
      }
      else
      {
        Node dividendEntry = find(dividend);
        Node divisorEntry = find(divisor);

        string dividendGraphId = dividendEntry.parent;
        string divisorGraphId = divisorEntry.parent;
        double dividendWeight = dividendEntry.weight;
        double divisorWeight = divisorEntry.weight;

        if (dividendGraphId != divisorGraphId)
        {
          // not part of the same group
          results.push_back(-1);
        }
        else
        {
          results.push_back(dividendWeight / divisorWeight);
        }
      }
    }

    for (auto gi : graphIdWeight)
    {
      cout << gi.first << " : " << gi.second.parent << " == " << gi.second.weight << endl;
    }

    return results;
  }

private:
  Node find(string nodeId)
  {
    // If there is no node yet then create it.
    if (graphIdWeight.find(nodeId) == graphIdWeight.end())
    {
      graphIdWeight[nodeId] = Node(nodeId);
    }

    Node entry = graphIdWeight[nodeId];

    if (entry.parent != nodeId)
    {
      // recursively get the root
      Node newEntry = find(entry.parent);
      graphIdWeight[nodeId] = Node(newEntry.parent, entry.weight * newEntry.weight);
    }

    return graphIdWeight[nodeId];
  }

  void unionise(string dividend, string divisor, double value)
  {
    Node dividendEntry = find(dividend);
    Node divisorEntry = find(divisor);

    string dividendGraphId = dividendEntry.parent;
    string divisorGraphId = divisorEntry.parent;
    double dividendWeight = dividendEntry.weight;
    double divisorWeight = divisorEntry.weight;

    // merge the two groups together
    if (dividendGraphId != divisorGraphId)
    {
      graphIdWeight[dividendGraphId] = Node(divisorGraphId, (divisorWeight * value) / dividendWeight);
    }
  }
};

int main()
{

  // expected: [1.33333, 1.00000, -1.00000]
  vector<vector<string>> equations{
      {"a", "e"},
      {"b", "e"},
  };

  vector<double> values{
      4.0, 3.0};

  vector<vector<string>> queries{
      {"a", "b"},
      {"e", "e"},
      {"x", "x"},
  };

  // expected: [3.75000,0.40000,5.00000,0.20000]
  // vector<vector<string>> equations{
  //     {"a", "b"},
  //     {"b", "c"},
  //     {"bc", "cd"},
  // };

  // vector<double> values{
  //     1.5, 2.5, 5.0};

  // vector<vector<string>> queries{
  //     {"a", "c"},
  //     {"c", "b"},
  //     {"bc", "cd"},
  //     {"cd", "bc"},
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
  // vector<vector<string>> equations{
  //     {"x1", "x2"},
  //     {"x2", "x3"},
  //     {"x1", "x4"},
  //     {"x2", "x5"},
  // };

  // vector<double> values{
  //     3.0, 0.5, 3.4, 5.6};

  // vector<vector<string>> queries{
  //     {"x2", "x4"},
  //     {"x1", "x5"},
  //     {"x1", "x3"},
  //     {"x5", "x5"},
  //     {"x5", "x1"},
  //     {"x3", "x4"},
  //     {"x4", "x3"},
  //     {"x6", "x6"},
  //     {"x0", "x0"},
  // };

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