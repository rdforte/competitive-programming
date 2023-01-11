#include "bits/stdc++.h"

using namespace std;

/**
 * When completing this question if we look at the binary search tree example in the assets folder we can see that each
 * time we look at a node we find ourselves saying "it must be greater than A but less than B" this then leads us to
 * the assumption that the node must be within a Range. We then must keep track of the range at each traversal to see
 * if the node adheres to the rules of a Binary Search Tree.
 */

struct TreeNode
{
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

bool dfsValidBST(TreeNode *root, long lower, long upper)
{
  if (root == nullptr)
    return true;

  if (root->val >= upper or root->val <= lower)
    return false;

  bool leftValid = dfsValidBST(root->left, lower, root->val);
  bool rightValid = dfsValidBST(root->right, root->val, upper);

  return leftValid and rightValid;
}

bool isValidBST(TreeNode *root)
{
  return dfsValidBST(root, LLONG_MIN, LLONG_MAX);
}
