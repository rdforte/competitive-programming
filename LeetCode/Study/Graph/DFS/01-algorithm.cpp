#include <iostream>
#include <vector>
#include <unordered_map>
#include <string>

using namespace std;

// ITERATIVE APPROACH
void depthFirstPrint(unordered_map<string, vector<string>> graph, string source)
{
  vector<string> stack(1, source);

  while (stack.size() > 0)
  {
    string node = stack[stack.size() - 1];
    stack.pop_back();

    cout << node;

    vector<string> neighbours = graph[node];

    for (auto n : neighbours)
    {
      stack.push_back(n);
    }
  }
};

// RECURSIVE APPROACH - utilising the callstack so no need to use a stack like the above approach.
void depthFirstPrintRecursive(unordered_map<string, vector<string>> graph, string source)
{
  cout << source;
  for (auto n : graph[source])
  {
    depthFirstPrintRecursive(graph, n);
  }
};

int main()
{
  unordered_map<string, vector<string>> graph{
      {"a", {"c", "b"}},
      {"b", {"d"}},
      {"c", {"e"}},
      {"d", {"f"}},
      {"e", {}},
      {"f", {}},
  };

  depthFirstPrint(graph, "a"); // abdfce
  cout << "\n-------------------------\n";
  depthFirstPrintRecursive(graph, "a"); // abdfce
}