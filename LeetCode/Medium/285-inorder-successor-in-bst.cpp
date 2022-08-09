// https://leetcode.com/problems/inorder-successor-in-bst/
#include <vector>
#include <stack>
#include <algorithm>
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
  TreeNode *inorderSuccessor(TreeNode *root, TreeNode *p)
  {
    TreeNode *successor = nullptr;

    while (root != nullptr)
    {
      if (p->val >= root->val)
      {
        root = root->right;
      }
      else
      {
        successor = root;
        root = root->left;
      }
    }

    return successor;
  }
};

int main()
{
  TreeNode *one = new TreeNode(1);
  TreeNode *two = new TreeNode(2, one, nullptr);
  TreeNode *four = new TreeNode(4);
  TreeNode *seven = new TreeNode(7);
  TreeNode *nine = new TreeNode(9);
  TreeNode *eight = new TreeNode(8, seven, nine);
  TreeNode *three = new TreeNode(3, two, four);
  TreeNode *five = new TreeNode(5, three, eight);

  TreeNode *sol = Solution().inorderSuccessor(five, nine);
  if (sol == NULL)
    cout << "null";
  else
    cout << sol->val;
}