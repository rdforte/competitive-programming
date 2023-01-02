// https://leetcode.com/problems/n-ary-tree-preorder-traversal/?envType=study-plan&id=level-1
#include <bits/stdc++.h>

using namespace std;

class Node
{
public:
  int val;
  vector<Node *> children;

  Node() {}

  Node(int _val)
  {
    val = _val;
  }

  Node(int _val, vector<Node *> _children)
  {
    val = _val;
    children = _children;
  }
};

class Solution
{
public:
  vector<int> preorder(Node *root)
  {
    vector<int> res;
    if (!root)
      return res;
    dfs(root, res);
    return res;
  }

  void dfs(Node *root, vector<int> &nodes)
  {
    nodes.push_back(root->val);

    if (root->children.size() == 0)
      return;

    for (int i = 0; i < root->children.size(); i++)
    {
      dfs(root->children[i], nodes);
    }
  }
};