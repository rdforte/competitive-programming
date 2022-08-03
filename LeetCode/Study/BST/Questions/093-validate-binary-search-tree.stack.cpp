// https://leetcode.com/problems/validate-binary-search-tree/
#include <queue>
#include <iostream>
#include <stack>

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

struct TreeNodeBounds
{
  TreeNode *upper;
  TreeNode *lower;
  TreeNode *parent;

  TreeNodeBounds() : parent(nullptr), upper(nullptr), lower(nullptr) {}
  TreeNodeBounds(TreeNode *parent) : parent(parent), upper(nullptr), lower(nullptr) {}
  TreeNodeBounds(TreeNode *parent, TreeNode *upper) : parent(parent), upper(upper), lower(nullptr) {}
  TreeNodeBounds(TreeNode *parent, TreeNode *upper, TreeNode *lower) : parent(parent), upper(upper), lower(lower) {}
};

class Solution
{
public:
  bool isValidBST(TreeNode *root)
  {

    stack<TreeNodeBounds> s;

    s.push({root});

    while (!s.empty())
    {
      auto node = s.top();
      s.pop();

      if ((node.upper != nullptr && node.parent->val >= node.upper->val) or
          (node.lower != nullptr && node.parent->val <= node.lower->val))
      {
        return false;
      }

      // add left node to stack
      if (node.parent->left != nullptr)
        s.push({node.parent->left, node.parent, node.lower});

      // add right node
      if (node.parent->right != nullptr)
      {
        s.push({node.parent->right, node.upper, node.parent});
      }
    }

    return true;
  }
};

int main()
{
  TreeNode *val3 = new TreeNode(3);
  TreeNode *val7 = new TreeNode(7);
  TreeNode *val4 = new TreeNode(4);
  TreeNode *val6 = new TreeNode(6, val3, val7);
  TreeNode *root = new TreeNode(5, val4, val6);

  // TreeNode *val3 = new TreeNode(3);
  // TreeNode *val7 = new TreeNode(7);
  // TreeNode *root = new TreeNode(4, val3, val7);

  cout << boolalpha;

  cout << Solution().isValidBST(root);
}