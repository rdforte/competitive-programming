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

typedef long long ll;
class Solution
{
public:
  long long kthLargestLevelSum(TreeNode *root, int k)
  {
    queue<TreeNode *> q;
    q.push(root);
    vector<ll> v;

    while (!q.empty())
    {
      ll sum = 0;
      ll qSize = q.size();
      while (qSize > 0)
      {
        auto n = q.front();
        q.pop();
        sum += n->val;
        qSize--;
        if (n->left)
          q.push(n->left);
        if (n->right)
          q.push(n->right);
      }
      v.push_back(sum);
    }
    sort(v.begin(), v.end(), greater<ll>());
    return k > v.size() ? -1 : v[k - 1];
  }
};

int main()
{
  // Solution
}
