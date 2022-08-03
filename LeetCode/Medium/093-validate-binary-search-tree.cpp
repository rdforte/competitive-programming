// https://leetcode.com/problems/validate-binary-search-tree/
#include <queue>
#include <iostream>

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

class Solution
{
public:
  bool isValidBST(TreeNode *root)
  {
    return recursivelyCheckSubTree(root, nullptr, nullptr);
  }

  bool recursivelyCheckSubTree(TreeNode *parent, TreeNode *lower, TreeNode *upper)
  {
    // null is a valid tree.
    if (parent == nullptr)
      return true;

    if ((upper != nullptr && parent->val >= upper->val) or
        (lower != nullptr && parent->val <= lower->val))
      return false;

    return recursivelyCheckSubTree(parent->right, parent, upper) and
           recursivelyCheckSubTree(parent->left, lower, parent);
  }
};

int main()
{
  TreeNode *val3 = new TreeNode(3);
  TreeNode *val7 = new TreeNode(7);
  TreeNode *val4 = new TreeNode(4);
  TreeNode *val6 = new TreeNode(6, val3, val7);
  TreeNode *root = new TreeNode(5, val4, val6);

  cout << boolalpha;

  cout << Solution().isValidBST(root);
}