// https://leetcode.com/problems/inorder-successor-in-bst/
#include <vector>
#include <stack>
#include <algorithm>

using namespace std;

struct TreeNode
{
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

struct NodeSuccessor
{
  TreeNode *node;
  vector<TreeNode *> successors;

  NodeSuccessor() : node(nullptr), successors({}) {}
  NodeSuccessor(TreeNode *n) : node(n), successors({}) {}
  NodeSuccessor(TreeNode *n, vector<TreeNode *> p) : node(n), successors(p) {}
};

class Solution
{
public:
  TreeNode *inorderSuccessor(TreeNode *root, TreeNode *p)
  {

    if (root == p)
      return root;

    stack<NodeSuccessor> s;

    s.push({root});

    while (!s.empty())
    {
      auto n = s.top();
      s.pop();

      if (n.node == p)
      {
        sort(n.successors.begin(), n.successors.end(), [](auto left, auto right)
             { return left->val < right->val; });
        for (auto pNode : n.successors)
        {
          if (pNode->val > p->val)
          {
            return pNode;
          }
        }
      }

      if (n.node->left != nullptr)
      {
        vector<TreeNode *> v(n.successors.begin(), n.successors.end());
        v.push_back(n.node);
        s.push({n.node->left,
                v});
      }

      if (n.node->right != nullptr)
      {
        vector<TreeNode *> v(n.successors.begin(), n.successors.end());
        v.push_back(n.node);
        s.push({n.node->right,
                v});
      }
    }

    return NULL;
  }
};