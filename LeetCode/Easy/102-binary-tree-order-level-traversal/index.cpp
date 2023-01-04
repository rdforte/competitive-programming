// https://leetcode.com/problems/binary-tree-level-order-traversal/?envType=study-plan&id=level-1
#include <bits/stdc++.h>

using namespace std;

struct TreeNode
{
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class SolutionOne
{
public:
  vector<vector<int>> levelOrder(TreeNode *root)
  {
    vector<vector<int>> res;

    if (!root)
      return res;

    queue<pair<int, TreeNode *>> q;
    q.push({0, root});

    while (!q.empty())
    {
      auto node = q.front();
      q.pop();

      if (node.first + 1 > res.size())
      {
        res.push_back({node.second->val});
      }
      else
      {
        res[node.first].push_back(node.second->val);
      }

      if (node.second->left)
      {
        q.push({node.first + 1, node.second->left});
      }

      if (node.second->right)
      {
        q.push({node.first + 1, node.second->right});
      }
    }

    return res;
  }
};

class Solution
{
public:
  vector<vector<int>> levelOrder(TreeNode *root)
  {
    vector<vector<int>> res;

    if (!root)
      return res;

    queue<TreeNode *> q;
    q.push(root);

    while (!q.empty())
    {
      vector<int> temp;
      int qSize = q.size();
      for (int i = 0; i < qSize; i++)
      {
        auto node = q.front();
        q.pop();
        if (node->left)
          q.push(node->left);
        if (node->right)
          q.push(node->right);
        temp.push_back(node->val);
      }
      res.push_back(temp);
    }

    return res;
  }
};