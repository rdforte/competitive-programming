#include "../../../stdc++.h"

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

bool dfsValidBST(TreeNode *root, long lower, long upper)
{
  if (root == nullptr)
    return true;

  if (root->val >= upper || root->val <= lower)
    return false;

  bool leftValid = dfsValidBST(root->left, lower, root->val);
  bool rightValid = dfsValidBST(root->right, root->val, upper);

  return leftValid && rightValid;
}

bool isValidBST(TreeNode *root)
{
  return dfsValidBST(root, LLONG_MIN, LLONG_MAX);
}

int main()
{
}