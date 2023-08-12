/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
const lowestCommonAncestor = function (node, p, q) {
  if (node.val > p.val && node.val > q.val) {
    return lowestCommonAncestor(node.left, p, q);
  } else if (node.val < p.val && node.val < q.val) {
    return lowestCommonAncestor(node.right, p, q);
  } else {
    return node;
  }
};
