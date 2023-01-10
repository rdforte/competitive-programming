// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
#include <bits/stdc++.h>

using namespace std;

struct TreeNode
{
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution
{
public:
  TreeNode *lowestCommonAncestor(TreeNode *root, TreeNode *p, TreeNode *q)
  {
    if (q->val < root->val and p->val < root->val)
    {
      return lowestCommonAncestor(root->left, p, q);
    }
    else if (q->val > root->val and p->val > root->val)
    {
      return lowestCommonAncestor(root->right, p, q);
    }
    else
      return root;
  }
};

int main()
{
  TreeNode *left = new TreeNode(2);
  TreeNode *right = new TreeNode(8);
  TreeNode *root = new TreeNode(6);
  root->left = left;
  root->right = right;

  auto sol = Solution().lowestCommonAncestor(root, left, right);
  cout << sol->val;
}