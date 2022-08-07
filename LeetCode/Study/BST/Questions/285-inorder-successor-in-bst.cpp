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

    if (root == p)
      return root;

    stack<TreeNode *> s;
    vector<TreeNode *> prev{};

    s.push(root);

    bool isFound = false;

    while (!s.empty())
    {
      auto n = s.top();
      s.pop();

      prev.push_back(n);

      if (n == p)
      {
        isFound = true;
        if (n->right != nullptr)
        {
          s.push(n->right);
        }
        else
        {
          TreeNode *node = new TreeNode(INT_MAX);
          for (int i = prev.size() - 1; i >= 0; i--)
          {
            if (prev[i]->val > p->val && prev[i]->val < node->val)
            {
              node = prev[i];
            }
          }
          return node->val == INT_MAX ? NULL : node;
        }
      }
      else if (isFound)
      {
        if (n->left != nullptr)
        {
          s.push(n->left);
        }
        else
        {
          return prev[prev.size() - 1];
        }
      }
      else
      {
        if (n->right != nullptr)
          s.push(n->right);

        if (n->left != nullptr)
          s.push(n->left);
      }
    }

    return NULL;
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